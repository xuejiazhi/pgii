package pg

import (
	"fmt"
	"github.com/spf13/cast"
	"pgii/src/util"
	"strings"
)

func (s *Params) Show() {
	if len(s.Param) == ZeroCMDLength {
		util.PrintColorTips(util.LightRed, CmdLineWrong)
	}

	//获取CMD
	cmd := strings.ToLower(strings.Trim(s.Param[0], ""))
	switch CheckParamType(cmd) {
	case VersionStyle:
		s.ShowVersion()
	case DatabaseStyle:
		s.ShowDatabases()
	case TableStyle, ViewStyle:
		s.ShowTableView(cmd)
	case SelectStyle:
		util.PrintColorTips(util.LightGreen, fmt.Sprintf("DataBase:%s;Schema:%s", *Database, P.Schema))
	case SchemaStyle:
		s.ShowSchema()
	case TriggerStyle:
		s.ShowTrigger()
	case ConnectionStyle:
		s.ShowConnection()
	case ProcessStyle:
		s.ShowProcess()
	default:
		util.PrintColorTips(util.LightRed, CmdLineWrong)
	}
}

// ShowProcess 查看当前进程
func (s *Params) ShowProcess() {
	//存在第三个
	var param []interface{}
	if len(s.Param) > OneCMDLength {
		//sCMD
		thirdCmd := strings.ToLower(strings.Trim(s.Param[1], ""))
		param = append(param, thirdCmd)
		//查询PID的范围
		if thirdCmd == "pid" {
			//len
			if len(s.Param) != FiveCMDLength {
				util.PrintColorTips(util.LightRed, CmdLineWrong)
				return
			}

			//judge start end
			startPid := cast.ToInt(s.Param[2])
			endPid := cast.ToInt(s.Param[4])
			if startPid > endPid {
				util.PrintColorTips(util.LightRed, StartThanEndError)
				return
			}

			midCmd := strings.ToLower(strings.Trim(s.Param[3], ""))
			if midCmd != "and" {
				util.PrintColorTips(util.LightRed, CmdLineWrong)
				return
			}

			param = append(param, startPid, endPid)
		}
	}

	//proc
	if procList, err := P.Process(param...); err == nil {
		var ps [][]interface{}
		for _, v := range procList {
			var sbs []interface{}
			//oid
			pid := cast.ToString(v["pid"])
			datName := cast.ToString(v["datname"])
			applicationName := cast.ToString(v["application_name"])
			state := cast.ToString(v["state"])
			userName := cast.ToString(v["usename"])
			clientAddr := cast.ToString(v["client_addr"])
			clientPort := cast.ToString(v["client_port"])
			//状态为active
			if cast.ToString(v["state"]) == "active" {
				pid = util.SetColor(pid, util.LightGreen)
				datName = util.SetColor(datName, util.LightGreen)
				applicationName = util.SetColor(applicationName, util.LightGreen)
				state = util.SetColor(state, util.LightGreen)
			}
			//加入数据
			sbs = append(sbs,
				pid,
				datName,
				userName,
				clientAddr,
				clientPort,
				applicationName,
				state,
			)
			//加入数据列
			ps = append(ps, sbs)
		}
		ShowTable(ProcessHeader, ps)
	} else {
		util.PrintColorTips(util.LightRed, ShowDatabaseError, err.Error())
	}
}

