package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"os"
	"pgii/src/pg/db"
	"pgii/src/pg/global"
	"pgii/src/util"
	"strings"
	"time"
)

// CheckParamType 检查传过来的参数
func CheckParamType(types string) int {
	switch types {
	case "database", "db": //数据库
		return global.DatabaseStyle
	case "table", "tb": //表
		return global.TableStyle
	case "index", "idx": //索引
		return global.IndexStyle
	case "view", "vw": //视图
		return global.ViewStyle
	case "sd", "selectdb":
		return global.SelectStyle
	case "sc", "schema":
		return global.SchemaStyle
	case "tg", "trigger":
		return global.TriggerStyle
	case "ver", "version":
		return global.VersionStyle
	case "conn", "connection":
		return global.ConnectionStyle
	case "proc", "process": //当前进程
		return global.ProcessStyle
	case "tablespace", "tbsp": //表空间
		return global.TableSpaceStyle
	default:
		return global.NoneStyle
	}
}

// Handler ------func handle---------------//
func Handler(param ...string) HandlerInterface {
	return HandlerInterface(getInstance(param...))
}

func getInstance(param ...string) *Params {
	return &Params{
		Param: param,
	}
}

func (s *Params) judgeTable() (tbParam []string, errorMsg string, err error) {
	//赋值
	tbParam = s.Param[1:]
	//必须要指定表
	if len(tbParam) == global.ZeroCMDLength {
		errorMsg = global.CmdLineWrong
		err = errors.New(errorMsg)
		return
	}

	//判断指定的sc是否存在
	scInfo, err := db.P.GetSchemaFromNS(db.P.Schema)
	if err != nil {
		errorMsg = fmt.Sprintf("%s %s", global.SizeFailedNoSchema, err.Error())
		return
	}

	if len(scInfo) == global.ZeroCMDLength {
		errorMsg = global.SizeFailedNoSchema
		err = errors.New(errorMsg)
		return
	}
	//判断table是否存在
	tbInfo, err := db.P.GetTableByName(tbParam[0])
	if err != nil {
		errorMsg = fmt.Sprintf("%s %s", global.SizeFailedNoTable, err.Error())
		return
	}

	if len(tbInfo) == global.ZeroCMDLength {
		errorMsg = global.SizeFailedNoTable
		err = errors.New(errorMsg)
		return
	}

	if errorMsg != "" {
		err = errors.New(errorMsg)
	}
	//
	return
}

func (s *Params) judgeTableSpace() (tbParam []string, errorMsg string, err error) {
	//赋值
	tbParam = s.Param[1:]
	//必须要指定表
	if len(tbParam) == global.ZeroCMDLength {
		errorMsg = global.SizeFailedPointTable
	}
	//判断table是否存在
	tbInfo, err := db.P.GetTableSpaceNameBySpcName(tbParam[0])
	if err != nil {
		errorMsg = fmt.Sprintf("%s %s", global.SizeFailedNoTable, err.Error())
	}
	if len(tbInfo) == global.ZeroCMDLength {
		errorMsg = global.SizeFailedNoTable
	}

	if errorMsg != "" {
		err = errors.New(errorMsg)
	}
	//
	return
}

func fileClose(f *os.File) {
	err := f.Close()
	if err != nil {
		util.PrintColorTips(util.LightRed, global.CloseFileFailed)
	}
}

// ---------------------------Operate SQL Func-------------------//
// exec sql
func execUnzipSQL(fileName string) (err error) {
	//Open File
	ft, err := os.Open(fileName)
	defer ft.Close()

	//judge
	if err != nil {
		util.PrintColorTips(util.LightRed, global.LoadNoFile, fileName, err.Error())
		return
	}

	//read content
	content, err := ioutil.ReadAll(ft)

	//解压
	unZipSQL, err := util.UnCompress(content)
	affect, err := db.P.ExecSQL(string(unZipSQL))
	if err != nil {
		util.PrintColorTips(util.LightRed, global.LoadTableExecSQLFailed, err.Error())
		return
	}

	//print Success
	util.PrintColorTips(util.LightGreen, global.LoadTableSQLSuccess, fmt.Sprintf(" [%s] Affect Nums:%d", fileName, affect))

	//return
	return
}

