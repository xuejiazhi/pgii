package pg

import (
	"fmt"
	"github.com/spf13/cast"
	"pgii/src/util"
)

func (s *Params) Kill() {
	//参数长度为2位
	if len(s.Param) != TwoCMDLength {
		util.PrintColorTips(util.LightRed, SizeFailed)
		return
	}

	//查看DDL的类型
	sCmd := util.TrimLower(s.Param[0])
	switch sCmd {
	case "pid":
		s.KillPid()
	}
}

// todo:KillPid KILL PID
func (s *Params) KillPid() {
	//判断PID是否存在
	if len(s.Param) != TwoCMDLength {
		util.PrintColorTips(util.LightRed, SizeFailed)
		return
	}

	//PID是否存在
	pid := cast.ToInt(s.Param[1])
	procInfo, err := P.GetProcessByPid(pid)
	if err != nil {
		util.PrintColorTips(util.LightRed, KillProcessFailed)
		return
	}

	if len(procInfo) == 0 {
		util.PrintColorTips(util.LightRed, fmt.Sprintf("%s pid not exists", KillProcessFailed))
		return
	}

	//kill pid
	if err := P.CancelProcessByPid(pid); err == nil {
		util.PrintColorTips(util.LightGreen, fmt.Sprintf("%s,pid[%d]", KillProcessSuccess, pid))
	} else {
		util.PrintColorTips(util.LightRed, KillProcessFailed)
	}
}
