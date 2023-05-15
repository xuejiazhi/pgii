package util

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func If(condition bool, x, y interface{}) interface{} {
	if condition {
		return x
	}
	return y
}

func IfCmdFunc(condition bool, cmd string, param []string, x, y func(string, ...string)) {
	if condition {
		x(cmd, param...)
	} else {
		y(cmd, param...)
	}
}

func InArray(param string, params ...string) bool {
	for _, v := range params {
		if param == v {
			return true
		}
	}
	return false
}

func TrimLower(str string) string {
	return strings.Trim(strings.ToLower(str), "")
}

func RemoveNullStr(params *[]string) {
	flag := 0
	for {
		for k, v := range *params {
			if v == "" {
				*params = append((*params)[:k], (*params)[k+1:]...)
				break
			}

			if k == len(*params)-1 {
				flag = 1
			}
		}

		if flag == 1 {
			break
		}
	}
	return
}

func Substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}

func TypeTransForm(typename, fieldname string) (columnStr string) {
	switch strings.ToLower(typename) {
	case "timestamptz":
		columnStr = fmt.Sprintf("to_char(%s,'YYYY-MM-DD hh24:mi:ss') AS %s", fieldname, fieldname)
	case "timestamp":
		columnStr = fmt.Sprintf("to_char(%s,'YYYY-MM-DD hh24:mi:ss') AS %s", fieldname, fieldname)
	default:
		columnStr = fieldname
	}
	return
}

func String2Bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
