package pg

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

// QueryTableNums 获取表的行数
func (p *PgDsn) QueryTableNums(tableName string) (count int) {
	sqlStr := fmt.Sprintf("select count(1) as count from %s", tableName)
	p.PgConn.Raw(sqlStr).First(&count)
	return
}
