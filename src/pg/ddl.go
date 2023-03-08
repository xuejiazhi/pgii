package pg

import (
	"fmt"
	"github.com/spf13/cast"
	"pgii/src/util"
	"strings"
)

func DDL(cms ...string) {
	if len(cms) != 2 {
		fmt.Println("Failed:DDL Cmd fail")
		return
	}

	//查看DDL的类型
	types := util.TrimLower(cms[0])
	name := util.TrimLower(cms[1])
	switch types {
	case "sc", "schema": //查看schema的DDL
		DDLSchema(name)
	case "tb", "table": //查看table的DDL
		DDLTable(name)
	case "vw", "view": //查看view的DDL
		DDLView(name)
	default:
		fmt.Println("Failed:DDL Cmd fail")
		return
	}
}

// DDLSchema 生成SCHEAM的DDL
func DDLSchema(name string) {
	//校验schema 是否存在
	if info, err := P.GetSchemaFromNS(name); err == nil {
		if len(info) == 0 {
			fmt.Println("Failed:DDL Cmd Schema fail,Schema not exists!")
			return
		}
		//print Create schema SQL
		fmt.Println("========= Create Schema Success ============")
		fmt.Println(fmt.Sprintf("-- DROP SCHEMA %s;", name))
		fmt.Println(fmt.Sprintf("CREATE SCHEMA \"%s\" AUTHORIZATION %s;", name, *UserName))
	} else {
		fmt.Println("Failed:DDL Cmd Schema fail,error ", err.Error())
		return
	}
}

// DDLTable 生成Table的DDL
func DDLTable(name string) {
	if info, err := P.GetTableByName(name); err == nil {
		if len(info) == 0 {
			fmt.Println("Failed:DDL Cmd Table fail,Table not exists!")
			return
		}
		//print Create Table SQL
		fmt.Println("========= Create Table Success ============")
		fmt.Println(fmt.Sprintf("-- DROP Table;"))
		fmt.Println(fmt.Sprintf("-- DROP Table %s;", name))

		//获取column
		if column, err := P.Column(name); err == nil {
			columnList := generateColumn(column)

			//获取主键或唯一约束
			puInfo, err := P.GetPriMaryUniqueKey(name)
			constraintList, indexList := generateConstraint(puInfo)
			if err == nil {
				//加上约束
				columnList = append(columnList, constraintList...)
			}

			//打印建表语句
			fmt.Println(fmt.Sprintf("CREATE TABLE \"%s\".%s (\n%s\n);",
				P.Schema,
				name,
				strings.Join(columnList, ",\n")))

			//获取Index
			indexDef, err := P.GetIndexDef(name, indexList)
			if err == nil {
				for _, v := range indexDef {
					if def, ok := v["indexdef"]; ok {
						fmt.Println(fmt.Sprintf("%v;", def))
					}
				}
			}
		}
	} else {
		fmt.Println("Failed:DDL Cmd Schema fail,error ", err.Error())
		return
	}
}

// DDLView 生成view视图的DDL
// viewName 视图名称
func DDLView(viewName string) {
	if viewInfo, err := P.Views("filter", viewName); err == nil {
		//判断信息不为空
		if len(viewInfo) == 0 {
			fmt.Println("Failed:DDL Cmd View fail,View not exists!")
			return
		}

		//print Create View SQL
		fmt.Println("========= Create View Success ============")
		if def, ok := viewInfo[0]["definition"]; ok {
			fmt.Println(fmt.Sprintf(" CREATE OR REPLACE VIEW \"%s\".%s\n AS%s", P.Schema, viewName, def))
		}
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
