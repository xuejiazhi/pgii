package pg

import (
	"fmt"
	"github.com/spf13/cast"
	"pgii/src/util"
	"strings"
)

func (s *Params) DDL() {
	if len(s.Param) != TwoCMDLength {
		fmt.Println("Failed:DDL Cmd fail")
		return
	}

	//查看DDL的类型
	sCmd := util.TrimLower(s.Param[0])
	name := util.TrimLower(s.Param[1])
	switch CheckParamType(sCmd) {
	case SchemaStyle: //查看schema的DDL
		s.DDLSchema(name)
	case TableStyle: //查看table的DDL
		s.DDLTable(name)
	case ViewStyle: //查看view的DDL
		s.DDLView(name)
	default:
		fmt.Println("Failed:DDL Cmd fail")
		return
	}
}

// DDLSchema 生成SCHEAM的DDL
func (s *Params) DDLSchema(name string) {
	//校验schema 是否存在
	info, err := P.GetSchemaFromNS(name)
	if err != nil {
		util.PrintColorTips(util.LightRed, DDLSchemaError, err.Error())
		return
	}

	if len(info) == 0 {
		util.PrintColorTips(util.LightRed, DDLSchemaNotExists)
		return
	}

	//print schema ddl
	fmt.Println(generateSchema(name))
}

// DDLTable 生成Table的DDL
func (s *Params) DDLTable(name string) {
	//get table info
	info, err := P.GetTableByName(name)
	if err != nil {
		util.PrintColorTips(util.LightRed, DDLTableError, err.Error())
		return
	}

	if len(info) == 0 {
		util.PrintColorTips(util.LightRed, DDLTableNoExists)
		return
	}

	//print
	util.PrintColorTips(util.LightSeaBlue, getTableDdlSql(name))
}

// DDLView 生成view视图的DDL
// viewName 视图名称
func (s *Params) DDLView(viewName string) {
	//获取view信息
	viewInfo, err := P.Views("filter", viewName)
	if err != nil {
		util.PrintColorTips(util.LightRed, DDLViewError, err.Error())
		return
	}

	//判断信息不为空
	if len(viewInfo) == 0 {
		util.PrintColorTips(util.LightRed, DDLViewNoExists)
		return
	}

	//print Create View SQL
	fmt.Println("========= Create View Success ============")
	if def, ok := viewInfo[0]["definition"]; ok {
		fmt.Println(fmt.Sprintf(" CREATE OR REPLACE VIEW \"%s\".%s\n AS%s", P.Schema, viewName, def))
	}
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
			if util.InArray(ud, []string{"int2", "int4", "int8"}) {
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
		case "int2":
			return "smallserial"
		case "int4":
			return "serial4"
		case "int8":
			return "bigserial"
		default:
			return types
		}
	}
	return types
}

// 获取DDL T-SQL
func getTableDdlSql(tbName string) (sqlStr string) {
	//print Create Table SQL
	sqlStr = fmt.Sprintf(`========= Create Table Success ============
-- DROP Table;
DROP Table %s;`, tbName) + "\n"

	//获取column
	column, err := P.Column(tbName)
	if err != nil {
		fmt.Println(util.SetColor(DDLColumnNoExists, util.LightRed))
		return
	}

	columnList := generateColumn(column)

	//获取主键或唯一约束
	puInfo, err := P.GetPriMaryUniqueKey(tbName)
	constraintList, indexList := generateConstraint(puInfo)
	if err == nil {
		//加上约束
		columnList = append(columnList, constraintList...)
	}

	//获取建表语句
	sqlStr += fmt.Sprintf("CREATE TABLE \"%s\".%s (\n%s\n);\n",
		P.Schema,
		tbName,
		strings.Join(columnList, ",\n"))

	//获取Index
	indexDef, err := P.GetIndexDef(tbName, indexList)
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

func getTriggerDef(tbName string) (triggerDef string, err error) {
	//判断是否有触发器
	pgcData, err := P.GetPgClassValueForTbName(tbName, "relhastriggers", "oid")
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
	triggerData, err := P.GetTriggerByTgRelid(cast.ToInt(oid))
	if err != nil || len(triggerData) == 0 {
		return
	}

	// trigger oid是否存在
	_, ok := triggerData[0]["oid"]
	if !ok {
		return
	}

	//查询triggerdef
	defInfo, err := P.GetPgTriggerDef(cast.ToInt(triggerData[0]["oid"]))
	if err != nil || len(defInfo) == 0 {
		return
	}

	if _, ok := defInfo["def"]; ok {
		triggerDef = cast.ToString(defInfo["def"])
	}
	return
}
