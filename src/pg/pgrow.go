package pg

import "github.com/jedib0t/go-pretty/v6/table"

var (
	//Show
	VersionShowHeader  = []interface{}{"#", "Version"}
	TriggerShowHeader  = table.Row{"database", "schema", "trigger_name", "event_manipulation", "event_object_table", "action_orientation", "action_timing"}
	TableShowHeader    = []interface{}{"Schema", "tablename", "tableowner", "tablespace", "tablesize", "indexsize"}
	ViewShowHeader     = table.Row{"Schema", "viewname", "viewowner"}
	DatabaseShowHeader = []interface{}{"#oid", "DbName", "Auth", "Encoding", "LC_COLLATE", "LC_CTYPE", "AllowConn", "ConnLimit", "LastSysOid", "TableSpace", "size"}
	SchemaShowHeader   = []interface{}{"#oid", "SchemaName", "Owner", "Acl"}
	ConnectionHeader   = []interface{}{"max_connection", "superuser_reserved_connections", "remaining_connections", "inuse_connections"}

	//Size Show Header
	DatabaseSizeHeader = []interface{}{"database", "database_size"}
	TableSizeHeader    = []interface{}{"tablename", "table_size"}
	IndexSizeHeader    = []interface{}{"tablename", "index_size"}

	//DESC Header
	DescTableHeader = []interface{}{"#", "column", "dataType", "length", "isnull", "defaultValue"}
)
