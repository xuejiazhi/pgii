package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cast"
	"os"
	"pgii/src/pg/db"
	"time"
)

func Psql(cmdRun, cmdStr string) {
	//执行SQL语句
	bTime := time.Now().UnixMilli()
	//执行指令
	switch cmdRun {
	case
		"select",
		"explain":
		if val, err := db.P.RunSQL(cmdStr); err != nil {
			fmt.Println("Run T-SQL Error,error ", err.Error())
		} else {
			eTime := time.Now().UnixMilli()
			ShowQuery(eTime-bTime, val)
		}
	case
		"update",
		"insert",
		"delete",
		"alter",
		"create",
		"drop":
		if affectRows, err := db.P.ExecSQL(cmdStr); err != nil {
			fmt.Println(fmt.Sprintf("Run %s TSQL Error,error %s", cmdRun, err.Error()))
		} else {
			fmt.Println(fmt.Sprintf("Run %s TSQL Success,Affect Rows %d Line", cmdRun, affectRows))
		}
	default:
		fmt.Println("Run T-SQL Failed ")
	}
}

// ShowQuery 显示
func ShowQuery(timed int64, val []map[string]interface{}) {
	/// Set Header
	var Header table.Row
	if len(val) > 0 {
		for k, _ := range val[0] {
			Header = append(Header, k)
		}
	}

	timeStr := "RunTimes "
	if timed >= 100 {
		timeStr += fmt.Sprintf("%0.2fs", cast.ToFloat64(timed)/1000)
	} else {
		timeStr += fmt.Sprintf("%dns", timed)
	}
	//序列化输出
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(Header)
	//数据记录
	var tbs []table.Row
	for _, v := range val {
		var sbs []interface{}
		//oid
		for _, sv := range Header {
			sbs = append(sbs, v[cast.ToString(sv)])
		}
		tbs = append(tbs, sbs)
	}
	t.AppendRows(tbs)
	t.SetCaption(fmt.Sprintf("[Total: %d Rows]  [%s]\n", len(val), timeStr))
	t.Render()
}
