package cmd

import (
	"fmt"
	"pgii/src/pg/db"
	"pgii/src/pg/global"
	"pgii/src/util"
)

func (s *Params) Size() {
	//参数长度为1位或2位
	if len(s.Param) > global.TwoCMDLength || len(s.Param) < global.OneCMDLength {
		util.PrintColorTips(util.LightRed, global.SizeFailed)
		return
	}

	//查看DDL的类型
	types := util.TrimLower(s.Param[0])
	switch CheckParamType(types) {
	case global.DatabaseStyle:
		s.SizeDatabase(s.Param[1:]...)
	case global.TableStyle:
		s.SizeTable(s.Param[1:])
	case global.IndexStyle:
		s.SizeIndex()
	case global.TableSpaceStyle:
		s.SizeTableSpace()
	default:
		util.PrintColorTips(util.LightRed, global.SizeFailed)
	}
}

// SizeTableSpace 取表空间大小
func (s *Params) SizeTableSpace() {
	//判断table信息
	tbParam, juErrMsg, err := s.judgeTableSpace()
	if err != nil {
		util.PrintColorTips(util.LightRed, juErrMsg)
		return
	}

	//获取数据
	if sizeInfo, err := db.P.GetSizeInfo(global.TableSpaceStyle, tbParam[0]); err == nil {
		//开始打印
		data := []interface{}{tbParam[0], sizeInfo["size"]}
		db.ShowTable(db.TableSpaceHeader, [][]interface{}{data})
	} else {
		fmt.Println(global.SizeFailedDataNull)
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
	if sizeInfo, err := db.P.GetSizeInfo(global.IndexStyle, tbParam[0]); err == nil {
		//开始打印
		data := []interface{}{tbParam[0], sizeInfo["size"]}
		db.ShowTable(db.IndexSizeHeader, [][]interface{}{data})
	} else {
		fmt.Println(global.SizeFailedDataNull)
	}
}

// SizeDatabase 取database的size
func (s *Params) SizeDatabase(dbName ...string) {
	//如果没有指定数据库，查询当前数据库
	useDb := ""
	if len(dbName) == global.ZeroCMDLength {
		//数据库是否为空
		if db.P.DataBase == "" {
			util.PrintColorTips(util.LightRed, global.SizeFailedNull)
			return
		}
		useDb = db.P.DataBase
	} else {
		//判断指定的db是否存在
		dbInfo, err := db.P.GetDatabaseInfoByName(dbName[0])
		if err != nil || len(dbInfo) == global.ZeroCMDLength {
			util.PrintColorTips(util.LightRed, global.SizeFailedNull)
			return
		}
		useDb = dbName[0]
	}

	//获取数据
	if sizeInfo, err := db.P.GetSizeInfo(global.DatabaseStyle, useDb); err == nil {
		//开始打印
		data := []interface{}{useDb, sizeInfo["size"]}
		db.ShowTable(db.DatabaseSizeHeader, [][]interface{}{data})
	} else {
		util.PrintColorTips(util.LightRed, global.SizeFailedDataNull)
	}
}

// SizeTable 取表的size
func (s *Params) SizeTable(param []string) {
	//判断table信息
	tbParam, juErrMsg, err := s.judgeTable()
	if err != nil {
		util.PrintColorTips(util.LightRed, juErrMsg)
		return
	}

	//获取数据
	if sizeInfo, err := db.P.GetSizeInfo(global.TableStyle, tbParam[0]); err == nil {
		//开始打印
		data := []interface{}{param[0], sizeInfo["size"]}
		db.ShowTable(db.TableSizeHeader, [][]interface{}{data})
	} else {
		fmt.Println(global.SizeFailedDataNull)
	}
}
