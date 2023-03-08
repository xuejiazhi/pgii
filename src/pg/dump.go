package pg

import (
	"fmt"
	"pgii/src/util"
)

// Dump DUMP PGSQL
func Dump(params ...string) {
	//参数长度至少为1位
	if len(params) < 1 {
		fmt.Println(util.SetColor(DumpFailed, util.LightRed))
		return
	}

	//查看Dump的类型
	style := util.TrimLower(params[0])
	switch style {
	case "tb", "table": //生成表的备份
		DumpTable(params[1])
	case "sc", "schema": //生成schema的备份
		DumpSchema()
	case "db", "database": //生成database的备份
		DumpDatabase()
	default:
		fmt.Println(util.SetColor(DumpFailed, util.LightRed))
	}

}

func DumpSchema() {

}

// DumpTable 生成一个创建Table 的SQL
// tbName 表名
func DumpTable(tbName string) {
	if tbName == "" {
		return
	}
	//校验表是否存在
	if tbInfo, err := P.GetTableByName(tbName); err != nil || len(tbInfo) == 0 {
		fmt.Println(DumpFailedNoTable)
		return
	}

}

// DumpDatabase 生成Database的备份
func DumpDatabase() {

}
