package pg

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"pgii/src/util"
	"strings"
)

var P PgDsn

// Connect 链接数据库
func (p *PgDsn) Connect() error {
	//Connect
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s database=%s  sslmode=disable",
		p.Host,
		p.Port,
		p.User,
		p.Password,
		p.DataBase,
	)

	if p.Schema != "" {
		dsn += fmt.Sprintf(" search_path=%s", p.Schema)
	}

	PgConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		fmt.Println("PostgresSQL  Connect Is failed! ")
		err = errors.New("PostgresSQL  Connect Is failed! ")
		return err
	}
	p.PgConn = PgConn
	return nil
}

// Version  show pg version
func (p *PgDsn) Version() (string, error) {
	//查询pgsql的版本
	//sql: show server_version;
	var version string
	if err := p.PgConn.Raw("show server_version").Row().Scan(&version); err == nil {
		return version, nil
	} else {
		return "", err
	}
}

// Database 获取数据库
func (p *PgDsn) Database() (pgDatabases []map[string]interface{}, err error) {
	/**
	  pg_database 数据库的信息表
		datname 数据库名字
		datdba 数据库的拥有者，通常是创建它的用户
		encoding 此数据库的字符编码的编号（pg_encoding_to_char()可将此编号转换成编码的名字）
		datcollate  此数据库的LC_COLLATE
		datctype  此数据库的LC_CTYPE
		datallowconn  如果为假则没有人能连接到这个数据库。这可以用来保护template0数据库不被修改。
		datconnlimit  设置能够连接到这个数据库的最大并发连接数。-1表示没有限制。
		datlastsysoid  数据库中最后一个系统OID，对pg_dump特别有用
		dattablespace  此数据库的默认表空间。在此数据库中，所有pg_class.reltablespace为0的表都将被存储在这个表空间中，尤其是非共享系统目录都会在其中。
		datacl   访问权限，更多信息参见 GRANT和 REVOKE
	*/
	sqlStr := `
		select 
			oid,
			datname,
			datdba,
			encoding,
			datcollate,
			datctype,
			datallowconn,
			datconnlimit,
			datlastsysoid,
			dattablespace, 
			(select pg_size_pretty( pg_database_size(datname) ) ) as size 
		from 
			pg_database`

	//query
	err = p.PgConn.Raw(sqlStr).Scan(&pgDatabases).Error

	return
}

func (p *PgDsn) Tables(cmd string, param ...string) (pgTables []map[string]interface{}, err error) {
	/**
	pg_tables 提供对数据库中每个表的信息的访问
		schemaname  包含表的模式名
		tablename   表名  	pg_class.relname
		tableowner  表拥有者的名字  pg_authid.rolname
		tablespace  包含表的表空间的名字（如果使用数据库的默认表空间，此列为空） pg_tablespace.spcname
	*/
	sqlStr := `
		select
			schemaname,
			tablename,
			tableowner,
			tablespace
		from 
			pg_tables 
`

	//是否选择Schema
	condition := p.getTableViewCondition("table", cmd, param...)

	if condition != "" {
		sqlStr += fmt.Sprintf(" where %s", condition)
	}
	//query
	err = p.PgConn.Raw(sqlStr).Scan(&pgTables).Error

	return
}

func (p *PgDsn) Views(cmd string, param ...string) (pgTables []map[string]interface{}, err error) {
	/**
	pg_tables 提供对数据库中每个表的信息的访问
		schemaname  包含表的模式名
		tablename   表名  	pg_class.relname
		tableowner  表拥有者的名字  pg_authid.rolname
		tablespace  包含表的表空间的名字（如果使用数据库的默认表空间，此列为空） pg_tablespace.spcname
	*/
	sqlStr := "select * from pg_views"

	//是否选择Schema
	condition := p.getTableViewCondition("view", cmd, param...)

	if condition != "" {
		sqlStr += fmt.Sprintf(" where %s", condition)
	}
	//query
	err = p.PgConn.Raw(sqlStr).Scan(&pgTables).Error

	return
}

func (p *PgDsn) Column(schemaName, tableName string) (tbColumns []map[string]interface{}, err error) {
	///Get Column TSQL
	if schemaName == "" {
		schemaName = p.Schema
	}
	sqlStr := fmt.Sprintf("select * from information_schema.columns where table_schema='%s' and table_name='%s'", schemaName, tableName)
	//query
	err = p.PgConn.Raw(sqlStr).Scan(&tbColumns).Error

	return
}

