package util

import (
	"fmt"
	"strings"
)

type Color int

const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Purple  //紫色
	SeaBlue //海蓝
	Grey    //灰
	White
)

// 高亮
const (
	LightGrey  Color = iota + 90 //亮灰
	LightRed                     //亮红
	LightGreen                   //亮绿
	LightYellow
	LightBlue
	LightSeaBlue
	LightWhite
)

/*
*
设置颜色
*/
func SetColor(str string, color Color, para ...bool) string {
	prefix := ""
	//加粗
	if len(para) > 0 {
		if para[0] == true {
			prefix = "\u001b[1m"
		}
	}
	//下划线
	if len(para) > 1 {
		if para[1] == true {
			prefix = fmt.Sprintf("%s%s%s", prefix, "\u001b[4m", "\u001b[7m")
		}
	}

	return fmt.Sprintf("%s\u001B[%dm%s%s", prefix, color, str, "\u001B[0m")
}

func PrintColorTips(color Color, tips ...string) {
	//定义彩字
	colorStr := SetColor(strings.Join(tips, ""), color)
	fmt.Println(colorStr)
}
