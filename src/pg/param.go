package pg

const (
	ZeroCMDLength = iota
	OneCMDLength
	TwoCMDLength
	ThreeCMDLength
	FourCMDLength
	FiveCMDLength
)

const (
	NoneStyle = iota
	DatabaseStyle
	TableStyle
	IndexStyle
	ViewStyle
	SelectStyle
	SchemaStyle
	TriggerStyle
	VersionStyle
	ConnectionStyle
	ProcessStyle
	TableSpaceStyle //表空间
)

const (
	MaxConnections               = iota //最大连接数
	SuperuserReservedConnections        //超级用户保留的连接数
	RemainingConnections                //剩余连接数
	InUseConnections                    //正在使用的链接数
)

const (
	DDL = iota
	DUMP
)

var (
	//CMD
	ShowCMD    = "show"
	UseCMD     = "use"
	DescCMD    = "desc"
	HelpCMD    = "help"
	DdlCMD     = "ddl"
	DumpCMD    = "dump"
	SizeCMD    = "size"
	KillCMD    = "kill"
	ExplainCMD = "explain"
	SetCMD     = "set"
	LoadCMD    = "load"
	ExitCMD    = "exit"

	//config set
	DefaultHost     = "127.0.0.1"
	DefaultUser     = "postgres"
	DefaultPassword = "123456"
	DefaultDB       = "postgres"
	DefaultPort     = 5432

	//array
	EqualAndFilter   = []string{"equal", "eq", "filter", "fi"}
	EqualVar         = []string{"equal", "eq"}
	FilterVar        = []string{"filter", "fi"}
	TableAndView     = []string{"tb", "table", "view", "vw"}
	TableVar         = []string{"tb", "table"}
	SystemSchemaList = []string{"'pg_toast'", "'pg_temp_1'", "'pg_toast_temp_1'", "'pg_catalog'", "'information_schema'"}
	SystemCmd        = []string{ShowCMD, UseCMD, DescCMD, HelpCMD, DdlCMD, DumpCMD, SizeCMD, KillCMD, ExplainCMD}

	//get column limit
	PgLimit = 50000

	//int type
	Int2Type = "int2"
	Int4Type = "int4"
	Int8Type = "int8"

	ZhCN    = "cn"
	ZhEN    = "en"
	INIFile = "_init_"
)
