package pg

import (
	"os"
	"pgii/src/help"
	"pgii/src/util"
	"strings"
)

func Route(cmd string) {
	//命令行列表处理
	cmdStr := strings.Trim(cmd, "")
	///命令行split
	cmdList := strings.Split(cmdStr, " ")
	cmdRun := strings.ToLower(cmdList[0])

	//需要去掉多余空格的指令
	if util.InArray(cmdRun, []string{"show", "use", "desc", "help", "ddl"}) {
		util.RemoveNullStr(&cmdList)
	}

	//根据指令route到各个指令
	switch cmdRun {
	case "show":
		Show(cmdList[1:])
	case "use":
		Use(cmdList[1:])
	case "desc": //查看表结构
		Desc(cmdList[1:]...)
	case "help": //打印帮助
		help.Help(cmdList[1:]...)
	case "ddl": //查看建模式与表的语句
		DDL(cmdList[1:]...)
	case "size": //查看库和表的空间大小
		Size(cmdList[1:]...)
	case "dump": //备份数据
		Dump(cmdList[1:]...)
	case "exit":
		os.Exit(0)
	default:
		Psql(cmdRun, cmdStr)
	}
}
