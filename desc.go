package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

// Desc 查看表结构
func Desc(cms ...string) {
	if len(cms) != 1 {
		fmt.Println("Failed:Describe Table fail")
		return
	}
	//获取表名
	tableName := TrimLower(cms[0])
	//校验是否存在表
	if tbInfo, err := P.GetTableByName(tableName); err != nil {
		fmt.Println("Failed:Describe Table fail!Error ", err.Error())
		return
	} else {
		if len(tbInfo) == 0 {
			fmt.Println("Failed:Describe Table fail!Table not exists")
			return
		}
	}

	if columnInfo, err := P.Column(tableName); err == nil {
		//序列化输出
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "column", "dataType", "length", "isnull", "defaultValue"})
		var tbs []table.Row
		for _, v := range columnInfo {
			var sbs []interface{}
			//oid
			sbs = append(sbs,
				v["dtd_identifier"],
				v["column_name"],
				v["udt_name"],
				v["character_maximum_length"],
				v["is_nullable"],
				v["column_default"])
			tbs = append(tbs, sbs)
		}
		t.AppendRows(tbs)
		t.Render()
	}
}
