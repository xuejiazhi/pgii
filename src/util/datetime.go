package util

import (
	"fmt"
	"github.com/spf13/cast"
	"time"
)

var timeLayout = "2006-01-02 15:04:05"

func GetFormatDateTime() string {
	timeUnix := time.Now().UnixNano() / 1e6 //已知的时间戳
	miroTime := timeUnix - timeUnix/1000*1000
	formatTimeStr := time.Unix(timeUnix/1000, 0).Format(timeLayout)
	//fmt.Println(formatTimeStr) //打印结果：2017-04-11 13:30:39
	return fmt.Sprintf("%s.%s", formatTimeStr, cast.ToString(miroTime))
}

func ParseDate(strTime string) time.Time {
	dateTime, _ := time.Parse(timeLayout, strTime)
	return dateTime
}

func GetTimeArr(start, end string) int64 {
	// 转成时间戳
	startUnix, _ := time.ParseInLocation(timeLayout, start, time.Local)
	endUnix, _ := time.ParseInLocation(timeLayout, end, time.Local)
	startTime := startUnix.Unix()
	endTime := endUnix.Unix()
	// 求相差天数
	date := (endTime - startTime) / 86400
	return date
}

func GetTimeDiff(datestr string, d, p int) string {
	//转化时间
	timeLocal, _ := time.ParseInLocation(timeLayout, datestr, time.Local)

	//时间戳
	timeUnix := timeLocal.Unix()

	//judgement
	if p == 1 {
		timeUnix += 86400 * int64(d)
	} else {
		timeUnix -= 86400 * int64(d)
	}

	//return
	return time.Unix(timeUnix, 0).Format(timeLayout)
}

// AddTimeSecond
func AddTimeSecond(datestr string, d, p int) string {
	//转化时间
	timeLocal, _ := time.ParseInLocation(timeLayout, datestr, time.Local)
	//时间戳
	timeUnix := timeLocal.Unix()

	//judgement
	if p == 1 {
		timeUnix += int64(d) * 60
	} else {
		timeUnix -= int64(d) * 60
	}

	//return
	return time.Unix(timeUnix, 0).Format(timeLayout)
}

func ConvUnixTime(datetime int64) string {
	length := len(fmt.Sprintf("%d", datetime))
	if length == 10 {
		return time.Unix(datetime, 0).Format(timeLayout)
	} else if length == 13 {
		return time.Unix(datetime/1000, 0).Format(timeLayout)
	} else {
		return ""
	}
}
