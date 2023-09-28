package pg

import (
	"bufio"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
	"pgii/src/pg/db"
	"pgii/src/pg/global"
	"pgii/src/util"
	"strings"
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
	//SetColor()

	//connect pgsql
	db.P.Host = *global.Host
	db.P.User = *global.UserName
	db.P.Password = *global.PassWord
	db.P.Port = *global.Port
	db.P.DataBase = *global.Database
	db.P.TimeZone = "Asia/Shanghai"
	if err := db.P.Connect(); err != nil {
		fmt.Println("Connect Pgsql Error:", err.Error())
	} else {
		//获取版本
		version, _ := db.P.Version()
		//欢迎信息
		WelCome(version)
		for {
			ReadLine()
		}
	}
}

func WelCome(v string) {
	fmt.Println(util.SetColor(fmt.Sprintf("Connect Pgsql Success Host %s", *global.Host), util.LightGreen))
	//todo
	fmt.Println(util.SetColor(fmt.Sprintf("PostgresSql Version: %s", v), util.LightGreen))
}

// ReadLine 获取键盘输入
func ReadLine() {
	//CMD
	cmdLine := ""
	//键盘输入
	//print header
	fmt.Print(util.SetColor(fmt.Sprintf("pgii~[%s/%s]# ", *global.Database, db.P.Schema), util.LightBlue))
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
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
				util.PrintColorTips(util.LightRed, global.CmdLineError)
			}
			break
		}
	}
}
