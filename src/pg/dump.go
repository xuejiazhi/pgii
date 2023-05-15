package pg

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"os"
	"pgii/src/util"
	"strings"
	"time"
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

// DumpSchema dump schema
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

	//创建一个文件夹
	filePath := fmt.Sprintf("dump_schema_%s_%d", P.Schema, time.Now().Unix())
	if err := util.CreateDir(filePath); err != nil {
		util.PrintColorTips(util.LightRed, DumpFailedNoSelectSchema)
		return
	}

	//step1 生成init文件
	initFile := fmt.Sprintf("%s/%s", filePath, INIFile)
	f, _ := os.OpenFile(initFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer fileClose(f)

	//step2 生成schema文件
	scFile := filePath + "/schema.pgi"
	_, _ = f.Write(util.String2Bytes(scFile + "\n"))
	fs, _ := os.OpenFile(scFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer fileClose(fs)

	//生成schema
	scStr := util.String2Bytes(generateSchema(P.Schema, DUMP))
	util.Compress(&scStr)
	_, _ = fs.Write(scStr)
	//print success
	util.PrintColorTips(util.LightGreen, fmt.Sprintf("%s [%s]", DumpSchemaSuccess, P.Schema))

	//生成table
	if tbs, err := P.Tables(""); err == nil {
		if len(tbs) == 0 {
			util.PrintColorTips(util.LightRed, DumpFailedSchemaNoTable)

			return
		}
		//range table
		for _, tb := range tbs {
			tn, ok := tb["tablename"]
			if !ok {
				util.PrintColorTips(util.LightRed, DumpFailedNoTable)
				continue
			}

			//校验表是否存在
			fns, err := splitTableFile(filePath, cast.ToString(tn), SchemaStyle)
			if err != nil {
				util.PrintColorTips(util.LightRed, DumpFailedNoTable+strings.Join(fns, ","))
				continue
			}

			//write fileName
			for _, fn := range fns {
				_, _ = f.Write(util.String2Bytes(fn + "\n"))
			}

			//print tips
			util.PrintColorTips(util.LightGreen, DumpTableSuccess, fmt.Sprintf(" [%v].....", tn))
		}
	}
}

// DumpTable 生成一个创建Table 的SQL
// tbName 表名
func (s *Params) DumpTable() {
	//必须选中模式
	if P.Schema == "" {
		util.PrintColorTips(util.LightRed, DumpFailedNoSelectSchema)
		return
	}

	//取表名
	if "" == s.Param[1] {
		util.PrintColorTips(util.LightRed, DumpFailedNoTable)
		return
	}

	//创建一个文件夹
	filePath := fmt.Sprintf("dump_table_%s_%d", P.Schema, time.Now().Unix())
	if err := util.CreateDir(filePath); err != nil {
		util.PrintColorTips(util.LightRed, DumpFailedNoSelectSchema)
		return
	}

	//step1 生成init文件
	initFile := fmt.Sprintf("%s/%s", filePath, INIFile)
	f, _ := os.OpenFile(initFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer fileClose(f)

	//get tablename
	tbName := cast.ToString(s.Param[1])

	//save file
	fns, err := splitTableFile(filePath, tbName, TableStyle)
	if err != nil {
		util.PrintColorTips(util.LightRed, DumpFailedNoTable+strings.Join(fns, ","))
	}

	//circle write filename
	for _, fn := range fns {
		_, _ = f.Write(util.String2Bytes(fn + "\n"))
	}

	//print tips
	util.PrintColorTips(util.LightGreen, DumpTableSuccess, fmt.Sprintf(" [%s].....", tbName))
}

// DumpDatabase 生成Database的备份
func (s *Params) DumpDatabase() {
	//是否选中了database
	if P.DataBase == "" {
		util.PrintColorTips(util.LightRed, DumpFailedNoSelectDatabase)
		return
	}

	//把schema遍历出来
	scList, err := P.SchemaNS()
	if err != nil || len(scList) == 0 {
		util.PrintColorTips(util.LightRed, DumpDatabaseFailedNoSchema)
		return
	}

	//要生成的文件名
	fileName, err := genDumpFile(DatabaseStyle)
	if err != nil {
		util.PrintColorTips(util.LightRed, DumpFailed)
		return
	}

	util.PrintColorTips(util.LightGreen, ">"+DumpDataBaseBegin)
	//打开要生成的文件句柄
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		util.PrintColorTips(util.LightRed, DumpFailed)
		return
	}
	defer fileClose(f)

	//generate db sql
	dbSQL := []byte(genDataBaseSQL(P.DataBase))

	//压缩数据
	util.Compress(&dbSQL)
	//写入文件
	_, _ = f.Write(dbSQL)
	util.PrintColorTips(util.LightGreen, DumpDataBaseStructSuccess)
	//遍历schema
	for _, m := range scList {
		util.PrintColorTips(util.LightBlue, LineOperate)
		schmaName := cast.ToString(m["nspname"])
		//校验schema 是否存在
		if info, err := P.GetSchemaFromNS(schmaName); err == nil {
			if len(info) == 0 {
				util.PrintColorTips(util.LightRed, ">>"+DumpSchemaNotExists)
			}
		}
		//开始写入schema
		//生成schema
		schemaStr := []byte(generateSchema(schmaName))
		util.Compress(&schemaStr)
		//写入文件
		_, _ = f.Write(schemaStr)
		util.PrintColorTips(util.LightGreen, fmt.Sprintf(">>%s[%s]", DumpSchemaSuccess, schmaName))

		//查询所有的表
		if tbs, err := P.GetTableBySchema(schmaName); err == nil {
			if len(tbs) == 0 {
				util.PrintColorTips(util.LightRed, DumpFailedSchemaNoTable)
				continue
			}

			for _, tb := range tbs {
				tn, ok := tb["tablename"]
				if !ok {
					util.PrintColorTips(util.LightRed, ">>>"+DumpFailedNoTable)
					continue
				}

				tbName := cast.ToString(tn)
				fullTbName := fmt.Sprintf(`"%s".%s`, schmaName, tbName)
				//校验表是否存在
				if tbInfo, err := P.GetTableByName(cast.ToString(tbName)); err != nil || len(tbInfo) == 0 {
					util.PrintColorTips(util.LightRed, ">>>"+DumpFailedNoTable)
					continue
				}

				//生成Table 的DDL
				tbsql := []byte(getTableDdlSql(schmaName, tbName))
				//压缩数据
				util.Compress(&tbsql)
				//写入文件
				_, _ = f.Write(tbsql)
				//print success
				util.PrintColorTips(util.LightGreen, fmt.Sprintf(">>>%s [%s]", DumpTableStructSuccess, cast.ToString(tbName)))

				//处理SQL语句
				//获取表的行数
				cnt := P.QueryTableNums(fullTbName)
				pgCount := 0
				if cnt > 0 {
					pgCount = cnt/PgLimit + 1
				}
				//开始处理表的数据
				//获取表的column
				columnList := P.GetColumnList(schmaName, tbName)
				columnType := P.GetColumnsType(tbName, columnList...)
				for i := 0; i < pgCount; i++ {
					batchSql := ""
					//定义定入的SQL
					batchValue := generateBatchValue(i, fullTbName, columnList, columnType)
					if len(batchValue) > 0 {
						batchSql = fmt.Sprintf(`Insert into "%s".%s(%s) values %s;\n`, schmaName, tbName, strings.Join(columnList, ","), strings.Join(batchValue, ","))
					}
					//压缩数据
					tbSqlByte := []byte(batchSql)
					util.Compress(&tbSqlByte)
					//写入文件
					_, _ = f.Write(tbSqlByte)
				}
				//print success
				util.PrintColorTips(util.LightBlue, fmt.Sprintf(" >>>>%s [%s]", DumpTableRecordSuccess, cast.ToString(fullTbName)))
			}
		}
	}
}

func genDumpFile(style int, param ...string) (fileName string, err error) {
	//要生成的文件名
	switch style {
	case DatabaseStyle:
		fileName = fmt.Sprintf("dump_Database_%s_%d.pgi", P.DataBase, time.Now().Unix())
	case SchemaStyle:
		fileName = fmt.Sprintf("dump_schema_%s_%d.pgi", P.Schema, time.Now().Unix())
	case TableStyle:
		if len(param) > 0 {
			fileName = fmt.Sprintf("dump_table_%s_%d.pgi", cast.ToString(param[0]), time.Now().Unix())
		} else {
			fileName = fmt.Sprintf("dump_table_null_%d.pgi", time.Now().Unix())
		}
	default:
		err = errors.New("dump Style is error")
	}

	if _, err := os.Stat(fileName); err == nil {
		_ = os.Remove(fileName)
	}

	return
}

func genDataBaseSQL(dbName string) (genDBSQL string) {
	return fmt.Sprintf("drop database if exists %s;create database %s;\n", dbName, dbName)
}

// save table file (saveTableFile)
func _(filePath, tbName string) (fileName string, err error) {
	var tbInfo map[string]interface{}
	if tbInfo, err = P.GetTableByName(tbName); err != nil || len(tbInfo) == 0 {
		util.PrintColorTips(util.LightRed, DumpFailedNoTable)
		return
	}

	//打开要生成的文件句柄
	fileName = fmt.Sprintf("%s/%s%s%s", filePath, "tb_", tbName, ".pgi")
	ft, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer fileClose(ft)

	//生成Table 的DDL
	tbSql := getCreateTableSql(tbName)

	//处理SQL语句
	//获取表的行数
	cnt := P.QueryTableNums(fmt.Sprintf(`"%s"."%s"`, P.Schema, tbName))
	pgCount := 0
	if cnt > 0 {
		pgCount = cnt/PgLimit + 1
	}
	//开始处理表的数据
	//获取表的column
	columnList := P.GetColumnList(P.Schema, tbName)
	columnType := P.GetColumnsType(tbName, columnList...)
	for i := 0; i < pgCount; i++ {
		//get batchSQL
		batchSql := ""
		//定义定入的SQL
		batchValue := generateBatchValue(i, fmt.Sprintf(`"%s"."%s"`, P.Schema, tbName), columnList, columnType)
		if len(batchValue) > 0 {
			batchSql = fmt.Sprintf(`Insert into "%s"(%s) values %s;`,
				tbName,
				strings.Join(columnList, ","),
				strings.Join(batchValue, ","))
		}

		//join tbsql
		tbSql = string(append([]byte(tbSql), batchSql...))
	}
	//压缩数据
	tbSqlByte := util.String2Bytes(tbSql)
	util.Compress(&tbSqlByte)
	//写入文件
	_, _ = ft.Write(tbSqlByte)

	//打印
	util.PrintColorTips(util.LightGreen, fmt.Sprintf("Table[%s] %s", tbName, DumpTableSuccess))
	//return
	return
}

// Save Split Table file
func splitTableFile(filePath, tbName string, style int) (fileNames []string, err error) {
	var tbInfo map[string]interface{}
	if tbInfo, err = P.GetTableByName(tbName); err != nil || len(tbInfo) == 0 {
		util.PrintColorTips(util.LightRed, DumpFailedNoTable)
		return
	}

	//打开要生成的文件句柄
	fileName := fmt.Sprintf("%s/%s%s%s", filePath, "tb_", tbName, ".pgi")
	ft, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)

	//生成Table 的DDL
	tbSql := getCreateTableSql(tbName)

	//处理SQL语句
	//获取表的行数
	cnt := P.QueryTableNums(fmt.Sprintf(`"%s"."%s"`, P.Schema, tbName))
	pgCount := 0
	if cnt > 0 {
		pgCount = cnt/PgLimit + 1
	}
	//开始处理表的数据
	//获取表的column
	columnList := P.GetColumnList(P.Schema, tbName)
	columnType := P.GetColumnsType(tbName, columnList...)

	//先获取创建表语句和第一批插入语句
	batchSql := getBatchSql(0, style, tbName, columnList, columnType)
	//join tbsql
	tbSql = string(append([]byte(tbSql), batchSql...))
	//压缩数据
	tbSqlByte := util.String2Bytes(tbSql)
	util.Compress(&tbSqlByte)
	//写入文件
	_, _ = ft.Write(tbSqlByte)
	//加入文件名
	fileNames = append(fileNames, fileName)
	//close ft
	fileClose(ft)

	//The data is larger, so split it
	if pgCount > 1 {
		for i := 1; i < pgCount; i++ {
			//打开要生成的文件句柄
			newFileName := fmt.Sprintf("%s/%s%s_%d.pgi", filePath, "tb_", tbName, i)
			newFt, _ := os.OpenFile(newFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
			newBatchSql := getBatchSql(i, style, tbName, columnList, columnType)
			//压缩数据
			newSqlByte := util.String2Bytes(newBatchSql)
			util.Compress(&newSqlByte)
			//写入文件
			_, _ = newFt.Write(newSqlByte)
			fileClose(newFt)
			//加入文件名
			fileNames = append(fileNames, newFileName)
		}
	}
	//return value
	return
}

// getBatchSql
func getBatchSql(i, style int, tbName string, columnList []string, columnType map[string]string) (batchSql string) {
	//get batchSQL
	//定义定入的SQL
	fullTbName := fmt.Sprintf(`"%s"`, tbName)
	if style == SchemaStyle {
		fullTbName = fmt.Sprintf(`"%s"."%s"`, P.Schema, tbName)
	}
	//get batch value
	batchValue := generateBatchValue(i, fullTbName, columnList, columnType)
	if len(batchValue) > 0 {
		batchSql = fmt.Sprintf(`Insert into %s(%s) values %s;`,
			fullTbName,
			strings.Join(columnList, ","),
			strings.Join(batchValue, ","))
	}
	//返回
	return
}
