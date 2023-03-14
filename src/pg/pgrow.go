package pg

import "github.com/jedib0t/go-pretty/v6/table"

var (
	// TriggerHeader show header
	TriggerHeader = table.Row{"database", "schema", "trigger_name", "event_manipulation", "event_object_table", "action_orientation", "action_timing"}
	TableHeader   = table.Row{"Schema", "tablename", "tableowner", "tablespace"}
	ViewHeader    = table.Row{"Schema", "viewname", "viewowner"}
)
