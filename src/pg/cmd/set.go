package cmd

import (
	"pgii/src/pg/global"
	"pgii/src/util"
)

func (s *Params) Set() {
	//参数长度为2位
	if len(s.Param) != global.TwoCMDLength {
		util.PrintColorTips(util.LightRed, global.SetError)
		return
	}

	//查看DDL的类型
	sCmd := util.TrimLower(s.Param[0])
	if sCmd == "language" {
		s.SetLanguage()
	}
}

func (s *Params) SetLanguage() {
	lang := util.TrimLower(s.Param[1])
	if util.InArray(lang, global.ZhCN, global.ZhEN) {
		global.Language = lang
		util.PrintColorTips(util.LightGreen, global.SetLanguageSuccess)
	} else {
		util.PrintColorTips(util.LightRed, global.SetError)
	}
}
