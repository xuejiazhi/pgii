package pg

var (
	EqualAndFilter = []string{"equal", "eq", "filter", "fi"}
	EqualVar       = []string{"equal", "eq"}
	FilterVar      = []string{"filter", "fi"}
	TableAndView   = []string{"tb", "table", "view", "vw"}
	TableVar       = []string{"tb", "table"}
)

var (
	PgLimit = 10000 //每次
)
