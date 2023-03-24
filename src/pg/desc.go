package pg

import (
	"pgii/src/util"
)

// Desc 查看表结构
func (s *Params) Desc() {
	if len(s.Param) != OneCMDLength {
		util.PrintColorTips(util.LightRed, DescTableFailed)
		return
	}
	//获取表名
	tableName := util.TrimLower(s.Param[0])

	//校验是否存在表
	tbInfo, err := P.GetTableByName(tableName)
	if err != nil {
		util.PrintColorTips(util.LightRed, DescTableError, err.Error())
		return
	}

	if len(tbInfo) == 0 {
		util.PrintColorTips(util.LightRed, DescTableNoExists)
		return
	}

	if columnInfo, err := P.Column(tableName); err == nil {
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
		ShowTable(DescTableHeader, desc)
	}
}
