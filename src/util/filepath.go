package util

import (
	"os"
)

// 判断文件夹是否存在
func HasDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

// CreateDir 创建文件夹
func CreateDir(path string) error {
	exist, err := HasDir(path)
	if err != nil {
		return err
	}
	if !exist {
		err = os.Mkdir(path, os.ModePerm)
	}
	return err
}
