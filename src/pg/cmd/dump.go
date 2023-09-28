package cmd

import (
	"fmt"
	"github.com/spf13/cast"
	"os"
	"pgii/src/pg/db"
	"pgii/src/pg/global"
	"pgii/src/util"
	"strings"
	"time"
)

// Dump DUMP PGSQL
func (s *Params) Dump() {
	//参数长度至少为1位
	if len(s.Param) < global.OneCMDLength {
		util.PrintColorTips(util.LightRed, global.DumpFailed)
		return
	}

	//查看Dump的类型
	sCmd := util.TrimLower(s.Param[0])
	switch CheckParamType(sCmd) {
	case global.TableStyle: //生成表的备份
		s.DumpTable()
	case global.SchemaStyle: //生成schema的备份
		s.DumpSchema()
	case global.DatabaseStyle: //生成database的备份
		s.DumpDatabase()
	default:
		util.PrintColorTips(util.LightRed, global.DumpFailed)
	}

}

// DumpSchema dump schema
func (s *Params) DumpSchema() {
	if db.P.Schema == "" {
		util.PrintColorTips(util.LightRed, global.DumpFailedNoSelectSchema)
		return
	}

	//校验schema 是否存在
	if info, err := db.P.GetSchemaFromNS(db.P.Schema); err == nil {
		if len(info) == 0 {
			util.PrintColorTips(util.LightRed, global.DumpFailedNoSelectSchema)
			return
		}
	}

	//创建一个文件夹
	filePath := fmt.Sprintf("dump_schema_%s_%d", db.P.Schema, time.Now().Unix())
	if err := util.CreateDir(filePath); err != nil {
		util.PrintColorTips(util.LightRed, global.DumpFailedNoSelectSchema)
		return
	}

	//step1 生成init文件
	initFile := fmt.Sprintf("%s/%s", filePath, global.INIFile)
	f, _ := os.OpenFile(initFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer fileClose(f)

	//step2 生成schema文件
	scFile := filePath + "/schema.pgi"
	_, _ = f.Write(util.ZeroCopyByte(scFile + "\n"))
	fs, _ := os.OpenFile(scFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer fileClose(fs)

	//生成schema
	scStr := util.ZeroCopyByte(db.GenerateSchema(db.P.Schema, global.DUMP))
	util.Compress(&scStr)
	_, _ = fs.Write(scStr)
	//print success
	util.PrintColorTips(util.LightGreen, fmt.Sprintf("%s [%s]", global.DumpSchemaSuccess, db.P.Schema))

	//生成table
	if tbs, err := db.P.Tables(""); err == nil {
		if len(tbs) == 0 {
			util.PrintColorTips(util.LightRed, global.DumpFailedSchemaNoTable)

			return
		}
		//range table
		for _, tb := range tbs {
			tn, ok := tb["tablename"]
			if !ok {
				util.PrintColorTips(util.LightRed, global.DumpFailedNoTable)
				continue
			}

			//校验表是否存在
			fns, err := splitTableFile(filePath, cast.ToString(tn), global.SchemaStyle)
			if err != nil {
				util.PrintColorTips(util.LightRed, global.DumpFailedNoTable+strings.Join(fns, ","))
				continue
			}

			//write fileName
			for _, fn := range fns {
				_, _ = f.Write(util.ZeroCopyByte(fn + "\n"))
			}

			//print tips
			util.PrintColorTips(util.LightGreen, global.DumpTableSuccess, fmt.Sprintf(" [%v].....", tn))
		}
	}
}

// DumpTable 生成一个创建Table 的SQL
// tbName 表名
func (s *Params) DumpTable() {
	//必须选中模式
	if db.P.Schema == "" {
		util.PrintColorTips(util.LightRed, global.DumpFailedNoSelectSchema)
		return
	}

	//取表名
	if "" == s.Param[1] {
		util.PrintColorTips(util.LightRed, global.DumpFailedNoTable)
		return
	}

	//创建一个文件夹
	filePath := fmt.Sprintf("dump_table_%s_%d", db.P.Schema, time.Now().Unix())
	if err := util.CreateDir(filePath); err != nil {
		util.PrintColorTips(util.LightRed, global.DumpFailedNoSelectSchema)
		return
	}

	//step1 生成init文件
	initFile := fmt.Sprintf("%s/%s", filePath, global.INIFile)
	f, _ := os.OpenFile(initFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer fileClose(f)

	//get tablename
	tbName := cast.ToString(s.Param[1])

	//save file
	fns, err := splitTableFile(filePath, tbName, global.TableStyle)
	if err != nil {
		util.PrintColorTips(util.LightRed, global.DumpFailedNoTable+strings.Join(fns, ","))
	}

	//circle write filename
	for _, fn := range fns {
		_, _ = f.Write(util.ZeroCopyByte(fn + "\n"))
	}

	//print tips
	util.PrintColorTips(util.LightGreen, global.DumpTableSuccess, fmt.Sprintf(" [%s].....", tbName))
}

// DumpDatabase 生成Database的备份
func (s *Params) DumpDatabase() {
	//是否选中了database
	if db.P.DataBase == "" {
		util.PrintColorTips(util.LightRed, global.DumpFailedNoSelectDatabase)
		return
	}

	//把schema遍历出来
	scList, err := db.P.SchemaNS()
	if err != nil || len(scList) == 0 {
		util.PrintColorTips(util.LightRed, global.DumpDatabaseFailedNoSchema)
		return
	}

	//要生成的文件名
	fileName, err := genDumpFile(global.DatabaseStyle)
	if err != nil {
		util.PrintColorTips(util.LightRed, global.DumpFailed)
		return
	}

	util.PrintColorTips(util.LightGreen, ">"+global.DumpDataBaseBegin)
	//打开要生成的文件句柄
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		util.PrintColorTips(util.LightRed, global.DumpFailed)
		return
	}
	defer fileClose(f)

	//generate db sql
	dbSQL := util.ZeroCopyByte(genDataBaseSQL(db.P.DataBase))

	//压缩数据
	util.Compress(&dbSQL)
	//写入文件
	_, _ = f.Write(dbSQL)
	util.PrintColorTips(util.LightGreen, global.DumpDataBaseStructSuccess)
	//遍历schema
	for _, m := range scList {
		util.PrintColorTips(util.LightBlue, global.LineOperate)
		schemaName := cast.ToString(m["nspname"])
		//校验schema 是否存在
		if info, err := db.P.GetSchemaFromNS(schemaName); err == nil {
			if len(info) == 0 {
				util.PrintColorTips(util.LightRed, ">>"+global.DumpSchemaNotExists)
			}
		}
		//开始写入schema
		//生成schema
		schemaStr := util.ZeroCopyByte(db.GenerateSchema(schemaName))
		util.Compress(&schemaStr)
		//写入文件
		_, _ = f.Write(schemaStr)
		util.PrintColorTips(util.LightGreen, fmt.Sprintf(">>%s[%s]", global.DumpSchemaSuccess, schemaName))

		//查询所有的表
		if tbs, err := db.P.GetTableBySchema(schemaName); err == nil {
			if len(tbs) == 0 {
				util.PrintColorTips(util.LightRed, global.DumpFailedSchemaNoTable)
				continue
			}

			for _, tb := range tbs {
				tn, ok := tb["tablename"]
				if !ok {
					util.PrintColorTips(util.LightRed, ">>>"+global.DumpFailedNoTable)
					continue
				}

				tbName := cast.ToString(tn)
				fullTbName := fmt.Sprintf(`"%s".%s`, schemaName, tbName)
				//校验表是否存在
				if tbInfo, err := db.P.GetTableByName(cast.ToString(tbName)); err != nil || len(tbInfo) == 0 {
					util.PrintColorTips(util.LightRed, ">>>"+global.DumpFailedNoTable)
					continue
				}

				//生成Table 的DDL
				tbsql := util.ZeroCopyByte(getTableDdlSql(schemaName, tbName))
				//压缩数据
				util.Compress(&tbsql)
				//写入文件
				_, _ = f.Write(tbsql)
				//print success
				util.PrintColorTips(util.LightGreen, fmt.Sprintf(">>>%s [%s]", global.DumpTableStructSuccess, cast.ToString(tbName)))

				//处理SQL语句
				//获取表的行数
				cnt := db.P.QueryTableNums(fullTbName)
				pgCount := 0
				if cnt > 0 {
					pgCount = cnt/global.PgLimit + 1
				}
				//开始处理表的数据
				//获取表的column
				columnList := db.P.GetColumnList(schemaName, tbName)
				columnType := db.P.GetColumnsType(tbName, columnList...)
				for i := 0; i < pgCount; i++ {
					batchSql := ""
					//定义定入的SQL
					batchValue := db.GenerateBatchValue(i, fullTbName, columnList, columnType)
					if len(batchValue) > 0 {
						batchSql = fmt.Sprintf(`Insert into "%s".%s(%s) values %s;\n`, schemaName, tbName, strings.Join(columnList, ","), strings.Join(batchValue, ","))
					}
					//压缩数据
					tbSqlByte := util.ZeroCopyByte(batchSql)
					util.Compress(&tbSqlByte)
					//写入文件
					_, _ = f.Write(tbSqlByte)
				}
				//print success
				util.PrintColorTips(util.LightBlue, fmt.Sprintf(" >>>>%s [%s]", global.DumpTableRecordSuccess, cast.ToString(fullTbName)))
			}
		}
	}
}
