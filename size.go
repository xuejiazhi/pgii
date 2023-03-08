package main

import "fmt"

func Size(params ...string) {
	//参数长度为1位或2位
	if len(params) > 2 || len(params) < 1 {
		fmt.Println("Failed:DDL Cmd fail")
		return
	}

	//查看DDL的类型
	types := TrimLower(params[0])
	switch types {
	case "database", "db": //查看数据库大小
		SizeDatabase(params[1:]...)
	case "table", "tb": //查看表大小
		SizeTable(params[1:])
	default:
		fmt.Println("Failed:CmdLine is Wrong!")
	}

}

func SizeDatabase(dbName ...string) {
	//如果没有指定数据库，查询当前数据库
	useDb := ""
	if len(dbName) == 0 {
		//数据库是否为空
		if P.DataBase == "" {
			fmt.Println("Failed:Size Database is Null!")
			return
		}
		useDb = P.DataBase
	} else {
		//判断指定的db是否存在
		dbInfo, err := P.GetDatabaseInfoByName(dbName[0])
		if err != nil || len(dbInfo) == 0 {
			fmt.Println("Failed:Size Database is Not Exists!")
			return
		}
		useDb = dbName[0]
	}

	//获取数据
	if sizeInfo, err := P.GetSizeInfo("db", useDb); err == nil {
		//开始打印
		header := []interface{}{"database", "size"}
		data := []interface{}{useDb, sizeInfo["size"]}
		ShowTable(header, [][]interface{}{data})
	} else {
		fmt.Println("Failed:Size Database Get Data Nil!")
	}
}

func SizeTable(param []string) {
	//必须要指定表
	if len(param) == 0 {
		fmt.Println("Failed:Size Table Must Point Table Name!")
		return
	}

	//判断指定的sc是否存在
	if scInfo, err := P.GetSchemaFromNS(P.Schema); err != nil || len(scInfo) == 0 {
		fmt.Println("Failed:Size Cmd Schema fail,Schema not exists!")
		return
	}

	//判断table是否存在
	if tbInfo, err := P.GetTableByName(param[0]); err != nil || len(tbInfo) == 0 {
		fmt.Println("Failed:Size Cmd Table fail,Table not exists!")
		return
	}

	//获取数据
	if sizeInfo, err := P.GetSizeInfo("tb", param[0]); err == nil {
		//开始打印
		header := []interface{}{"tablename", "size"}
		data := []interface{}{param[0], sizeInfo["size"]}
		ShowTable(header, [][]interface{}{data})
	} else {
		fmt.Println("Failed:Size TableName Get Data Nil!")
	}

}
