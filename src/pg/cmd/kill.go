package cmd

import (
	"fmt"
	"github.com/spf13/cast"
	"pgii/src/pg/db"
	"pgii/src/pg/global"
	"pgii/src/util"
)

func (s *Params) Kill() {
	//参数长度为2位
	if len(s.Param) != global.TwoCMDLength {
		util.PrintColorTips(util.LightRed, global.KillProcessFailed)
		return
	}

	//查看DDL的类型
	sCmd := util.TrimLower(s.Param[0])
	switch sCmd {
	case "pid":
		s.KillPid()
		//case "deadlock":

	}
}

// KillDeadLock kill deadlock by database name
// todo:一键关闭所有死锁的进程
func (s *Params) KillDeadLock() {
	//判断PID是否存在
	if len(s.Param) != global.OneCMDLength {
		util.PrintColorTips(util.LightRed, global.SizeFailed)
		return
	}

}

// todo:KillPid KILL PID
func (s *Params) KillPid() {
	//判断PID是否存在
	if len(s.Param) != global.TwoCMDLength {
		util.PrintColorTips(util.LightRed, global.SizeFailed)
		return
	}

	//PID是否存在
	pid := cast.ToInt(s.Param[1])
	procInfo, err := db.P.GetProcessByPid(pid)
	if err != nil {
		util.PrintColorTips(util.LightRed, global.KillProcessFailed)
		return
	}

	if len(procInfo) == 0 {
		util.PrintColorTips(util.LightRed, fmt.Sprintf("%s pid not exists", global.KillProcessFailed))
		return
	}

	//kill pid
	if err := db.P.CancelProcessByPid(pid); err == nil {
		util.PrintColorTips(util.LightGreen, fmt.Sprintf("%s,pid[%d]", global.KillProcessSuccess, pid))
	} else {
		util.PrintColorTips(util.LightRed, global.KillProcessFailed)
	}
}
