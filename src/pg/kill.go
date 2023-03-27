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
		fmt.Println()
		return
	}

	if len(procInfo) == 0 {
		fmt.Println()
		return
	}

}
