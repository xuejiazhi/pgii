package pg

var (
	//COMMON

	// Dump
	DumpFailed               = "Failed:Dump Cmd is failed"
	DumpFailedNoTable        = "Failed:Dump Cmd Table fail,Table not exists!"
	DumpFailedNoSelectSchema = "Failed:Dump Cmd Table fail,Schema not Selected!"
	DumpFailedSchemaNoTable  = "Failed:Dump Cmd Table fail,Schema not Selected!"
	DumpTableSuccess         = "Dump Table Success"
	DumpTableStructSuccess   = "Dump Table Struct Success"
	DumpTableRecordSuccess   = "Dump Table Record Success"
	DumpSchemaSuccess        = "Dump Schema Success"
	DumpTableNotExists       = "Dump Table Not Exists"

	// Size
	SizeFailed           = "Failed:SIZE Cmd fail"
	SizeFailedNull       = "Failed:Size Database is Nil!"
	SizeFailedDataNull   = "Failed:Size Database Get Data Nil!"
	SizeFailedPointTable = "Failed:Size Table Must Point Table Name!"
	SizeFailedNoSchema   = "Failed:Size Cmd Schema fail,Schema not exists!"
	SizeFailedNoTable    = "Failed:Size Cmd Table fail,Table not exists!"

	//USE
	UseFailed          = "Failed:Use Cmd is failed"
	UseDBFailed        = "Failed:Use Database failed"
	UseDBNotExists     = "Failed:Use Database fail,DataBase Not Exists!"
	UseDBSucc          = "Use Database Success!"
	UseSchemaFailed    = "Failed:Use Schema fail!"
	UseSchemaNotExists = "Failed:Use Schema fail,Schema Not Exists!"
	UseSchemaSucc      = "Use Schema Success!"

	//DDL
	DDLTableNoExists   = "Failed:DDL Cmd Table fail,Table not exists!"
	DDLColumnNoExists  = "Failed:DDL Cmd Table fail,Column not exists!"
	DDLViewNoExists    = "Failed:DDL Cmd View fail,View not exists!"
	DDLTableError      = "Failed:DDL Cmd Table fail,error "
	DDLViewError       = "Failed:DDL Cmd View fail,error "
	DDLSchemaError     = "Failed:DDL Cmd Schema fail,error "
	DDLSchemaNotExists = "Failed:DDL Cmd Schema fail,Schema not exists!"

	//SHOW
	ShowTriggerCmdFailed = "Failed Show Trigger CMD is Failed"
	ShowDatabaseError    = "Failed:Show DataBase is Wrong! error"
	CmdLineError         = "CmdLine Must be With ; ending"
	CmdLineWrong         = "Failed:CmdLine is Wrong!"
	StartThanEndError    = "Failed:End Pid Must Than Start Pid"

	//DESC
	DescTableError    = "Failed:Describe Table fail!Error"
	DescTableNoExists = "Failed:Describe Table fail!Table not exists"
	DescTableFailed   = "Failed:Describe Table fail"

	//KILL
	KillProcessSuccess = "Kill Process Success"
	KillProcessFailed  = "Kill Process Failed"
)
