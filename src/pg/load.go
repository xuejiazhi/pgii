package pg

import "pgii/src/util"

func (s *Params) Load() {
	//参数长度至少为1位
	if len(s.Param) < OneCMDLength {
		util.PrintColorTips(util.LightRed, DumpFailed)
		return
	}
}

// LoadTable 载入表
func (s *Params) LoadTable(dumpName string) {
	//todo:
}

// LoadSchema 载入模式
func (s *Params) LoadSchema(dumpName string) {
	//todo:
}

// LoadDataBase 载入库
func (s *Params) LoadDataBase(dumpName string) {
	//todo:
}
