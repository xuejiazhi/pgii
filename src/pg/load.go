package pg

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"pgii/src/util"
)

func (s *Params) Load() {
	//参数长度至少为2位
	if len(s.Param) != TwoCMDLength {
		util.PrintColorTips(util.LightRed, LoadFailed)
		return
	}
	//get cmd 2
	sCmd := util.TrimLower(s.Param[0])
	switch CheckParamType(sCmd) {
	case TableStyle:
		s.LoadTable()
	case SchemaStyle:
		s.LoadSchema()
	case DatabaseStyle:
		s.LoadDataBase()
	default:
		util.PrintColorTips(util.LightRed, LoadFailed)
	}
}

// LoadTable 载入表
func (s *Params) LoadTable() {
	//必须选中模式
	if P.Schema == "" {
		util.PrintColorTips(util.LightRed, LoadFailedNoSelectSchema)
		return
	}
	//todo:
	pgFile := s.Param[1]
	//pgFile := "../dump_load/dump_table_user_1681300666.pgi"
	//judge
	if _, err := os.Stat(pgFile); err != nil {
		util.PrintColorTips(util.LightRed, LoadTableNOFile)
		return
	}

	//Open File
	f, err := os.Open(pgFile)
	if err != nil {
		util.PrintColorTips(util.LightRed, LoadNoFile, err.Error())
		return
	}

	defer f.Close() //

	content, err := ioutil.ReadAll(f)

	//解压
	unZipSQL, err := util.UnCompress(content)
	affect, err := P.ExecSQL(string(unZipSQL))
	if err != nil {
		util.PrintColorTips(util.LightRed, LoadTableExecSQLFailed, err.Error())
		return
	}
	//print Success
	util.PrintColorTips(util.LightGreen, LoadTableSQLSuccess, fmt.Sprintf(" Affect Nums:%d", affect))
}

// LoadSchema 载入模式
func (s *Params) LoadSchema() {
	//必须选中模式
	if P.DataBase == "" {
		util.PrintColorTips(util.LightRed, LoadFailedNoSelectDB)
		return
	}

	filePath := s.Param[1]
	//pgFile := "../dump_load/dump_table_user_1681300666.pgi"
	//judge
	if _, err := os.Stat(filePath); err != nil {
		util.PrintColorTips(util.LightRed, LoadSchemaNOPath)
		return
	}

	//取inifile
	iniFile := fmt.Sprintf("%s/_init_", filePath)
	if _, err := os.Stat(iniFile); err != nil {
		util.PrintColorTips(util.LightRed, LoadSchemaNOPath)
		return
	}

	f, err := os.Open(iniFile)
	defer f.Close()
	if err != nil {
		util.PrintColorTips(util.LightRed, LoadSchemaNOPath)
		return
	}

	reader := bufio.NewReader(f)
	//buffer := bytes.NewBuffer(make([]byte, 1024))

	for {
		//get file line
		part, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		fileLine := string(part)
		//Open File
		ft, err := os.Open(fileLine)
		if err != nil {
			util.PrintColorTips(util.LightRed, LoadNoFile, err.Error())
			continue
		}

		//read content
		content, err := ioutil.ReadAll(ft)

		//解压
		unZipSQL, err := util.UnCompress(content)
		affect, err := P.ExecSQL(string(unZipSQL))
		if err != nil {
			util.PrintColorTips(util.LightRed, LoadTableExecSQLFailed, err.Error())
			return
		}
		//print Success
		util.PrintColorTips(util.LightGreen, LoadTableSQLSuccess, fmt.Sprintf(" Affect Nums:%d", affect))
		ft.Close()
	}

	//todo:
}

// LoadDataBase 载入库
func (s *Params) LoadDataBase() {
	//todo:
}