// Process 查看数据库列表
// all 参数 查看所有的proc; 不加只显示当前数据库下面的.
func (p *PgDsn) Process(param ...interface{}) (process []map[string]interface{}, err error) {
	//T-SQL
	var sqlStr string
	if len(param) > 0 {
		//sonCMD
		cmd := cast.ToString(param[0])
		switch cmd {
		case "all":
			sqlStr = "select pid,datname,usename,client_addr,client_port,application_name,state from pg_stat_activity"
		case "pid":
			if len(param) != ThreeCMDLength {
				err = errors.New("show proc pid param error")
				return
			} else {
				sqlStr = fmt.Sprintf(`
					select 
						pid,datname,usename,client_addr,client_port,application_name,state 
					from 
						pg_stat_activity 
					where 
						pid>=%d and pid<=%d`,
					cast.ToInt(param[1]), cast.ToInt(param[2]))
			}
		}
	} else {
		sqlStr = fmt.Sprintf(`
				select 
					pid,datname,usename,client_addr,client_port,application_name,state 
				from 
					pg_stat_activity 
				where datname='%s'`, P.DataBase)
	}
	//query
	err = p.PgConn.Raw(sqlStr).Scan(&process).Error
	//return
	return
}

func (p *PgDsn) Trigger(cmd, value string) (triggerInfo []map[string]interface{}, err error) {
	//T-SQL
	sqlStr := "select * from information_schema.triggers "

	//设置的条件值
	var conditionList []string

	//是否选择了DB
	//trigger_catalog
	if p.DataBase != "" {
		conditionList = append(conditionList, fmt.Sprintf("trigger_catalog='%s'", p.DataBase))
	}
	//是否选择了schema
	// column : trigger_schema
	if p.Schema != "" {
		conditionList = append(conditionList, fmt.Sprintf("trigger_schema='%s'", p.Schema))
	}

	// trigger_name 进行 filter 和 equal的操作
	if util.InArray(cmd, EqualAndFilter...) {

		//eq的处理
		if util.InArray(cmd, EqualVar...) {
			conditionList = append(conditionList, fmt.Sprintf("trigger_name ='%s'", value))
		}

		//filter的处理
		if util.InArray(cmd, FilterVar...) {
			conditionList = append(conditionList, fmt.Sprintf("trigger_name like '%%%s%%'", value))
		}
	}

	conditionStr := strings.Join(conditionList, " and ")

	if conditionStr != "" {
		sqlStr += fmt.Sprintf(" where %s", conditionStr)
	}

	//query
	err = p.PgConn.Raw(sqlStr).Scan(&triggerInfo).Error

	return
}

// SchemaNS 取Schema
func (p *PgDsn) SchemaNS() (pgSchema []map[string]interface{}, err error) {
	//query
	err = p.PgConn.Raw(
		fmt.Sprintf(`select oid,* from pg_namespace where nspname not in(%s)`,
			strings.Join(SystemSchemaList, ","))).
		Scan(&pgSchema).
		Error
	//return
	return
}

func (p *PgDsn) RunSQL(sqlStr string) (value []map[string]interface{}, err error) {
	//
	err = p.PgConn.Raw(sqlStr).Scan(&value).Error
	return
}

func (p *PgDsn) ExecSQL(sqlStr string) (affect int64, err error) {
	db := p.PgConn.Exec(sqlStr)
	affect = db.RowsAffected
	err = db.Error
	return
}

// ------------------------FUNCTION---------------------------------------//

func (p *PgDsn) GetEncodingChar(code int) string {
	var encoding string
	sqlStr := fmt.Sprintf("select pg_encoding_to_char(%d)", code)
	if err := p.PgConn.Raw(sqlStr).Row().Scan(&encoding); err == nil {
		return encoding
	} else {
		return ""
	}
}

func (p *PgDsn) GetRoleNameByOid(oid int) string {
	var rolName string
	sqlStr := fmt.Sprintf("select rolname from pg_authid where oid=%d", oid)
	if err := p.PgConn.Raw(sqlStr).Row().Scan(&rolName); err == nil {
		return rolName
	} else {
		return ""
	}
}

func (p *PgDsn) GetTableSpaceNameByOid(oid int) string {
	var spcName string
	sqlStr := fmt.Sprintf("select spcname from pg_tablespace where oid=%d", oid)
	if err := p.PgConn.Raw(sqlStr).Row().Scan(&spcName); err == nil {
		return spcName
	} else {
		return ""
	}
}

func (p *PgDsn) GetTableSpaceNameBySpcName(spcName string) (spcData []map[string]interface{}, err error) {
	sqlStr := fmt.Sprintf("select * from pg_tablespace where spcname='%s'", spcName)
	err = p.PgConn.Raw(sqlStr).Scan(&spcData).Error
	return
}