func generateColumn(column []map[string]interface{}) (columnList []string) {
	///字段名说明
	//column_name              //字段名
	//udt_name                 //类型
	//character_maximum_length //长度
	//is_nullable              //YES 可以为空，NO 不能为空
	//identity_generation      //GENERATED <identity_generation> AS IDENTITY
	//column_default  默认值
	//遍历字段列表
	for _, v := range column {
		st := ""
		//字段类型
		if c, ok := v["column_name"]; ok {
			st += "    " + cast.ToString(c)
		}

		//类型
		tb := strings.ToLower(cast.ToString(v["table_name"]))
		cn := strings.ToLower(cast.ToString(v["column_name"]))
		nextVal := fmt.Sprintf("nextval('%s_%s_seq'::regclass)", tb, cn)
		if u, ok := v["udt_name"]; ok {
			ud := strings.ToLower(cast.ToString(u))
			if util.InArray(ud, global.Int2Type, global.Int4Type, global.Int8Type) {
				st += fmt.Sprintf(" %s", getSerial(nextVal, ud, v["column_default"]))
			} else {
				st += fmt.Sprintf(" %s", ud)
			}
		}

		//长度
		if cml, ok := v["character_maximum_length"]; ok {
			if cml != nil {
				st += fmt.Sprintf("(%s)", cast.ToString(cml))
			}
		}

		//是否为空
		if nan, ok := v["is_nullable"]; ok {
			nanStr := strings.ToLower(cast.ToString(nan))
			if nanStr == "yes" {
				st += " NULL"
			} else {
				st += " NOT NULL"
			}
		}

		//identity_generation
		if ig, ok := v["identity_generation"]; ok {
			if ig != nil {
				st += fmt.Sprintf(" GENERATED %s AS IDENTITY", cast.ToString(ig))
			}
		}

		//column_default
		if cd, ok := v["column_default"]; ok {
			if cd != nil {
				if cast.ToString(cd) != nextVal {
					st += fmt.Sprintf(" DEFAULT %s", cast.ToString(cd))
				}
			}
		}
		//将字符串加入
		columnList = append(columnList, st)
	}
	return
}

func generateConstraint(puInfo []map[string]interface{}) (constraintList, indexList []string) {
	//
	//  column_name 约束字段
	//	constraint_name 约束名称
	//	constraint_type 约束类型
	daList := map[string][]string{}
	for _, v := range puInfo {
		//取值
		cln, okl := v["column_name"]
		csn, oks := v["constraint_name"]
		cst, okt := v["constraint_type"]
		if okl && oks && okt {
			key := cast.ToString(csn) + "|" + cast.ToString(cst)
			daList[key] = append(daList[key], cast.ToString(cln))
			indexList = append(indexList, cast.ToString(csn))
		}
	}

	if len(daList) > 0 {
		for k, v := range daList {
			kc := strings.Split(k, "|")
			if len(kc) == 2 {
				puStr := fmt.Sprintf("    CONSTRAINT %s %s (%s)", kc[0], kc[1], strings.Join(v, ","))
				constraintList = append(constraintList, puStr)
			}
		}
	}

	return
}

func getSerial(nextval, types string, dv interface{}) string {
	//为空直接返回原来的类型
	if dv == nil {
		return types
	}

	//校验是否是serial类型
	if cast.ToString(dv) == nextval {
		switch types {
		case global.Int2Type:
			return "smallserial"
		case global.Int4Type:
			return "serial4"
		case global.Int8Type:
			return "bigserial"
		default:
			return types
		}
	}
	return types
}

