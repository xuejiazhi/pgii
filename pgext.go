package main

import "fmt"

func (p *PgDsn) GetSizeInfo(style, Name string) (sizeInfo map[string]interface{}, err error) {
	sqlStr := ""
	if style == "db" {
		//T-SQL
		sqlStr = fmt.Sprintf("select pg_size_pretty( pg_database_size('%s') ) as size", Name)
	} else {
		//T-SQL
		sqlStr = fmt.Sprintf("select pg_size_pretty( pg_total_relation_size('%s') ) as size", Name)
	}

	//query
	err = p.PgConn.Raw(sqlStr).Scan(&sizeInfo).Error

	return
}
