package pg

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cast"
	"os"
	"pgii/src/util"
	"strings"
)

// CheckParamType 检查传过来的参数
func CheckParamType(types string) int {
	switch types {
	case "database", "db": //数据库
		return DatabaseStyle
	case "table", "tb": //表
		return TableStyle
	case "index", "idx": //索引
		return IndexStyle
	case "view", "vw": //视图
		return ViewStyle
	case "sd", "selectdb":
		return SelectStyle
	case "sc", "schema":
		return SchemaStyle
	case "tg", "trigger":
		return TriggerStyle
	case "ver", "version":
		return VersionStyle
	case "conn", "connection":
		return ConnectionStyle
	case "proc", "process": //当前进程
		return ProcessStyle
	case "tablespace", "tbsp": //表空间
		return TableSpaceStyle
	default:
		return NoneStyle
	}
}

func ShowTable(header string, data [][]interface{}) {
	//默认为英文
	if !util.InArray(Language, ZhCN, ZhEN) {
		Language = ZhEN
	}
	showHeader := ShowPrettyHeader[Language][header]
	prettyTable := table.NewWriter()
	//prettyTable.SetStyle(table.StyleLight)
	prettyTable.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:         "total",
			Colors:       text.Colors{text.BgHiGreen, text.Bold},
			ColorsHeader: text.Colors{text.BgHiGreen, text.FgHiYellow, text.Bold},
			ColorsFooter: text.Colors{text.BgHiGreen, text.FgHiYellow},
		},
		{
			Name:         "used%",
			Colors:       text.Colors{text.BgHiBlack, text.FgHiGreen, text.Bold},
			ColorsHeader: text.Colors{text.BgHiRed, text.FgGreen, text.Bold},
			ColorsFooter: text.Colors{text.BgHiRed, text.FgGreen},
		},
	})
	prettyTable.SetOutputMirror(os.Stdout)
	prettyTable.AppendHeader(showHeader)
	if len(data) > 0 {
		for _, v := range data {
			prettyTable.AppendRow(v)
			prettyTable.AppendSeparator()
		}
	}
	prettyTable.Render()
}

// ================DDL DUMP 使用===============================//
func generateSchema(scName string, style ...int) (scStr string) {
	//print Create schema SQL
	scStr = "-- Create Schema Success \n"
	//Drop Schema
	scStr += fmt.Sprintf("-- DROP SCHEMA %s;\n", scName)

	//get create schema
	scStr += fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS \"%s\" AUTHORIZATION %s;", scName, *UserName)

	//search path
	if len(style) > 0 {
		if style[0] == DUMP {
			scStr += fmt.Sprintf("SET search_path TO \"%s\";\n", scName)
		}
	}

	//return
	return
}

// =======================DUMP使用==========================//
func generateBatchValue(idx int, tbName string, columnList []string, columnType map[string]string) (batchValue []string) {
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

func fileClose(f *os.File) {
	err := f.Close()
	if err != nil {
		util.PrintColorTips(util.LightRed, CloseFileFailed)
	}
}
