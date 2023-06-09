package db

import (
	"fmt"
	"github.com/spf13/cast"
	"pgii/src/pg/global"
	"strings"
)

// GenerateSchema  ================DDL DUMP 使用===============================//
func GenerateSchema(scName string, style ...int) (scStr string) {
	//print Create schema SQL
	scStr = "-- Create Schema Success \n"
	//Drop Schema
	scStr += fmt.Sprintf("-- DROP SCHEMA %s;\n", scName)

	//get create schema
	scStr += fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS \"%s\" AUTHORIZATION %s;", scName, *global.UserName)

	//search path
	if len(style) > 0 {
		if style[0] == global.DUMP {
			scStr += fmt.Sprintf("SET search_path TO \"%s\";\n", scName)
		}
	}

	//return
	return
}

func GenerateBatchValue(idx int, tbName string, columnList []string, columnType map[string]string) (batchValue []string) {
	//查询的SQL
	querySQL := P.GetQuerySql(tbName, columnList, columnType, idx)
	//run sql
	data, err := P.RunSQL(querySQL)
	if err != nil || len(data) == 0 {
		return
	}

	//循环
	for _, v := range data {
		var valSon []string
		l := 0
		for _, sv := range columnList {
			//judge
			if _, ok := v[sv]; !ok {
				break
			}
			//处理当数据中存在' 符号的存在
			//Handles errors raised when the ' symbol is present in the data
			valStr := ""
			if v[sv] == nil {
				valSon = append(valSon, "NULL")
			} else {
				valStr = strings.Replace(cast.ToString(v[sv]), "'", "''", -1)
				valSon = append(valSon, fmt.Sprintf("'%s'", valStr))
			}
			l++
		}

		//加入数组
		batchValue = append(batchValue, fmt.Sprintf("(%s)", strings.Join(valSon, ",")))
	}
	return
}

// GenerateBatchSql
func GenerateBatchSql(i, style int, tbName string, columnList []string, columnType map[string]string) (batchSql string) {
	//get batchSQL
	//定义定入的SQL
	fullTbName := fmt.Sprintf(`"%s"`, tbName)
	if style == global.SchemaStyle {
		fullTbName = fmt.Sprintf(`"%s"."%s"`, P.Schema, tbName)
	}
	//get batch value
	batchValue := GenerateBatchValue(i, fullTbName, columnList, columnType)
	if len(batchValue) > 0 {
		batchSql = fmt.Sprintf(`Insert into %s(%s) values %s;`,
			fullTbName,
			strings.Join(columnList, ","),
			strings.Join(batchValue, ","))
	}
	//返回
	return
}
