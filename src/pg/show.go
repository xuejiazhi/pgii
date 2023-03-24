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
	default:
		util.PrintColorTips(util.LightRed, CmdLineWrong)
	}
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

		if !util.InArray(sonCmd, EqualAndFilter) ||
			!util.InArray(cmd, TableAndView) {
			fmt.Println("Failed:CmdLine Show Table Or View filter is Wrong!")
			return
		}

		//校验是查表还是视图
		util.IfCmdFunc(
			util.InArray(cmd, TableVar),
			sonCmd,
			params,
			s.ShowTables,
			s.ShowView,
		)

	} else {
		if !util.InArray(cmd, TableAndView) {
			fmt.Println("Failed:CmdLine Show Table Or View filter is Wrong!")
			return
		}

		util.IfCmdFunc(
			util.InArray(cmd, TableVar),
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
			var sbs []interface{}
			//
			tableName := cast.ToString(v["tablename"])
			if len(filter) > 0 {
				for _, v := range filter {
					tableName = strings.Replace(tableName, v, util.SetColor(v, util.LightGreen), -1)
				}
			}
			//oid
			sbs = append(sbs, v["schemaname"], tableName, v["tableowner"], v["tablespace"])
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
