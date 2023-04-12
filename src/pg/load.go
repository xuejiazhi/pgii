package pg

import (
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
	//todo:
}

// LoadDataBase 载入库
func (s *Params) LoadDataBase() {
	//todo:
}
