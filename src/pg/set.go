package pg

import "pgii/src/util"

func (s *Params) Set() {
	//参数长度为2位
	if len(s.Param) != TwoCMDLength {
		util.PrintColorTips(util.LightRed, SetError)
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
	if util.InArray(lang, []string{"en", "cn"}) {
		Language = lang
		util.PrintColorTips(util.LightGreen, SetLanguageSuccess)
	} else {
		util.PrintColorTips(util.LightRed, SetError)
	}
}
