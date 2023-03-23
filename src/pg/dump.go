package pg

import (
	"fmt"
	"github.com/spf13/cast"
	"os"
	"pgii/src/util"
	"strings"
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
	fileName := fmt.Sprintf("dump_table_%s.pgi", tbName)
	if _, err := os.Stat(fileName); err == nil {
		_ = os.Remove(fileName)
	}
	f, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
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
	//开始处理表的数据
	//获取表的column
	columnList := P.GetColumnList(tbName)
	columnType := P.GetColumnsType(tbName, columnList...)
	for i := 0; i < pgCount; i++ {
		batchSql := ""
		//定义定入的SQL
		batchValue := generateBatchValue(i, tbName, columnList, columnType)
		if len(batchValue) > 0 {
			batchSql = fmt.Sprintf("Insert into %s.%s(%s) values %s;", P.Schema, tbName, strings.Join(columnList, ","), strings.Join(batchValue, ","))
		}
		//压缩数据
		tbSqlByte := []byte(batchSql)
		util.Compress(&tbSqlByte)
		//写入文件
		_, _ = f.Write(tbSqlByte)
	}
	fmt.Println(util.SetColor(DumpTableSuccess, util.LightGreen))
}

func generateBatchValue(idx int, tbName string, columnList []string, columnType map[string]string) (batchValue []string) {
	//查询的SQL
	querySQL := P.GetQuerySql(tbName, columnList, columnType, idx)
	//run sql
	data, err := P.RunSQL(querySQL)
	if err != nil {
		fmt.Println("value is null")
		return
	}

	if len(data) == 0 {
		fmt.Println("value is null")
		return
	}
	//循环
	for _, v := range data {
		valSon := []string{}
		l := 0
		for _, sv := range columnList {
			if _, ok := v[sv]; ok {
				valSon = append(valSon, fmt.Sprintf("'%s'", cast.ToString(v[sv])))
				l++
			} else {
				break
			}
		}
		//加入数组
		batchValue = append(batchValue, "("+strings.Join(valSon, ",")+")")
	}
	return
}

// 获取
func getColumn(columnList []map[string]interface{}) (cols []string) {
	if len(columnList) == 0 {
		return
	}

	//拼接column
	for _, c := range columnList {
		if _, ok := c["column_name"]; ok {
			cols = append(cols, cast.ToString(c["column_name"]))
		}
	}
	return
}

// DumpDatabase 生成Database的备份
func DumpDatabase() {

}