// 获取DDL T-SQL
func getTableDdlSql(schema, tbName string) (sqlStr string) {
	//schema 不能为空
	if schema == "" {
		schema = db.P.Schema
	}
	//print Create Table SQL
	sqlStr = fmt.Sprintf(`-- ========= Create Table Success ============
-- DROP Table;
DROP table  IF exists "%s"."%s" cascade;`, schema, tbName) + "\n"

	//获取column
	column, err := db.P.Column(schema, tbName)
	if err != nil {
		fmt.Println(util.SetColor(global.DDLColumnNoExists, util.LightRed))
		return
	}

	columnList := generateColumn(column)

	//获取主键或唯一约束
	puInfo, err := db.P.GetPriMaryUniqueKey(tbName)
	constraintList, indexList := generateConstraint(puInfo)
	if err == nil {
		//加上约束
		columnList = append(columnList, constraintList...)
	}

	//获取建表语句
	sqlStr += fmt.Sprintf("CREATE TABLE \"%s\".%s (\n%s\n);\n",
		schema,
		tbName,
		strings.Join(columnList, ",\n"))

	//获取Index
	indexDef, err := db.P.GetIndexDef(tbName, indexList)
	if err == nil {
		for k, v := range indexDef {
			if def, ok := v["indexdef"]; ok {
				sqlStr += fmt.Sprintf("%v;", def)
				if k < len(indexDef)-1 {
					sqlStr += "\n"
				}
			}
		}
	}
	//判断是否有触发器
	if triggerDef, err := getTriggerDef(tbName); err == nil {
		if triggerDef != "" {
			sqlStr += "\n" + triggerDef + ";\n"
		}
	}

	return
}

// 获取Create SQL
func getCreateTableSql(tbName string) (sqlStr string) {
	//print Create Table SQL
	sqlStr = fmt.Sprintf(`DROP table  IF exists "%s" cascade;`, tbName) + "\n"

	//获取column
	column, err := db.P.Column(db.P.Schema, tbName)
	if err != nil {
		fmt.Println(util.SetColor(global.DDLColumnNoExists, util.LightRed))
		return
	}

	columnList := generateColumn(column)

	//获取主键或唯一约束
	puInfo, err := db.P.GetPriMaryUniqueKey(tbName)
	constraintList, indexList := generateConstraint(puInfo)
	if err == nil {
		//加上约束
		columnList = append(columnList, constraintList...)
	}

	//获取建表语句
	sqlStr += fmt.Sprintf(`CREATE TABLE "%s"(%s);`,
		tbName,
		strings.Join(columnList, ","))

	replySchema := fmt.Sprintf("\"%s\".", db.P.Schema)
	//获取Index
	indexDef, err := db.P.GetIndexDef(tbName, indexList)
	if err == nil {
		for _, v := range indexDef {
			if def, ok := v["indexdef"]; ok {
				defStr := strings.Replace(cast.ToString(def), replySchema, "", -1)
				sqlStr += fmt.Sprintf("%v;", defStr)
			}
		}
	}
	//判断是否有触发器
	if triggerDef, err := getTriggerDef(tbName); err == nil {
		if triggerDef != "" {
			sqlStr += strings.Replace(triggerDef, replySchema, "", -1) + ";"
		}
	}

	return
}

func getTriggerDef(tbName string) (triggerDef string, err error) {
	//判断是否有触发器
	pgcData, err := db.P.GetPgClassValueForTbName(tbName, "relhastriggers", "oid")
	if err != nil {
		return
	}

	//judge field
	hasTrigger, okr := pgcData["relhastriggers"]
	oid, oko := pgcData["oid"]

	if !okr || !oko {
		return
	}

	//是否存的触发器
	if !cast.ToBool(hasTrigger) {
		return
	}

	//查询trigger
	triggerData, err := db.P.GetTriggerByTgRelid(cast.ToInt(oid))
	if err != nil || len(triggerData) == 0 {
		return
	}

	// trigger oid是否存在
	_, ok := triggerData[0]["oid"]
	if !ok {
		return
	}

	//查询triggerdef
	defInfo, err := db.P.GetPgTriggerDef(cast.ToInt(triggerData[0]["oid"]))
	if err != nil || len(defInfo) == 0 {
		return
	}

	if _, ok := defInfo["def"]; ok {
		triggerDef = cast.ToString(defInfo["def"])
	}
	return
}

