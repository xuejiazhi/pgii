package pg

import "github.com/jedib0t/go-pretty/v6/table"

var (
	ShowPrettyHeader = map[string]map[string][]interface{}{
		"cn": {
			VersionShowHeader:  []interface{}{"#", "版本"},
			TriggerShowHeader:  table.Row{"数据库", "模式", "触发器名称", "触发事件", "触发器所属表", "触发方向", "触发时刻"},
			TableShowHeader:    []interface{}{"模式", "表名", "表的拥用者", "表空间", "表大小", "超索引大小"},
			ViewShowHeader:     table.Row{"模式", "视图大小", "视图拥用者"},
			DatabaseShowHeader: []interface{}{"#oid", "数据库", "数据库拥用者", "字符编码", "LC_COLLATE", "LC_CTYPE", "允许连接", "最大并发连接数", "最后一个系统OID", "表空间", "数据库尺寸"},
			SchemaShowHeader:   []interface{}{"#oid", "模式名称", "模式拥有者", "权限"},
			ConnectionHeader:   []interface{}{"最大连接数", "超级用户保留的连接数", "剩余连接数", "当前正使用的连接数"},
			ProcessHeader:      []interface{}{"数据库pid", "数据库名称", "用户名", "客户端地址", "客户端端口", "应用程序", "状态"},
			//Size Show Header
			DatabaseSizeHeader: []interface{}{"数据库名称", "数据库尺寸"},
			TableSizeHeader:    []interface{}{"表名称", "表尺寸"},
			IndexSizeHeader:    []interface{}{"表名称", "索引尺寸"},
			TableSpaceHeader:   []interface{}{"表空间名称", "表空间尺寸"},

			//DESC Header
			DescTableHeader: []interface{}{"#", "列名", "数据类型", "长度", "是否为空", "默认值"},
		},
		"en": {
			VersionShowHeader:  []interface{}{"#", "Version"},
			TriggerShowHeader:  table.Row{"database", "schema", "trigger_name", "event_manipulation", "event_object_table", "action_orientation", "action_timing"},
			TableShowHeader:    []interface{}{"Schema", "tablename", "tableowner", "tablespace", "tablesize", "indexsize"},
			ViewShowHeader:     table.Row{"Schema", "viewname", "viewowner"},
			DatabaseShowHeader: []interface{}{"#oid", "DbName", "Auth", "Encoding", "LC_COLLATE", "LC_CTYPE", "AllowConn", "ConnLimit", "LastSysOid", "TableSpace", "size"},
			SchemaShowHeader:   []interface{}{"#oid", "SchemaName", "Owner", "Acl"},
			ConnectionHeader:   []interface{}{"max_connection", "superuser_reserved_connections", "remaining_connections", "inuse_connections"},
			ProcessHeader:      []interface{}{"pid", "database_name", "user_name", "client_addr", "client_port", "application_name", "state"},
			//Size Show Header
			DatabaseSizeHeader: []interface{}{"database", "database_size"},
			TableSizeHeader:    []interface{}{"tablename", "table_size"},
			IndexSizeHeader:    []interface{}{"tablename", "index_size"},
			TableSpaceHeader:   []interface{}{"tablespace_name", "tablespace_size"},

			//DESC Header
			DescTableHeader: []interface{}{"#", "column", "dataType", "length", "isnull", "defaultValue"},
		},
	}
	//Show
	VersionShowHeader  = "VersionShowHeader"
	TriggerShowHeader  = "TriggerShowHeader"
	TableShowHeader    = "TableShowHeader"
	ViewShowHeader     = "ViewShowHeader"
	DatabaseShowHeader = "DatabaseShowHeader"
	SchemaShowHeader   = "SchemaShowHeader"
	ConnectionHeader   = "ConnectionHeader"
	ProcessHeader      = "ProcessHeader"

	//Size Show Header
	DatabaseSizeHeader = "DatabaseSizeHeader"
	TableSizeHeader    = "TableSizeHeader"
	IndexSizeHeader    = "IndexSizeHeader"
	TableSpaceHeader   = "TableSpaceHeader"

	//DESC Header
	DescTableHeader = "DescTableHeader"
)
