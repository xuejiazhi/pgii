package pg

import (
	"os"
	"pgii/src/help"
	cmd2 "pgii/src/pg/cmd"
	"pgii/src/pg/global"
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
	if util.InArray(cmdRun, global.SystemCmd...) {
		util.RemoveNullStr(&cmdList)
	}

	//根据指令route到各个指令
	switch cmdRun {
	case global.ShowCMD:
		cmd2.Handler(cmdList[1:]...).Show()
	case global.UseCMD:
		cmd2.Handler(cmdList[1:]...).Use()
	case global.DescCMD: //查看表结构
		cmd2.Handler(cmdList[1:]...).Desc()
	case global.HelpCMD: //打印帮助
		help.Help(cmdList[1:]...)
	case global.DdlCMD: //查看建模式与表的语句
		cmd2.Handler(cmdList[1:]...).DDL()
	case global.SizeCMD: //查看库和表的空间大小
		cmd2.Handler(cmdList[1:]...).Size()
	case global.DumpCMD: //备份数据
		cmd2.Handler(cmdList[1:]...).Dump()
	case global.KillCMD:
		cmd2.Handler(cmdList[1:]...).Kill()
	case global.SetCMD:
		cmd2.Handler(cmdList[1:]...).Set()
	case global.LoadCMD:
		cmd2.Handler(cmdList[1:]...).Load()
	case global.ClearCMD:
		cmd2.Handler(cmdList[1:]...).Clear()
	case global.ExitCMD:
		os.Exit(0)
	default:
		cmd2.Psql(cmdRun, cmdStr)
	}
}