func genDumpFile(style int, param ...string) (fileName string, err error) {
	//要生成的文件名
	switch style {
	case global.DatabaseStyle:
		fileName = fmt.Sprintf("dump_Database_%s_%d.pgi", db.P.DataBase, time.Now().Unix())
	case global.SchemaStyle:
		fileName = fmt.Sprintf("dump_schema_%s_%d.pgi", db.P.Schema, time.Now().Unix())
	case global.TableStyle:
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
	if tbInfo, err = db.P.GetTableByName(tbName); err != nil || len(tbInfo) == 0 {
		util.PrintColorTips(util.LightRed, global.DumpFailedNoTable)
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
	cnt := db.P.QueryTableNums(fmt.Sprintf(`"%s"."%s"`, db.P.Schema, tbName))
	pgCount := 0
	if cnt > 0 {
		pgCount = cnt/global.PgLimit + 1
	}
	//开始处理表的数据
	//获取表的column
	columnList := db.P.GetColumnList(db.P.Schema, tbName)
	columnType := db.P.GetColumnsType(tbName, columnList...)
	for i := 0; i < pgCount; i++ {
		//get batchSQL
		batchSql := ""
		//定义定入的SQL
		batchValue := db.GenerateBatchValue(i, fmt.Sprintf(`"%s"."%s"`, db.P.Schema, tbName), columnList, columnType)
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
	util.PrintColorTips(util.LightGreen, fmt.Sprintf("Table[%s] %s", tbName, global.DumpTableSuccess))
	//return
	return
}

// Save Split Table file
func splitTableFile(filePath, tbName string, style int) (fileNames []string, err error) {
	var tbInfo map[string]interface{}
	if tbInfo, err = db.P.GetTableByName(tbName); err != nil || len(tbInfo) == 0 {
		util.PrintColorTips(util.LightRed, global.DumpFailedNoTable)
		return
	}

	//打开要生成的文件句柄
	fileName := fmt.Sprintf("%s/%s%s%s", filePath, "tb_", tbName, ".pgi")
	ft, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)

	//生成Table 的DDL
	tbSql := getCreateTableSql(tbName)

	//处理SQL语句
	//获取表的行数
	cnt := db.P.QueryTableNums(fmt.Sprintf(`"%s"."%s"`, db.P.Schema, tbName))
	pgCount := 0
	if cnt > 0 {
		pgCount = cnt/global.PgLimit + 1
	}
	//开始处理表的数据
	//获取表的column
	columnList := db.P.GetColumnList(db.P.Schema, tbName)
	columnType := db.P.GetColumnsType(tbName, columnList...)

	//先获取创建表语句和第一批插入语句
	batchSql := db.GenerateBatchSql(0, style, tbName, columnList, columnType)
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
			newBatchSql := db.GenerateBatchSql(i, style, tbName, columnList, columnType)
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

func _judgeCommon(tbParam []string) (errorMsg string) {
	//必须要指定表
	if len(tbParam) == global.ZeroCMDLength {
		errorMsg = global.SizeFailedPointTable
	}

	//判断指定的sc是否存在
	scInfo, err := db.P.GetSchemaFromNS(db.P.Schema)
	if err != nil {
		errorMsg = fmt.Sprintf("%s %s", global.SizeFailedNoSchema, err.Error())
	}
	if len(scInfo) == global.ZeroCMDLength {
		errorMsg = global.SizeFailedNoSchema
	}
	return
}
