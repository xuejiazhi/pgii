package cmd

import (
	"pgii/src/pg/db"
	"pgii/src/pg/global"
	"pgii/src/util"
)

func (s *Params) Clear() {
	//参数长度至少为2位
	if len(s.Param) != global.TwoCMDLength {
		util.PrintColorTips(util.LightRed, global.ClearParamLengthErrror)
		return
	}

	//get cmd 2
	sCmd := util.TrimLower(s.Param[0])
	switch CheckParamType(sCmd) {
	case global.TableStyle:
		s.ClearTable(s.Param[1:])
	default:
		util.PrintColorTips(util.LightRed, global.ClearErr)
	}
}

// ClearTable 清空表数据
func (s *Params) ClearTable(param []string) {
	//判断table信息
	tbParam, juErrMsg, err := s.judgeTable()
	if err != nil {
		util.PrintColorTips(util.LightRed, juErrMsg)
		return
	}

	//must select schema
	if db.P.Schema == "" {
		util.PrintColorTips(util.LightRed, global.ClearFailedNoSelectSchema)
		return
	}

	//取表名
	if "" == s.Param[1] {
		util.PrintColorTips(util.LightRed, global.ClearFailedNoTable)
		return
	}

	//获取数据
	if _, err := db.P.ClearTable(tbParam[0]); err == nil {
		//开始打印
		util.PrintColorTips(util.LightGreen, global.ClearSuccess)
	} else {
		util.PrintColorTips(util.LightRed, global.ClearErr)
	}
}
