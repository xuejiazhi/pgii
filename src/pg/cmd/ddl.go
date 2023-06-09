package cmd

import (
	"fmt"
	"pgii/src/pg/db"
	"pgii/src/pg/global"
	"pgii/src/util"
)

func (s *Params) DDL() {
	if len(s.Param) != global.TwoCMDLength {
		fmt.Println("Failed:DDL Cmd fail")
		return
	}

	//查看DDL的类型
	sCmd := util.TrimLower(s.Param[0])
	name := util.TrimLower(s.Param[1])
	switch CheckParamType(sCmd) {
	case global.SchemaStyle: //查看schema的DDL
		s.DDLSchema(name)
	case global.TableStyle: //查看table的DDL
		s.DDLTable(name)
	case global.ViewStyle: //查看view的DDL
		s.DDLView(name)
	default:
		fmt.Println("Failed:DDL Cmd fail")
		return
	}
}

// DDLSchema 生成SCHEMA的DDL
func (s *Params) DDLSchema(name string) {
	//校验schema 是否存在
	info, err := db.P.GetSchemaFromNS(name)
	if err != nil {
		util.PrintColorTips(util.LightRed, global.DDLSchemaError, err.Error())
		return
	}

	if len(info) == 0 {
		util.PrintColorTips(util.LightRed, global.DDLSchemaNotExists)
		return

	}

	//print schema ddl
	fmt.Println(db.GenerateSchema(name))
}

// DDLTable 生成Table的DDL
func (s *Params) DDLTable(name string) {
	//get table info
	info, err := db.P.GetTableByName(name)
	if err != nil {
		util.PrintColorTips(util.LightRed, global.DDLTableError, err.Error())
		return
	}

	if len(info) == 0 {
		util.PrintColorTips(util.LightRed, global.DDLTableNoExists)
		return
	}

	//print
	util.PrintColorTips(util.LightSeaBlue, getTableDdlSql(db.P.Schema, name))
}

// DDLView 生成view视图的DDL
// viewName 视图名称
func (s *Params) DDLView(viewName string) {
	//获取view信息
	viewInfo, err := db.P.Views("filter", viewName)
	if err != nil {
		util.PrintColorTips(util.LightRed, global.DDLViewError, err.Error())
		return
	}

	//判断信息不为空
	if len(viewInfo) == 0 {
		util.PrintColorTips(util.LightRed, global.DDLViewNoExists)
		return
	}

	//print Create View SQL
	fmt.Println("========= Create View Success ============")
	if def, ok := viewInfo[0]["definition"]; ok {
		fmt.Println(fmt.Sprintf(" CREATE OR REPLACE VIEW \"%s\".%s\n AS%s", db.P.Schema, viewName, def))
	}
}
