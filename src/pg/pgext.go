package pg

import (
	"fmt"
	"github.com/spf13/cast"
	"pgii/src/util"
	"strings"
)

func (p *PgDsn) GetSizeInfo(style, Name string) (sizeInfo map[string]interface{}, err error) {
	sqlStr := ""

	// 获取哪种类型的size
	sizeType := util.If(style == "db", "pg_database_size", "pg_total_relation_size")

	//T-SQL
	sqlStr = fmt.Sprintf("select pg_size_pretty( %s('%s') ) as size", sizeType, Name)

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

// GetPgClassForTbName 根据tablename 获取pg_class 内容
// tableName 表名
func (p *PgDsn) GetPgClassForTbName(tableName string) (classInfo map[string]interface{}, err error) {
	//T-SQL
	sqlStr := fmt.Sprintf("select oid,* from  pg_class where relname='%s'", tableName)
	//query
	err = p.PgConn.Raw(sqlStr).First(&classInfo).Error
	//return
	return
}

// GetPgClassValueForTbName 获取pg_class中字段的值
// tableName 表名 fieldName 字段名
func (p *PgDsn) GetPgClassValueForTbName(tableName string, fieldName ...string) (value map[string]interface{}, err error) {
	//获取pg_class表内容
	classInfo, err := p.GetPgClassForTbName(tableName)
	if len(fieldName) == 0 {
		return
	}

	fv := map[string]interface{}{}
	//range filed
	for _, v := range fieldName {
		if _, ok := classInfo[v]; ok {
			fv[v] = classInfo[v]
		}
	}
	value = fv

	return
}

// GetPgTriggerDef 获取触发器
func (p *PgDsn) GetPgTriggerDef(oid int) (triggerDef map[string]interface{}, err error) {
	//T-SQL
	sqlStr := fmt.Sprintf("select pg_get_triggerdef(%d) as def", oid)
	//query
	err = p.PgConn.Raw(sqlStr).First(&triggerDef).Error
	//return
	return
}

// GetColumnList 获取字段列表
func (p *PgDsn) GetColumnList(tbName string) (cols []string) {
	//取co
	if columnList, err := p.Column(tbName); err == nil {
		if len(columnList) == 0 {
			return
		}
		//拼接column
		for _, c := range columnList {
			if _, ok := c["column_name"]; ok {
				cols = append(cols, cast.ToString(c["column_name"]))
			}
		}

	}

	return
}

// GetQuerySql 获取查询SQL
func (p *PgDsn) GetQuerySql(tbName string, fieldList []string, pageSize int) (sqlStr string) {
	//生成SQL
	sqlStr = fmt.Sprintf("SELECT %s FROM %s  OFFSET %d LIMIT %d",
		strings.Join(fieldList, ","),
		tbName,
		pageSize*PgLimit,
		PgLimit)
	return
}
