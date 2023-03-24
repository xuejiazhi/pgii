package pg

import (
	"fmt"
	"github.com/spf13/cast"
	"os"
	"pgii/src/util"
	"strings"
)

// Dump DUMP PGSQL
func (s *Params) Dump() {
	//参数长度至少为1位
	if len(s.Param) < OneCMDLength {
		util.PrintColorTips(util.LightRed, DumpFailed)
		return
	}

	//查看Dump的类型
	sCmd := util.TrimLower(s.Param[0])
	switch CheckParamType(sCmd) {
	case TableStyle: //生成表的备份
		s.DumpTable()
	case SchemaStyle: //生成schema的备份
		s.DumpSchema()
	case DatabaseStyle: //生成database的备份
		s.DumpDatabase()
	default:
		util.PrintColorTips(util.LightRed, DumpFailed)
	}

}

func (s *Params) DumpSchema() {
	if P.Schema == "" {
		util.PrintColorTips(util.LightRed, DumpFailedNoSelectSchema)
		return
	}

	//校验schema 是否存在
	if info, err := P.GetSchemaFromNS(P.Schema); err == nil {
		if len(info) == 0 {
			util.PrintColorTips(util.LightRed, DumpFailedNoSelectSchema)
			return
		}
	}

	//打开要生成的文件句柄
	fileName := fmt.Sprintf("dump_schema_%s_%s.pgi", P.Schema, util.GetFormatDateTime())
	if _, err := os.Stat(fileName); err == nil {
		_ = os.Remove(fileName)
	}

	f, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer f.Close()
	//生成schema
	schemaStr := []byte(generateSchema(P.Schema))
	util.Compress(&schemaStr)
	//写入文件
	_, _ = f.Write(schemaStr)
	//print success
	util.PrintColorTips(util.LightGreen, fmt.Sprintf("%s [%s]", DumpSchemaSuccess, P.Schema))
	//查询所有的表
	if tbs, err := P.Tables(""); err == nil {
		if len(tbs) == 0 {
			util.PrintColorTips(util.LightRed, DumpFailedSchemaNoTable)
			return
		}

		for _, tb := range tbs {
			tn, ok := tb["tablename"]
			if !ok {
				util.PrintColorTips(util.LightRed, DumpFailedNoTable)
				continue
			}

			tbName := cast.ToString(tn)
			//校验表是否存在
			if tbInfo, err := P.GetTableByName(cast.ToString(tbName)); err != nil || len(tbInfo) == 0 {
				util.PrintColorTips(util.LightRed, DumpFailedNoTable)
				return
			}

			//生成Table 的DDL
			tbsql := []byte(getTableDdlSql(cast.ToString(tbName)))
			//压缩数据
			util.Compress(&tbsql)
			//写入文件
			_, _ = f.Write(tbsql)
			//print success
			util.PrintColorTips(util.LightGreen, fmt.Sprintf("%s [%s]", DumpTableStructSuccess, cast.ToString(tbName)))

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
			//print success
			util.PrintColorTips(util.LightBlue, fmt.Sprintf(" ->%s [%s]", DumpTableRecordSuccess, cast.ToString(tbName)))
		}
	}
}

// DumpTable 生成一个创建Table 的SQL
// tbName 表名
func (s *Params) DumpTable() {
	//取表名
	tbName := ""
	if "" == s.Param[1] {
		util.PrintColorTips(util.LightRed, DumpFailedNoTable)
		return
	} else {
		tbName = s.Param[1]
	}

	//校验表是否存在
	if tbInfo, err := P.GetTableByName(tbName); err != nil || len(tbInfo) == 0 {
		util.PrintColorTips(util.LightRed, DumpFailedNoTable)
		return
	}

	//要生成的文件名
	fileName := fmt.Sprintf("dump_table_%s_%s.pgi", tbName, util.GetFormatDateTime())
	//是否存在文件
	if _, err := os.Stat(fileName); err == nil {
		_ = os.Remove(fileName)
	}

	//打开要生成的文件句柄
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
	//打印
	util.PrintColorTips(util.LightGreen, DumpTableSuccess)
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
func (s *Params) DumpDatabase() {

}
