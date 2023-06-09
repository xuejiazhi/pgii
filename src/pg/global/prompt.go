package global

import (
	prompt "github.com/c-bata/go-prompt"
)

var suggestions = []prompt.Suggest{
	// Command
	{"show", ""},
	{"show ver", ""},
	{"show version", ""},
	{"show db", ""},
	{"show database", ""},
	{"show tb", ""},
	{"show table", ""},
	{"show vw", ""},
	{"show view", ""},
	{"show sd", ""},
	{"show selectdb", ""},
	{"show sc", ""},
	{"show schema", ""},
	{"show tg", ""},
	{"show trigger", ""},
	{"show conn", ""},       //"conn", "connection":
	{"show connection", ""}, //"conn", "connection":
	{"show proc", ""},       //"proc", "process"
	{"show process", ""},    //"proc", "process"
	{"use", ""},
	{"use db", ""},
	{"use database", ""},
	{"use sc", ""},
	{"use schema", ""},
	{"desc", ""},
	{"ddl", ""},
	{"ddl sc", ""},
	{"ddl schema", ""},
	{"ddl tb", ""},
	{"ddl table", ""},
	{"ddl vw", ""},
	{"ddl view", ""},
	{"size", ""},
	{"size db", ""},
	{"size database", ""},
	{"size tb", ""},
	{"size table", ""},
	{"size idx", ""},
	{"size index", ""},
	{"size tablespace", ""}, //"tablespace", "tbsp"
	{"size tbsp", ""},
	{"dump", ""},
	{"dump tb", ""},
	{"dump table", ""},
	{"dump db", ""},
	{"dump database", ""},
	{"dump sc", ""},
	{"dump schema", ""},
	{"kill", ""},
	{"kill pid", ""},
	{"set", ""},
	{"set language", ""},
	{"exit", ""},
}

func completer(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasSuffix(suggestions, d.GetWordBeforeCursor(), true)
}
