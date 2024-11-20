package cmd

import (
	"fmt"
	"github.com/spf13/cast"
	"pgii/src/pg/db"
	"pgii/src/pg/global"
	"pgii/src/util"
	"strings"
)

var DefaultVersion = 15

func (s *Params) Show() {
	if len(s.Param) == global.ZeroCMDLength {
		util.PrintColorTips(util.LightRed, global.CmdLineWrong)
	}

	//获取CMD
	cmd := strings.ToLower(strings.Trim(s.Param[0], ""))
	switch CheckParamType(cmd) {
	case global.VersionStyle:
		s.ShowVersion()
	case global.DatabaseStyle:
		s.ShowDatabases()
	case global.TableStyle, global.ViewStyle:
		s.ShowTableView(cmd)
	case global.SelectStyle:
		util.PrintColorTips(util.LightGreen, fmt.Sprintf("DataBase:%s;Schema:%s", *global.Database, db.P.Schema))
	case global.SchemaStyle:
		s.ShowSchema()
	case global.TriggerStyle:
		s.ShowTrigger()
	case global.ConnectionStyle:
		s.ShowConnection()
	case global.ProcessStyle:
		s.ShowProcess()
	default:
		util.PrintColorTips(util.LightRed, global.CmdLineWrong)
	}
}

