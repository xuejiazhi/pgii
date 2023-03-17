package pg

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	flag "github.com/spf13/pflag"
	"os"
	"pgii/src/util"
	"strings"
)

var ColorGreenPrint *color.Color

// Host UserName PassWord Database Port (flag)
var Host = flag.StringP("host", "h", "127.0.0.1", "Input Your Postgresql Host")
var UserName = flag.StringP("user", "u", "postgres", "Input Your Postgresql User")
var PassWord = flag.StringP("password", "p", "123456", "Input Your Postgresql Password")
var Database = flag.StringP("db", "d", "postgres", "Input Your Postgresql databases")
var Port = flag.Int("port", 5432, "Input Your Postgresql Password")

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
	//print header
	fmt.Print(util.SetColor(fmt.Sprintf("pgii~[%s/%s]# ", *Database, P.Schema), util.LightBlue))

	//CMD
	cmdLine := ""

	//键盘输入
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//获取输入的值
		t := strings.Trim(scanner.Text(), "")
		//拼接到CMDLINE
		cmdLine += t
		if strings.HasSuffix(t, "\\") {
			//如果以 \ 结尾,继续输入
			fmt.Print(">")
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
}

// SetColor 设置颜色
func SetColor() {
	ColorGreenPrint = color.New()
	ColorGreenPrint.Add(color.FgHiGreen) // 绿色文字
}