func (p *PgDsn) GetDatabaseInfoByName(dbName string) (pgDatabase map[string]interface{}, err error) {
	sqlStr := fmt.Sprintf(`
				select 
					oid,datname,datdba,encoding,datcollate,datctype,datallowconn,datconnlimit,datlastsysoid,dattablespace,datacl 
				from pg_database where datname ='%s'`, dbName)

	//query
	err = p.PgConn.Raw(sqlStr).Scan(&pgDatabase).Error

	return
}

func (p *PgDsn) GetSchemaFromNS(name string) (pgNameSpace map[string]interface{}, err error) {
	sqlStr := fmt.Sprintf("select oid,nspname,nspowner,nspacl from pg_namespace where nspname='%s'", name)
	//query
	err = p.PgConn.Raw(sqlStr).Scan(&pgNameSpace).Error

	return
}

// GetTableByName 获取表信息
// tbName 表名
func (p *PgDsn) GetTableByName(name string) (tbInfo map[string]interface{}, err error) {
	//TSQL
	sqlStr := fmt.Sprintf("select schemaname,tablename,tableowner,tablespace from pg_tables where tablename='%s'", name)
	//query
	err = p.PgConn.Raw(sqlStr).Scan(&tbInfo).Error

	return
}

func (p *PgDsn) GetTableBySchema(schema string) (pgTables []map[string]interface{}, err error) {
	/**
	pg_tables 提供对数据库中每个表的信息的访问
		schemaname  包含表的模式名
		tablename   表名  	pg_class.relname
		tableowner  表拥有者的名字  pg_authid.rolname
		tablespace  包含表的表空间的名字（如果使用数据库的默认表空间，此列为空） pg_tablespace.spcname
	*/
	if schema == "" {
		err = errors.New("schema is null")
		return
	}

	sqlStr := fmt.Sprintf(`
				select 
					schemaname,
					tablename,
					tableowner,
					tablespace
				from 
					pg_tables 
 				where schemaname ='%s'`, schema)

	//query
	err = p.PgConn.Raw(sqlStr).Scan(&pgTables).Error

	return
}

// GetPriMaryUniqueKey 获取主键和唯一约束
// tbName 表名
func (p *PgDsn) GetPriMaryUniqueKey(tbName string) (puInfo []map[string]interface{}, err error) {
	//TSQL
	sqlStr := fmt.Sprintf(`
		select
			k.column_name ,
			t.constraint_name ,
			t.constraint_type
		from
			information_schema.key_column_usage k
		left join 
			information_schema.table_constraints t
		on  k.constraint_name = t.constraint_name
		where
			k.table_catalog = '%s'
			and k.table_schema = '%s'
			and k.table_name = '%s'
			and t.table_catalog ='%s'
			and t.table_schema  ='%s'
			and t.table_name ='%s'
`, p.DataBase, p.Schema, tbName, p.DataBase, p.Schema, tbName)
	//query
	err = p.PgConn.Raw(sqlStr).Scan(&puInfo).Error

	return
}

// GetIndexDef 获取索引
// tbName 表名
func (p *PgDsn) GetIndexDef(tbName string, notIName []string) (indexInfo []map[string]interface{}, err error) {
	// indexdef  创建索引语句
	// indexname 索引名字
	//取创建索引的值
	sqlStr := fmt.Sprintf(`
			select
				b.indexrelid,
				a.indexdef,
				a.indexname 
			from
				PG_INDEXES a
			left join 
				PG_STAT_ALL_INDEXES b
			on  a.indexname = b.indexrelname 
			where
				a.schemaname ='%s'
			and
				a.tablename ='%s'
			and 
				b.schemaname = '%s' 
			and 
				b.relname  ='%s'`,
		p.Schema, tbName, p.Schema, tbName)

	if len(notIName) > 0 {
		for k, v := range notIName {
			notIName[k] = fmt.Sprintf("'%s'", v)
		}
		//增加条件
		sqlStr += fmt.Sprintf(" and a.indexname not in(%s)", strings.Join(notIName, ","))
		sqlStr += fmt.Sprintf(" and b.indexrelname not in(%s)", strings.Join(notIName, ","))
	}

	sqlStr += " order by b.indexrelid desc"
	//fmt.Println(sqlStr)
	//query
	err = p.PgConn.Raw(sqlStr).Scan(&indexInfo).Error

	return
}

// GetTriggerByTgRelid 获取触发器
func (p *PgDsn) GetTriggerByTgRelid(tgrelid int) (triggerInfo []map[string]interface{}, err error) {
	//T-SQL
	sqlStr := fmt.Sprintf("SELECT oid,* FROM pg_trigger where tgrelid =%d", tgrelid)

	//query
	err = p.PgConn.Raw(sqlStr).Scan(&triggerInfo).Error

	//return
	return
}
