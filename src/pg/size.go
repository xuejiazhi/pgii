package pg

import (
	"errors"
	"fmt"
	"pgii/src/util"
)

func (s *Params) Size() {
	//参数长度为1位或2位
	if len(s.Param) > TwoCMDLength || len(s.Param) < OneCMDLength {
		util.PrintColorTips(util.LightRed, SizeFailed)
		return
	}

	//查看DDL的类型
	types := util.TrimLower(s.Param[0])
	switch CheckParamType(types) {
	case DatabaseStyle:
		s.SizeDatabase(s.Param[1:]...)
	case TableStyle:
		s.SizeTable(s.Param[1:])
	case IndexStyle:
		s.SizeIndex()
	default:
		util.PrintColorTips(util.LightRed, SizeFailed)
	}
}

// SizeIndex 取index索引大小
func (s *Params) SizeIndex() {
	//判断table信息
	tbParam, juErrMsg, err := s.judgeTable()
	if err != nil {
		util.PrintColorTips(util.LightRed, juErrMsg)
		return
	}

	//获取数据
	if sizeInfo, err := P.GetSizeInfo(IndexStyle, tbParam[0]); err == nil {
		//开始打印
		data := []interface{}{tbParam[0], sizeInfo["size"]}
		ShowTable(TableSizeHeader, [][]interface{}{data})
	} else {
		fmt.Println(SizeFailedDataNull)
	}
}

// SizeDatabase 取database的size
func (s *Params) SizeDatabase(dbName ...string) {
	//如果没有指定数据库，查询当前数据库
	useDb := ""
	if len(dbName) == ZeroCMDLength {
		//数据库是否为空
		if P.DataBase == "" {
			util.PrintColorTips(util.LightRed, SizeFailedNull)
			return
		}
		useDb = P.DataBase
	} else {
		//判断指定的db是否存在
		dbInfo, err := P.GetDatabaseInfoByName(dbName[0])
		if err != nil || len(dbInfo) == ZeroCMDLength {
			util.PrintColorTips(util.LightRed, SizeFailedNull)
			return
		}
		useDb = dbName[0]
	}

	//获取数据
	if sizeInfo, err := P.GetSizeInfo(DatabaseStyle, useDb); err == nil {
		//开始打印
		data := []interface{}{useDb, sizeInfo["size"]}
		ShowTable(DatabaseSizeHeader, [][]interface{}{data})
	} else {
		util.PrintColorTips(util.LightRed, SizeFailedDataNull)
	}
}

// SizeTable 取表的size
func (s *Params) SizeTable(param []string) {
	//必须要指定表
	if len(param) == ZeroCMDLength {
		util.PrintColorTips(util.LightRed, SizeFailedPointTable)
		return
	}

	//判断指定的sc是否存在
	if scInfo, err := P.GetSchemaFromNS(P.Schema); err != nil || len(scInfo) == ZeroCMDLength {
		util.PrintColorTips(util.LightRed, SizeFailedNoSchema)
		return
	}

	//判断table是否存在
	if tbInfo, err := P.GetTableByName(param[0]); err != nil || len(tbInfo) == ZeroCMDLength {
		util.PrintColorTips(util.LightRed, SizeFailedNoTable)
		return
	}

	//获取数据
	if sizeInfo, err := P.GetSizeInfo(TableStyle, param[0]); err == nil {
		//开始打印
		data := []interface{}{param[0], sizeInfo["size"]}
		ShowTable(TableSizeHeader, [][]interface{}{data})
	} else {
		fmt.Println(SizeFailedDataNull)
	}
}

func (s *Params) judgeTable() (tbParam []string, errorMsg string, err error) {
	//赋值
	tbParam = s.Param[1:]
	//必须要指定表
	if len(tbParam) == ZeroCMDLength {
		errorMsg = SizeFailedPointTable
	}

	//判断指定的sc是否存在
	scInfo, err := P.GetSchemaFromNS(P.Schema)
	if err != nil {
		errorMsg = fmt.Sprintf("%s %s", SizeFailedNoSchema, err.Error())
	}
	if len(scInfo) == ZeroCMDLength {
		errorMsg = SizeFailedNoSchema
		return
	}

	//判断table是否存在
	tbInfo, err := P.GetTableByName(tbParam[0])
	if err != nil {
		errorMsg = fmt.Sprintf("%s %s", SizeFailedNoTable, err.Error())
	}
	if len(tbInfo) == ZeroCMDLength {
		errorMsg = SizeFailedNoTable
	}

	err = errors.New(errorMsg)
	//
	return
}