// ShowConnection 当看当前链接
func (s *Params) ShowConnection() {
	//define
	//maxConnection 最大连接数
	//superConnection 超级用户保留的连接数
	//remainingConnection 剩余连接数
	//当前正使用的连接数
	maxConnection := 0
	superConnection := 0
	remainingConnection := 0
	inUseConnection := 0
	mc, err := P.GetConnectionNums(MaxConnections)
	if err == nil {
		if _, ok := mc["max_connections"]; ok {
			maxConnection = cast.ToInt(mc["max_connections"])
		}
	}

	sc, err := P.GetConnectionNums(SuperuserReservedConnections)
	if err == nil {
		if _, ok := sc["superuser_reserved_connections"]; ok {
			superConnection = cast.ToInt(sc["superuser_reserved_connections"])
		}
	}

	//查剩余连接数
	rc, err := P.GetUseConnection(RemainingConnections)
	if err == nil {
		if _, ok := rc["conn_nums"]; ok {
			remainingConnection = cast.ToInt(rc["conn_nums"])
		}
	}

	//正在使用连接数
	uc, err := P.GetUseConnection(InUseConnections)
	if err == nil {
		if _, ok := uc["conn_nums"]; ok {
			inUseConnection = cast.ToInt(uc["conn_nums"])
		}
	}

	//get data
	data := []interface{}{maxConnection, superConnection, remainingConnection, inUseConnection}

	//show table
	ShowTable(ConnectionHeader, [][]interface{}{data})
}

func (s *Params) ShowTrigger() {
	//define
	var triggerInfo []map[string]interface{}
	var err error

	if len(s.Param) == OneCMDLength {
		triggerInfo, err = P.Trigger("", "")
	} else {
		util.PrintColorTips(util.LightRed, ShowTriggerCmdFailed)
		return
	}

	fqv := ""
	if len(s.Param) == ThreeCMDLength {
		//带有like 或 filter
		sonCmd := strings.ToLower(strings.Trim(s.Param[1], ""))
		value := strings.ToLower(strings.Trim(s.Param[2], ""))
		fqv = value
		triggerInfo, err = P.Trigger(sonCmd, value)
	}

	//序列化输出
	if err == nil && len(triggerInfo) > 0 {
		//序列化输出
		var trbs [][]interface{}
		for _, v := range triggerInfo {
			var sbs []interface{}
			//
			triggerName := cast.ToString(v["trigger_name"])
			if fqv != "" {
				triggerName = strings.Replace(triggerName, fqv, util.SetColor(fqv, util.LightGreen), -1)
			}
			//oid
			sbs = append(sbs,
				v["trigger_catalog"],
				v["trigger_schema"],
				triggerName,
				v["event_manipulation"],
				v["event_object_table"],
				v["action_orientation"],
				v["action_timing"],
			)
			trbs = append(trbs, sbs)
		}
		ShowTable(TriggerShowHeader, trbs)
	}
}

// ShowSchema 获取模式
func (s *Params) ShowSchema() {
	if scList, err := P.SchemaNS(); err == nil {
		//序列化输出
		var scbs [][]interface{}
		for _, v := range scList {
			var sbs []interface{}
			sbs = append(sbs,
				v["oid"],
				util.If(cast.ToString(v["nspname"]) == P.Schema,
					util.SetColor(cast.ToString(v["nspname"])+"[✓]", util.LightGreen),
					cast.ToString(v["nspname"])),
				P.GetRoleNameByOid(cast.ToInt(v["nspowner"])),
				v["nspacl"],
			)
			//填入数组
			scbs = append(scbs, sbs)
		}
		ShowTable(SchemaShowHeader, scbs)
	}
}

// ShowVersion 获取版本
func (s *Params) ShowVersion() {
	//获取版本
	version, _ := P.Version()
	//序列化输出
	data := []interface{}{"PostgresSql", version}
	ShowTable(VersionShowHeader, [][]interface{}{data})
}