// ShowProcess 查看当前进程
func (s *Params) ShowProcess() {
	//存在第三个
	var param []interface{}
	if len(s.Param) > global.OneCMDLength {
		//sCMD
		thirdCmd := strings.ToLower(strings.Trim(s.Param[1], ""))
		param = append(param, thirdCmd)
		//查询PID的范围
		if thirdCmd == "pid" {
			//len
			if len(s.Param) != global.FiveCMDLength {
				util.PrintColorTips(util.LightRed, global.CmdLineWrong)
				return
			}

			//judge start end
			startPid := cast.ToInt(s.Param[2])
			endPid := cast.ToInt(s.Param[4])
			if startPid > endPid {
				util.PrintColorTips(util.LightRed, global.StartThanEndError)
				return
			}

			midCmd := strings.ToLower(strings.Trim(s.Param[3], ""))
			if midCmd != "and" {
				util.PrintColorTips(util.LightRed, global.CmdLineWrong)
				return
			}

			param = append(param, startPid, endPid)
		}
	}

	//proc
	if procList, err := db.P.Process(param...); err == nil {
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
		db.ShowTable(db.ProcessHeader, ps)
	} else {
		util.PrintColorTips(util.LightRed, global.ShowDatabaseError, err.Error())
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
	mc, err := db.P.GetConnectionNums(global.MaxConnections)
	if err == nil {
		if _, ok := mc["max_connections"]; ok {
			maxConnection = cast.ToInt(mc["max_connections"])
		}
	}

	sc, err := db.P.GetConnectionNums(global.SuperuserReservedConnections)
	if err == nil {
		if _, ok := sc["superuser_reserved_connections"]; ok {
			superConnection = cast.ToInt(sc["superuser_reserved_connections"])
		}
	}

	//查剩余连接数
	rc, err := db.P.GetUseConnection(global.RemainingConnections)
	if err == nil {
		if _, ok := rc["conn_nums"]; ok {
			remainingConnection = cast.ToInt(rc["conn_nums"])
		}
	}

	//正在使用连接数
	uc, err := db.P.GetUseConnection(global.InUseConnections)
	if err == nil {
		if _, ok := uc["conn_nums"]; ok {
			inUseConnection = cast.ToInt(uc["conn_nums"])
		}
	}

	//get data
	data := []interface{}{maxConnection, superConnection, remainingConnection, inUseConnection}

	//show table
	db.ShowTable(db.ConnectionHeader, [][]interface{}{data})
}

func (s *Params) ShowTrigger() {
	//define
	var triggerInfo []map[string]interface{}
	var err error

	if len(s.Param) == global.OneCMDLength {
		triggerInfo, err = db.P.Trigger("", "")
	} else {
		util.PrintColorTips(util.LightRed, global.ShowTriggerCmdFailed)
		return
	}

	fqv := ""
	if len(s.Param) == global.ThreeCMDLength {
		//带有like 或 filter
		sonCmd := strings.ToLower(strings.Trim(s.Param[1], ""))
		value := strings.ToLower(strings.Trim(s.Param[2], ""))
		fqv = value
		triggerInfo, err = db.P.Trigger(sonCmd, value)
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
		db.ShowTable(db.TriggerShowHeader, trbs)
	}
}

// ShowSchema 获取模式
func (s *Params) ShowSchema() {
	if scList, err := db.P.SchemaNS(); err == nil {
		//序列化输出
		var scbs [][]interface{}
		for _, v := range scList {
			var sbs []interface{}
			sbs = append(sbs,
				v["oid"],
				util.If(cast.ToString(v["nspname"]) == db.P.Schema,
					util.SetColor(cast.ToString(v["nspname"])+"[✓]", util.LightGreen),
					cast.ToString(v["nspname"])),
				db.P.GetRoleNameByOid(cast.ToInt(v["nspowner"])),
				v["nspacl"],
			)
			//填入数组
			scbs = append(scbs, sbs)
		}
		db.ShowTable(db.SchemaShowHeader, scbs)
	}
}

// ShowVersion 获取版本
func (s *Params) ShowVersion() {
	//获取版本
	version, _ := db.P.Version()
	//序列化输出
	data := []interface{}{"PostgresSql", version}
	db.ShowTable(db.VersionShowHeader, [][]interface{}{data})
}

// ShowDatabases 列出所有的数据库
func (s *Params) ShowDatabases() {
	//获取版本
	version := 0
	ver, _ := db.P.Version()
	vers := strings.Split(ver, ".")
	if len(vers) > 0 {
		version = cast.ToInt(vers[0])
	}

	//judge version support 15.X
	if dbList, err := db.P.Database(version); err == nil {
		var dbs [][]interface{}
		for _, v := range dbList {
			var sbs []interface{}
			//oid
			sbs = append(sbs,
				v["oid"],
				util.If(cast.ToString(v["datname"]) == *global.Database,
					util.SetColor(cast.ToString(v["datname"])+"[✓]", util.LightGreen),
					cast.ToString(v["datname"])),
				db.P.GetRoleNameByOid(cast.ToInt(v["datdba"])),
				db.P.GetEncodingChar(cast.ToInt(v["encoding"])),
				v["datcollate"],
				v["datctype"],
				v["datallowconn"],
				v["datconnlimit"])
			//judge version
			if version < DefaultVersion {
				sbs = append(sbs, v["datlastsysoid"])
			}

			sbs = append(sbs,
				//v["datlastsysoid"],
				db.P.GetTableSpaceNameByOid(cast.ToInt(v["dattablespace"])),
				v["size"],
			)
			//加入数据列
			dbs = append(dbs, sbs)
		}

		//judge version show
		if version >= 15 {
			db.ShowTable(db.Database15ShowHeader, dbs)
		} else {
			db.ShowTable(db.DatabaseShowHeader, dbs)
		}
	} else {
		util.PrintColorTips(util.LightRed, global.ShowDatabaseError, err.Error())
	}
}

func (s *Params) ShowTableView(cmd string) {
	//增加过滤过功能
	if len(s.Param) == global.ThreeCMDLength {
		//带有like 或 filter
		sonCmd := strings.ToLower(strings.Trim(s.Param[1], ""))
		//过滤参数处理
		param := strings.Replace(s.Param[2], "'", "", -1)
		param = strings.Replace(param, "\"", "", -1)
		params := strings.Split(param, "|")

		if !util.InArray(sonCmd, global.EqualAndFilter...) ||
			!util.InArray(cmd, global.TableAndView...) {
			fmt.Println("Failed:CmdLine Show Table Or View filter is Wrong!")
			return
		}

		//校验是查表还是视图
		util.IfCmdFunc(
			util.InArray(cmd, global.TableVar...),
			sonCmd,
			params,
			s.ShowTables,
			s.ShowView,
		)

	} else {
		if !util.InArray(cmd, global.TableAndView...) {
			fmt.Println("Failed:CmdLine Show Table Or View filter is Wrong!")
			return
		}

		util.IfCmdFunc(
			util.InArray(cmd, global.TableVar...),
			"",
			nil,
			s.ShowTables,
			s.ShowView,
		)
	}

}

// ShowTables 查询表信息列表
func (s *Params) ShowTables(cmd string, filter ...string) {
	if tb, err := db.P.Tables(cmd, filter...); err == nil {
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
				classInfo, err := db.P.GetPgClassForTbName(tableName)
				//judge
				if err != nil && len(classInfo) == 0 {
					return
				}
				//range
				for _, st := range []int{global.TableStyle, global.IndexStyle} {
					if sizeInfo, err := db.P.GetSizeInfo(st, fmt.Sprintf(`"%s".%s`, schemaName, tableName)); err == nil {
						switch st {
						case global.TableStyle:
							tbSize = sizeInfo["size"]
						case global.IndexStyle:
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
		db.ShowTable(db.TableShowHeader, tbs)
	}
}

// ShowView 查询视图
func (s *Params) ShowView(cmd string, filter ...string) {
	if tb, err := db.P.Views(cmd, filter...); err == nil {
		//序列化输出
		var vbs [][]interface{}
		for _, v := range tb {
			var sbs []interface{}
			//oid
			sbs = append(sbs, v["schemaname"], v["viewname"], v["viewowner"])
			vbs = append(vbs, sbs)
		}
		//打印表格
		db.ShowTable(db.ViewShowHeader, vbs)
	}
}
