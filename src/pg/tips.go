package pg

var (
	// Dump
	DumpFailed        = "Failed:Dump Cmd is failed"
	DumpFailedNoTable = "Failed:Dump Cmd Table fail,Table not exists!"

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
	DDLTableNoExists  = "Failed:DDL Cmd Table fail,Table not exists!"
	DDLColumnNoExists = "Failed:DDL Cmd Table fail,Column not exists!"
)
