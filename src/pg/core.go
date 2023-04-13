package pg

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	flag "github.com/spf13/pflag"
	"pgii/src/util"
	"strings"
)

var (
	ColorGreenPrint *color.Color
	Host            = flag.StringP("host", "h", DefaultHost, CmdTipsHost)
	UserName        = flag.StringP("user", "u", DefaultUser, CmdTipsUser)
	PassWord        = flag.StringP("password", "p", DefaultPassword, CmdTipsPassword)
	Database        = flag.StringP("db", "d", DefaultDB, CmdTipsDatabase)
	Port            = flag.Int("port", DefaultPort, CmdTipsPort)
	Language        = "en"
)

func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return flag.NormalizedName(name)
}

func Run() {
	// 设置标准化参数名称的函数
	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)
	flag.Parse()

	//设置颜色
	SetColor()

	//connect pgsql
	P.Host = *Host
	P.User = *UserName
	P.Password = *PassWord
	P.Port = *Port
	P.DataBase = *Database
	P.TimeZone = "Asia/Shanghai"
	if err := P.Connect(); err != nil {
		fmt.Println("Connect Pgsql Error:", err.Error())
	} else {
		//获取版本
		version, _ := P.Version()
		//欢迎信息
		WelCome(version)
		for {
			ReadLine()
		}
	}
}

func WelCome(v string) {
	fmt.Println(util.SetColor(fmt.Sprintf("Connect Pgsql Success Host %s", *Host), util.LightGreen))
	//todo
	fmt.Println(util.SetColor(fmt.Sprintf("PostgresSql Version: %s", v), util.LightGreen))
}

// ReadLine 获取键盘输入
func ReadLine() {
	//t := prompt.Input(util.SetColor(fmt.Sprintf("pgi~[%s/%s]# ", *Database, P.Schema), util.LightBlue), completer)
	//fmt.Println("You selected " + t)
	//
	////print header
	//fmt.Print(util.SetColor(fmt.Sprintf("pgi~[%s/%s]# ", *Database, P.Schema), util.LightBlue))
	//CMD
	cmdLine := ""
	//获取输入的值
	t := prompt.Input(fmt.Sprintf("pgii~[%s/%s]# ", *Database, P.Schema), completer)
	for {
		cmdLine += t
		if strings.HasSuffix(t, "\\") {
			//如果以 \ 结尾,继续输入
			//fmt.Print(">")
			t = prompt.Input(">", completer)
		} else {
			//使用;结束
			if strings.HasSuffix(t, ";") {
				//去掉 \和最后的 ;
				cmdStr := strings.Replace(cmdLine, "\\", " ", -1)
				cmdLine = util.Substring(cmdStr, 0, len(cmdStr)-1)
				//去掉 ;
				Route(cmdLine)
				break
			}
			//wrong
			if strings.Trim(cmdLine, "") != "" {
				util.PrintColorTips(util.LightRed, CmdLineError)
			}
			break
		}
	}

	//键盘输入
	//scanner := bu-fio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//
	//	//t := strings.Trim(scanner.Text(), "")
	//	//拼接到CMDLINE
	//	cmdLine += t
	//	if strings.HasSuffix(t, "\\") {
	//		//如果以 \ 结尾,继续输入
	//		fmt.Print(">")
	//	} else {
	//		//使用;结束
	//		if strings.HasSuffix(t, ";") {
	//			//去掉 \和最后的 ;
	//			cmdStr := strings.Replace(cmdLine, "\\", " ", -1)
	//			cmdLine = util.Substring(cmdStr, 0, len(cmdStr)-1)
	//			//去掉 ;
	//			Route(cmdLine)
	//			break
	//		}
	//		//wrong
	//		if strings.Trim(cmdLine, "") != "" {
	//			util.PrintColorTips(util.LightRed, CmdLineError)
	//		}
	//		break
	//	}
	//}
}

// SetColor 设置颜色
func SetColor() {
	ColorGreenPrint = color.New()
	ColorGreenPrint.Add(color.FgHiGreen) // 绿色文字
}
