package pg

var (
	// LineOperate CloseFileFailed COMMON
	LineOperate     = "----------------------------------------"
	CloseFileFailed = "Close File Failed"

	// DumpFailed Dump
	DumpFailed                 = "Failed:Dump Cmd is failed"
	DumpFailedNoTable          = "Failed:Dump Cmd Table fail,Table not exists!"
	DumpFailedNoSelectSchema   = "Failed:Dump Cmd Schema fail,Schema not Selected!"
	DumpFailedNoSelectDatabase = "Failed:Dump Cmd Database fail,Database not Selected!"
	DumpDatabaseFailedNoSchema = "Failed:Dump Cmd Database fail,No Schema!"
	DumpFailedSchemaNoTable    = "tips:Dump Cmd Table fail,no table in Schema!"
	DumpTableSuccess           = "Dump Table Success"
	DumpTableStructSuccess     = "Dump Table Struct Success"
	DumpTableRecordSuccess     = "Dump Table Record Success"
	DumpSchemaSuccess          = "Dump Schema Success"
	DumpTableNotExists         = "Dump Table Not Exists"
	DumpSchemaNotExists        = "Dump Schema Not Exists"
	DumpDataBaseBegin          = "Dump DataBase Begin"
	DumpDataBaseStructSuccess  = "Dump DataBase Struct Success"

	// LoadFailed Load
	LoadFailed               = "Failed:Load Cmd is failed"
	LoadTableNOFile          = "Failed:Table Pgi File is not exists"
	LoadNoFile               = "cannot read the file"
	LoadTableExecSQLFailed   = "Failed:Table Exec SQL Failed"
	LoadTableSQLSuccess      = "Load Table Success"
	LoadFailedNoSelectSchema = "Failed:Load Cmd Schema fail,Schema not Selected!"

	// SizeFailed Size
	SizeFailed           = "Failed:SIZE Cmd fail"
	SizeFailedNull       = "Failed:Size Database is Nil!"
	SizeFailedDataNull   = "Failed:Size Database Get Data Nil!"
	SizeFailedPointTable = "Failed:Size Table Must Point Table Name!"
	SizeFailedNoSchema   = "Failed:Size Cmd Schema fail,Schema not exists!"
	SizeFailedNoTable    = "Failed:Size Cmd Table fail,Table not exists!"

	// UseFailed USE
	UseFailed          = "Failed:Use Cmd is failed"
	UseDBFailed        = "Failed:Use Database failed"
	UseDBNotExists     = "Failed:Use Database fail,DataBase Not Exists!"
	UseDBSuch          = "Use Database Success!"
	UseSchemaFailed    = "Failed:Use Schema fail!"
	UseSchemaNotExists = "Failed:Use Schema fail,Schema Not Exists!"
	UseSchemaSuch      = "Use Schema Success!"

	// DDLTableNoExists DDL
	DDLTableNoExists   = "Failed:DDL Cmd Table fail,Table not exists!"
	DDLColumnNoExists  = "Failed:DDL Cmd Table fail,Column not exists!"
	DDLViewNoExists    = "Failed:DDL Cmd View fail,View not exists!"
	DDLTableError      = "Failed:DDL Cmd Table fail,error "
	DDLViewError       = "Failed:DDL Cmd View fail,error "
	DDLSchemaError     = "Failed:DDL Cmd Schema fail,error "
	DDLSchemaNotExists = "Failed:DDL Cmd Schema fail,Schema not exists!"

	// ShowTriggerCmdFailed SHOW
	ShowTriggerCmdFailed = "Failed Show Trigger CMD is Failed"
	ShowDatabaseError    = "Failed:Show DataBase is Wrong! error"
	CmdLineError         = "CmdLine Must be With ; ending"
	CmdLineWrong         = "Failed:CmdLine is Wrong!"
	StartThanEndError    = "Failed:End Pid Must Than Start Pid"

	// DescTableError DESC
	DescTableError    = "Failed:Describe Table fail!Error"
	DescTableNoExists = "Failed:Describe Table fail!Table not exists"
	DescTableFailed   = "Failed:Describe Table fail"

	// KillProcessSuccess KILL
	KillProcessSuccess = "Kill Process Success"
	KillProcessFailed  = "Kill Process Failed"

	// SetError Set
	SetError           = "Set Cmd fail"
	SetLanguageSuccess = "Set Language Success"

	// CmdTipsHost cmd
	CmdTipsHost     = "Input Your Postgresql Host"
	CmdTipsUser     = "Input Your Postgresql User"
	CmdTipsPassword = "Input Your Postgresql Password"
	CmdTipsDatabase = "Input Your Postgresql database"
	CmdTipsPort     = "Input Your Postgresql Port"
)
