package cmd

import (
	"pgii/src/pg/db"
	"pgii/src/pg/global"
	"pgii/src/util"
)

// Desc 查看表结构
func (s *Params) Desc() {
	if len(s.Param) != global.OneCMDLength {
		util.PrintColorTips(util.LightRed, global.DescTableFailed)
		return
	}
	//获取表名
	tableName := util.TrimLower(s.Param[0])

	//校验是否存在表
	tbInfo, err := db.P.GetTableByName(tableName)
	if err != nil {
		util.PrintColorTips(util.LightRed, global.DescTableError, err.Error())
		return
	}

	if len(tbInfo) == 0 {
		util.PrintColorTips(util.LightRed, global.DescTableNoExists)
		return
	}

	if columnInfo, err := db.P.Column(db.P.Schema, tableName); err == nil {
		//序列化输出
		var desc [][]interface{}
		for _, v := range columnInfo {
			var sbs []interface{}
			//oid
			sbs = append(sbs,
				v["dtd_identifier"],
				v["column_name"],
				v["udt_name"],
				v["character_maximum_length"],
				v["is_nullable"],
				v["column_default"])
			desc = append(desc, sbs)
		}
		db.ShowTable(db.DescTableHeader, desc)
	}
}
