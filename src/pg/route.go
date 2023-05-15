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
	//命令行split
	cmdList := strings.Split(cmdStr, " ")
	cmdRun := strings.ToLower(cmdList[0])

	//需要去掉多余空格的指令
	if util.InArray(cmdRun, SystemCmd...) {
		util.RemoveNullStr(&cmdList)
	}

	//根据指令route到各个指令
	switch cmdRun {
	case ShowCMD:
		Handler(cmdList[1:]...).Show()
	case UseCMD:
		Handler(cmdList[1:]...).Use()
	case DescCMD: //查看表结构
		Handler(cmdList[1:]...).Desc()
	case HelpCMD: //打印帮助
		help.Help(cmdList[1:]...)
	case DdlCMD: //查看建模式与表的语句
		Handler(cmdList[1:]...).DDL()
	case SizeCMD: //查看库和表的空间大小
		Handler(cmdList[1:]...).Size()
	case DumpCMD: //备份数据
		Handler(cmdList[1:]...).Dump()
	case KillCMD:
		Handler(cmdList[1:]...).Kill()
	case SetCMD:
		Handler(cmdList[1:]...).Set()
	case LoadCMD:
		Handler(cmdList[1:]...).Load()
	case ExitCMD:
		os.Exit(0)
	default:
		Psql(cmdRun, cmdStr)
	}
}
