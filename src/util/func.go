package util

import "strings"

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

func InArray(param string, params []string) bool {
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
