package pg

import (
	"fmt"
	"github.com/spf13/cast"
	"pgii/src/util"
	"strings"
)

func (p *PgDsn) GetSizeInfo(style int, Name string) (sizeInfo map[string]interface{}, err error) {
	sqlStr := ""

	// 获取哪种类型的size
	sizeType := func() string {
		switch style {
		case DatabaseStyle:
			return "pg_database_size"
		case TableStyle:
			return "pg_total_relation_size"
		case IndexStyle:
			return "pg_indexes_size"
		default:
			return "pg_database_size"
		}
	}()

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

// GetColumnsType 获取表字段的类型
func (p *PgDsn) GetColumnsType(tbName string, columns ...string) (types map[string]string) {
	//取co
	newType := map[string]string{}
	if columnList, err := p.Column(tbName); err == nil {
		if len(columnList) == 0 {
			return
		}
		//拼接column
		for _, c := range columnList {
			if _, ok := c["column_name"]; !ok {
				continue
			}

			columnName := cast.ToString(c["column_name"])
			if util.InArray(columnName, columns) {
				newType[columnName] = cast.ToString(c["udt_name"])
			}
		}
	}
	types = newType
	return
}

// GetQuerySql 获取查询SQL
func (p *PgDsn) GetQuerySql(tbName string, fieldList []string, columnTypes map[string]string, pageSize int) (sqlStr string) {
	//range filed
	newFiledList := make([]string, len(fieldList))
	copy(newFiledList, fieldList)
	for k, v := range newFiledList {
		if ct, ok := columnTypes[v]; ok {
			newFiledList[k] = util.TypeTransForm(ct, v)
		}
	}
	//生成SQL
	sqlStr = fmt.Sprintf("SELECT %s FROM %s  OFFSET %d LIMIT %d",
		strings.Join(newFiledList, ","),
		tbName,
		pageSize*PgLimit,
		PgLimit)
	return
}

// 获取最链接数
func (p *PgDsn) GetConnectionNums(types int) (connection map[string]interface{}, err error) {
	//T-SQL
	prefixStr := func() string {
		switch types {
		case MaxConnections:
			return "max_connections"
		case SuperuserReservedConnections:
			return "superuser_reserved_connections"
		default:
			return "max_connections"
		}
	}()

	sqlStr := fmt.Sprintf("show %s", prefixStr)
	//query
	err = p.PgConn.Raw(sqlStr).Scan(&connection).Error
	//return
	return
}

// GetUseConnection 获取剩余连接数
func (p *PgDsn) GetUseConnection(types int) (resiConn map[string]interface{}, err error) {
	//
	sqlStr := ""
	switch types {
	case RemainingConnections:
		sqlStr = `
select
	max_conn-now_conn as conn_nums
from
	(
	select
		setting::int8 as max_conn,
		(
		select
			count(*)
		from
			pg_stat_activity) as now_conn
	from
		pg_settings
	where
		name = 'max_connections') t;
`
	case InUseConnections:
		sqlStr = `select
			count(*) as conn_nums
		from
			pg_stat_activity
`
	}

	//query
	err = p.PgConn.Raw(sqlStr).Scan(&resiConn).Error
	//return
	return
}
