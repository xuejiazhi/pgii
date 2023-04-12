package pg

type HandlerInterface interface {
	SizeInterface
	ShowInterface
	UserInterface
	DescInterface
	DDLInterface
	DumpInterface
	KillInterface
	SetInterface
	LoadInterface
}

func Handler(param ...string) HandlerInterface {
	return HandlerInterface(getInstance(param...))
}

func getInstance(param ...string) *Params {
	return &Params{
		Param: param,
	}
}

type Params struct {
	Param []string
}

// SizeInterface Size指令 Interface
type SizeInterface interface {
	Size()
	SizeIndex()
	SizeDatabase(...string)
	SizeTable([]string)
}

// ShowInterface show 指令
type ShowInterface interface {
	Show()
	ShowTrigger()
	ShowSchema()
	ShowVersion()
	ShowDatabases()
	ShowTableView(string)
	ShowTables(string, ...string)
	ShowView(string, ...string)
	ShowConnection()
	ShowProcess()
}

type UserInterface interface {
	Use()
	UseDatabase(string)
	UseSchema(string)
}

type DescInterface interface {
	Desc()
}

type DDLInterface interface {
	DDL()
	DDLSchema(string)
	DDLTable(string)
	DDLView(string)
}

type DumpInterface interface {
	Dump()
	DumpSchema()
	DumpTable()
	DumpDatabase()
}

type KillInterface interface {
	Kill()
}

type SetInterface interface {
	Set()
	SetLanguage()
}

type LoadInterface interface {
	Load()
	LoadTable(string)
	LoadSchema(string)
	LoadDataBase(string)
}