// ShowDatabases 列出所有的数据库
func (s *Params) ShowDatabases() {
	if dbList, err := P.Database(); err == nil {
		var dbs [][]interface{}
		for _, v := range dbList {
			var sbs []interface{}
			//oid
			sbs = append(sbs,
				v["oid"],
				util.If(cast.ToString(v["datname"]) == *Database,
					util.SetColor(cast.ToString(v["datname"])+"[✓]", util.LightGreen),
					cast.ToString(v["datname"])),
				P.GetRoleNameByOid(cast.ToInt(v["datdba"])),
				P.GetEncodingChar(cast.ToInt(v["encoding"])),
				v["datcollate"],
				v["datctype"],
				v["datallowconn"],
				v["datconnlimit"],
				v["datlastsysoid"],
				P.GetTableSpaceNameByOid(cast.ToInt(v["dattablespace"])),
				v["size"],
			)
			//加入数据列
			dbs = append(dbs, sbs)
		}
		ShowTable(DatabaseShowHeader, dbs)
	} else {
		util.PrintColorTips(util.LightRed, ShowDatabaseError, err.Error())
	}
}

func (s *Params) ShowTableView(cmd string) {
	//增加过滤过功能
	if len(s.Param) == ThreeCMDLength {
		//带有like 或 filter
		sonCmd := strings.ToLower(strings.Trim(s.Param[1], ""))
		//过滤参数处理
		param := strings.Replace(s.Param[2], "'", "", -1)
		param = strings.Replace(param, "\"", "", -1)
		params := strings.Split(param, "|")

		if !util.InArray(sonCmd, EqualAndFilter...) ||
			!util.InArray(cmd, TableAndView...) {
			fmt.Println("Failed:CmdLine Show Table Or View filter is Wrong!")
			return
		}

		//校验是查表还是视图
		util.IfCmdFunc(
			util.InArray(cmd, TableVar...),
			sonCmd,
			params,
			s.ShowTables,
			s.ShowView,
		)

	} else {
		if !util.InArray(cmd, TableAndView...) {
			fmt.Println("Failed:CmdLine Show Table Or View filter is Wrong!")
			return
		}

		util.IfCmdFunc(
			util.InArray(cmd, TableVar...),
			"",
			nil,
			s.ShowTables,
			s.ShowView,
		)
	}

}

// ShowTables 查询表信息列表
func (s *Params) ShowTables(cmd string, filter ...string) {
	if tb, err := P.Tables(cmd, filter...); err == nil {
		//序列化输出
		var tbs [][]interface{}
		for _, v := range tb {
			//define
			var sbs []interface{}

			//get schemaName  tableName
			schemaName := cast.ToString(v["schemaname"])
			tableName := func() (tbName string) {
				tbName = cast.ToString(v["tablename"])
				if len(filter) > 0 {
					for _, v := range filter {
						tbName = strings.Replace(tbName, v, util.SetColor(v, util.LightGreen), -1)
					}
				}
				return
			}()

			//判断relation是否存在
			tableSize, indexSize := func() (tbSize, idxSize interface{}) {
				classInfo, err := P.GetPgClassForTbName(tableName)
				//judge
				if err != nil && len(classInfo) == 0 {
					return
				}
				//range
				for _, st := range []int{TableStyle, IndexStyle} {
					if sizeInfo, err := P.GetSizeInfo(st, fmt.Sprintf(`"%s".%s`, schemaName, tableName)); err == nil {
						switch st {
						case TableStyle:
							tbSize = sizeInfo["size"]
						case IndexStyle:
							idxSize = sizeInfo["size"]
						}
					}
				}
				//return
				return
			}()

			//oid
			sbs = append(sbs, v["schemaname"], tableName, v["tableowner"], v["tablespace"], tableSize, indexSize)
			tbs = append(tbs, sbs)
		}
		//打印表格
		ShowTable(TableShowHeader, tbs)
	}
}

// ShowView 查询视图
func (s *Params) ShowView(cmd string, filter ...string) {
	if tb, err := P.Views(cmd, filter...); err == nil {
		//序列化输出
		var vbs [][]interface{}
		for _, v := range tb {
			var sbs []interface{}
			//oid
			sbs = append(sbs, v["schemaname"], v["viewname"], v["viewowner"])
			vbs = append(vbs, sbs)
		}
		//打印表格
		ShowTable(ViewShowHeader, vbs)
	}
}
