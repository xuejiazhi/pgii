package pg

import (
	"errors"
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
		case TableSpaceStyle: //表空间
			return "pg_tablespace_size"
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

func (p *PgDsn) GetProcessByPid(pid int) (process map[string]interface{}, err error) {
	///Get Column TSQL
	sqlStr := fmt.Sprintf("select pid,datname,application_name,state from pg_stat_activity where pid=%d", pid)
	//query
	err = p.PgConn.Raw(sqlStr).Scan(&process).Error
	//return
	return
}

// CancelProcessByPid 终止一个后台服务进程，同时释放此后台服务进程的资源
func (p *PgDsn) CancelProcessByPid(pid int) (err error) {
	//pg_terminate_backend
	if pInfo, e := p.GetProcessByPid(pid); e == nil {
		if len(pInfo) == 0 {
			err = errors.New(fmt.Sprintf("process %d is not exists", pid))
		} else {
			sqlStr := fmt.Sprintf("select pg_terminate_backend(%d)", pid)
			err = p.PgConn.Exec(sqlStr).Error
		}
	} else {
		err = e
	}
	return
}

func (p *PgDsn) getTableViewCondition(style, cmd string, param ...string) (condition string) {
	if p.Schema != "" {
		condition += fmt.Sprintf(" schemaname='%s'", p.Schema)
	}

	//查表还是查视图
	useName := util.If(style == "view", "viewname", "tablename")

	//加上过滤条件
	if util.InArray(cmd, EqualAndFilter) {
		if len(param) == 0 {
			return
		}

		inParam := ""
		for k, v := range param {
			//eq
			if util.InArray(cmd, EqualVar) {
				inParam += cast.ToString(util.If(len(param)-1 == k,
					fmt.Sprintf("'%s'", v),
					fmt.Sprintf("'%s',", v)))
			}

			//filter
			if util.InArray(cmd, FilterVar) {
				inParam += cast.ToString(util.If(len(param)-1 == k,
					fmt.Sprintf("%s like '%%%s%%'", useName, v),
					fmt.Sprintf("%s like '%%%s%%' or ", useName, v)))
			}
		}
		inParam = fmt.Sprintf("(%s)", inParam)
		//eq的处理
		if util.InArray(cmd, EqualVar) {
			condition += cast.ToString(util.If(condition == "",
				fmt.Sprintf(" %s in %s", useName, inParam),
				fmt.Sprintf(" and %s in %s", useName, inParam)))
		}

		//filter的处理
		if util.InArray(cmd, FilterVar) {
			condition += cast.ToString(
				util.If(condition == "",
					fmt.Sprintf("  %s", inParam),
					fmt.Sprintf(" and %s", inParam)))
		}

	}
	return
}
