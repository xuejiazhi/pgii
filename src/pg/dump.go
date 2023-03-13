package pg

import (
	"fmt"
	"os"
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
		fmt.Println(util.SetColor(DumpFailedNoTable, util.LightRed))
		return
	}

	//打开要生成的文件句柄
	f, _ := os.OpenFile("data.pgi", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer f.Close()
	//生成Table 的DDL
	tbsql := []byte(getTableDdlSql(tbName))

	//压缩数据
	util.Compress(&tbsql)
	//写入文件
	_, _ = f.Write(tbsql)

	//处理SQL语句
	//获取表的行数
	cnt := P.QueryTableNums(tbName)
	pgCount := 0
	if cnt > 0 {
		pgCount = cnt/PgLimit + 1
	}
	//开始处理
	for i := 0; i < pgCount; i++ {
		for {

		}
	}
}

// DumpDatabase 生成Database的备份
func DumpDatabase() {

}
