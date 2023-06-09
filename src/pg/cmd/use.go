package cmd

import (
	"pgii/src/pg/db"
	"pgii/src/pg/global"
	"pgii/src/util"
	"strings"
)

// Use 用做切换数据库
func (s *Params) Use() {
	//判断
	if len(s.Param) == global.TwoCMDLength {
		sw := strings.Trim(s.Param[0], "")
		para := strings.Trim(s.Param[1], "")
		switch CheckParamType(sw) {
		case global.DatabaseStyle:
			s.UseDatabase(para)
		case global.SchemaStyle:
			s.UseSchema(para)
		default:
			util.PrintColorTips(util.LightRed, global.UseFailed)
		}
	} else {
		util.PrintColorTips(util.LightRed, global.UseFailed)
	}
}

// UseDatabase 选择数据库
func (s *Params) UseDatabase(dbName string) {
	//判断是否存在这个数据库
	info, err := db.P.GetDatabaseInfoByName(dbName)
	if err != nil {
		util.PrintColorTips(util.LightRed, global.UseDBFailed, err.Error())
		return
	}

	if len(info) == 0 {
		util.PrintColorTips(util.LightRed, global.UseDBNotExists)
		return
	}

	//set database
	db.P.DataBase = dbName
	if err := db.P.Connect(); err == nil {
		*global.Database = dbName
		db.P.Schema = ""
		util.PrintColorTips(util.LightGreen, global.UseDBSuch)
	}
}

// UseSchema 选择模式
func (s *Params) UseSchema(schema string) {
	//判断是否存在这个模式
	info, err := db.P.GetSchemaFromNS(schema)
	if err != nil {
		util.PrintColorTips(util.LightRed, global.UseSchemaFailed, err.Error())
		return
	}

	if len(info) == 0 {
		util.PrintColorTips(util.LightRed, global.UseSchemaNotExists)
		return
	}

	//set schema
	db.P.Schema = schema
	if err := db.P.Connect(); err == nil {
		util.PrintColorTips(util.LightGreen, global.UseSchemaSuch)
	}
}
