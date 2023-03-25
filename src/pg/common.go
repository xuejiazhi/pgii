package pg

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cast"
	"os"
	"strings"
)

const (
	ZeroCMDLength = iota
	OneCMDLength
	TwoCMDLength
	ThreeCMDLength
)

const (
	NoneStyle = iota
	DatabaseStyle
	TableStyle
	IndexStyle
	ViewStyle
	SelectStyle
	SchemaStyle
	TriggerStyle
	VersionStyle
	ConnectionStyle
	ProcessStyle
	TableSpaceStyle //表空间
)

const (
	MaxConnections               = iota //最大连接数
	SuperuserReservedConnections        //超级用户保留的连接数
	RemainingConnections                //剩余连接数
	InUseConnections                    //正在使用的链接数
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

func ShowTable(header []interface{}, data [][]interface{}) {
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
	prettyTable.AppendHeader(header)
	if len(data) > 0 {
		for _, v := range data {
			prettyTable.AppendRow(v)
			prettyTable.AppendSeparator()
		}
	}
	prettyTable.Render()
}

// ================DDL DUMP 使用===============================//
func generateSchema(scName string) (scStr string) {
	//print Create schema SQL
	scStr = "========= Create Schema Success ============\n"
	scStr += fmt.Sprintf("-- DROP SCHEMA %s;\n", scName)
	scStr += fmt.Sprintf("CREATE SCHEMA \"%s\" AUTHORIZATION %s;", scName, *UserName)
	return
}

// =======================DUMP使用==========================//
func generateBatchValue(idx int, tbName string, columnList []string, columnType map[string]string) (batchValue []string) {
	//查询的SQL
	querySQL := P.GetQuerySql(tbName, columnList, columnType, idx)
	//run sql
	data, err := P.RunSQL(querySQL)
	if err != nil {
		fmt.Println("value is null")
		return
	}

	if len(data) == 0 {
		fmt.Println("value is null")
		return
	}
	//循环
	for _, v := range data {
		valSon := []string{}
		l := 0
		for _, sv := range columnList {
			if _, ok := v[sv]; ok {
				valSon = append(valSon, fmt.Sprintf("'%s'", cast.ToString(v[sv])))
				l++
			} else {
				break
			}
		}
		//加入数组
		batchValue = append(batchValue, "("+strings.Join(valSon, ",")+")")
	}
	return
}