package cmd

import (
	"bufio"
	"fmt"
	"os"
	"pgii/src/pg/db"
	"pgii/src/pg/global"
	"pgii/src/util"
)

func (s *Params) Load() {
	//参数长度至少为2位
	if len(s.Param) != global.TwoCMDLength {
		util.PrintColorTips(util.LightRed, global.LoadFailed)
		return
	}
	//get cmd 2
	sCmd := util.TrimLower(s.Param[0])
	switch CheckParamType(sCmd) {
	case global.TableStyle:
		s.LoadTable()
	case global.SchemaStyle:
		s.LoadSchema()
	case global.DatabaseStyle:
		s.LoadDataBase()
	default:
		util.PrintColorTips(util.LightRed, global.LoadFailed)
	}
}

// LoadTable 载入表
func (s *Params) LoadTable() {
	//必须选中模式
	if db.P.Schema == "" {
		util.PrintColorTips(util.LightRed, global.LoadFailedNoSelectSchema)
		return
	}

	//get ini file
	iniFile, err := s.getIniFile()
	if err != nil {
		util.PrintColorTips(util.LightRed, global.LoadTableNOFile, err.Error())
		return
	}

	//open file
	f, err := os.Open(iniFile)
	defer f.Close()
	if err != nil {
		util.PrintColorTips(util.LightRed, global.LoadTableNOFile)
		return
	}

	//define reader
	reader := bufio.NewReader(f)
	for {
		//get file line
		part, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		//run sql
		_ = execUnzipSQL(string(part))
	}
}

// LoadSchema 载入模式
func (s *Params) LoadSchema() {
	//必须选中模式
	if db.P.DataBase == "" {
		util.PrintColorTips(util.LightRed, global.LoadFailedNoSelectDB)
		return
	}

	//get ini file
	iniFile, err := s.getIniFile()
	if err != nil {
		util.PrintColorTips(util.LightRed, global.LoadSchemaNOPath, err.Error())
		return
	}

	//open file
	f, err := os.Open(iniFile)
	defer f.Close()
	if err != nil {
		util.PrintColorTips(util.LightRed, global.LoadSchemaNOPath)
		return
	}

	//define reader
	reader := bufio.NewReader(f)
	for {
		//get file line
		part, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		//run sql
		_ = execUnzipSQL(string(part))
	}
}

// LoadDataBase 载入库
func (s *Params) LoadDataBase() {
	//todo:
}

func (s *Params) getIniFile() (iniFile string, err error) {
	filePath := s.Param[1]
	//pgFile := "../dump_load/dump_table_user_1681300666.pgi"
	//judge
	if _, err = os.Stat(filePath); err != nil {
		util.PrintColorTips(util.LightRed, global.LoadSchemaNOPath)
		return
	}

	//取inifile
	iniFile = fmt.Sprintf("%s/%s", filePath, global.INIFile)
	if _, err = os.Stat(iniFile); err != nil {
		util.PrintColorTips(util.LightRed, global.LoadSchemaNOPath)
		return
	}

	//
	return
}
