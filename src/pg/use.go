package pg

import (
	"pgii/src/util"
	"strings"
)

// Use 用做切换数据库
func (s *Params) Use() {
	//判断
	if len(s.Param) == TwoCMDLength {
		sw := strings.Trim(s.Param[0], "")
		para := strings.Trim(s.Param[1], "")
		switch CheckParamType(sw) {
		case DatabaseStyle:
			s.UseDatabase(para)
		case SchemaStyle:
			s.UseSchema(para)
		default:
			util.PrintColorTips(util.LightRed, UseFailed)
		}
	} else {
		util.PrintColorTips(util.LightRed, UseFailed)
	}
}

// UseDatabase 选择数据库
func (s *Params) UseDatabase(dbName string) {
	//判断是否存在这个数据库
	info, err := P.GetDatabaseInfoByName(dbName)
	if err != nil {
		util.PrintColorTips(util.LightRed, UseDBFailed, err.Error())
		return
	}

	if len(info) == 0 {
		util.PrintColorTips(util.LightRed, UseDBNotExists)
		return
	}

	P.DataBase = dbName
	if err := P.Connect(); err == nil {
		*Database = dbName
		P.Schema = ""
		util.PrintColorTips(util.LightGreen, UseDBSucc)
	}
}

// UseSchema 选择模式
func (s *Params) UseSchema(schema string) {
	//判断是否存在这个模式
	info, err := P.GetSchemaFromNS(schema)
	if err != nil {
		util.PrintColorTips(util.LightRed, UseSchemaFailed, err.Error())
		return
	}

	if len(info) == 0 {
		util.PrintColorTips(util.LightRed, UseSchemaNotExists)
		return
	}
	P.Schema = schema
	if err := P.Connect(); err == nil {
		util.PrintColorTips(util.LightGreen, UseSchemaSucc)
	}
}
