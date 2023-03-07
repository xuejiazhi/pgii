package main

import (
	"fmt"
	"strings"
)

// Use 用做切换数据库
func Use(cmdList []string) {
	//判断
	if len(cmdList) == 2 {
		sw := strings.Trim(cmdList[0], "")
		para := strings.Trim(cmdList[1], "")
		switch sw {
		case "db", "database":
			UseDatabase(para)
		case "sc", "schema":
			UseSchema(para)
		default:
			fmt.Println("Failed:Use cmd fail!")
		}
	} else {
		fmt.Println("Failed:Use Database fail!")
	}
}

// UseDatabase 选择数据库
func UseDatabase(dbName string) {
	//判断是否存在这个数据库
	info, err := P.GetDatabaseInfoByName(dbName)
	if err != nil {
		fmt.Println("Failed:Use Database fail!", err.Error())
		return
	}

	if len(info) == 0 {
		fmt.Println("Failed:Use Database fail,DataBase Not Exists!")
		return
	}

	P.DataBase = dbName
	if err := P.Connect(); err == nil {
		*Database = dbName
		P.Schema = ""
		fmt.Println("Use Database Success!")
	}
}

// UseSchema 选择模式
func UseSchema(schema string) {
	//判断是否存在这个模式
	info, err := P.GetSchemaFromNS(schema)
	if err != nil {
		fmt.Println("Failed:Use Schema fail!", err.Error())
		return
	}

	if len(info) == 0 {
		fmt.Println("Failed:Use Schema fail,Schema Not Exists!")
		return
	}
	P.Schema = schema
	if err := P.Connect(); err == nil {
		fmt.Println("Use Schema Success!")
	}
}
