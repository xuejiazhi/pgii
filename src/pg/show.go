package pg

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cast"
	"os"
	"pgii/src/util"
	"strings"
)

func Show(cmdList []string) {
	if len(cmdList) > 0 {
		cmd := strings.ToLower(strings.Trim(cmdList[0], ""))
		if cmd != "" {
			switch cmd {
			case "ver", "version":
				ShowVersion()
			case "db", "database":
				ShowDatabases()
			case "tb", "table", "view", "vw":
				ShowTableView(cmd, cmdList)
				return
			case "sd", "selectdb":
				fmt.Println("DataBase:", *Database, ";Schema:", P.Schema)
			case "sc", "schema": //查询schema
				ShowSchema()
			case "tg", "trigger": //查询trigger
				ShowTrigger(cmdList)
			default:
				fmt.Println("Failed:CmdLine is Wrong!")
			}
		}
	} else {
		fmt.Println("Failed:CmdLine is Wrong!")
	}
}

func ShowTrigger(cmdList []string) {
	triggerInfo := []map[string]interface{}{}
	fqv := ""
	var err error
	if len(cmdList) == 3 {
		//带有like 或 filter
		sonCmd := strings.ToLower(strings.Trim(cmdList[1], ""))
		value := strings.ToLower(strings.Trim(cmdList[2], ""))
		fqv = value
		triggerInfo, err = P.Trigger(sonCmd, value)
	} else {
		if len(cmdList) == 1 {
			triggerInfo, err = P.Trigger("", "")
		} else {
			fmt.Println(ShowTriggerCmdFailed)
		}
	}

	//序列化输出
	if err == nil && len(triggerInfo) > 0 {
		//序列化输出
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(TriggerHeader)
		var tbs []table.Row
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
			tbs = append(tbs, sbs)
		}
		t.AppendRows(tbs)
		t.Render()
	}
}

// ShowSchema 获取模式
func ShowSchema() {
	if scList, err := P.SchemaNS(); err == nil {
		//序列化输出
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetStyle(table.StyleLight)
		t.AppendHeader(table.Row{"#oid", "SchemaName", "Owner", "Acl"})
		var dbs []table.Row
		for _, v := range scList {
			var sbs []interface{}
			sbs = append(sbs,
				v["oid"],
				util.If(cast.ToString(v["nspname"]) == P.Schema, cast.ToString(v["nspname"])+"[✓]", cast.ToString(v["nspname"])),
				P.GetRoleNameByOid(cast.ToInt(v["nspowner"])),
				v["nspacl"],
			)
			//填入数组
			dbs = append(dbs, sbs)
		}
		t.AppendRows(dbs)
		t.Render()
	}
}

// ShowVersion 获取版本
func ShowVersion() {
	//获取版本
	version, _ := P.Version()
	//序列化输出
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Version"})
	t.AppendRows([]table.Row{
		{"PostgresSql", version},
	})
	t.Render()
}

// ShowDatabases 列出所有的数据库
func ShowDatabases() {
	if dbList, err := P.Database(); err == nil {
		//序列化输出
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#oid", "DbName", "Auth", "Encoding", "LC_COLLATE", "LC_CTYPE", "AllowConn", "ConnLimit", "LastSysOid", "TableSpace", "size"})
		var dbs []table.Row
		for _, v := range dbList {
			var sbs []interface{}
			//oid
			sbs = append(sbs,
				v["oid"],
				util.If(cast.ToString(v["datname"]) == *Database,
					util.SetColor(cast.ToString(v["datname"])+"[✓]", util.LightSeaBlue),
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
			//
			dbs = append(dbs, sbs)
		}
		t.AppendRows(dbs)
		t.Render()
	} else {
		fmt.Println("Failed:Show DataBase is Wrong! error ", err.Error())
	}
}

func ShowTableView(cmd string, cmdList []string) {
	//增加过滤过功能
	if len(cmdList) == 3 {
		//带有like 或 filter
		sonCmd := strings.ToLower(strings.Trim(cmdList[1], ""))
		//过滤参数处理
		param := strings.Replace(cmdList[2], "'", "", -1)
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
			ShowTables,
			ShowView,
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
			ShowTables,
			ShowView,
		)
	}

}

func ShowTables(cmd string, filter ...string) {
	if tb, err := P.Tables(cmd, filter...); err == nil {
		//序列化输出
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(TableHeader)
		var tbs []table.Row
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
		t.AppendRows(tbs)
		t.Render()
	}
}

func ShowView(cmd string, filter ...string) {
	if tb, err := P.Views(cmd, filter...); err == nil {
		//序列化输出
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(ViewHeader)
		var tbs []table.Row
		for _, v := range tb {
			var sbs []interface{}
			//oid
			sbs = append(sbs, v["schemaname"], v["viewname"], v["viewowner"])
			tbs = append(tbs, sbs)
		}
		t.AppendRows(tbs)
		t.Render()
	}
}
