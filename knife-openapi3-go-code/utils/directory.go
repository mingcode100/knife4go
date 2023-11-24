package utils

import (
	"fmt"
	"os"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: FileExists
//@description: 文件目录是否存在
//@param: path string
//@return: bool, error

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDir
//@description: 批量创建文件夹
//@param: dirs ...string
//@return: err error
func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := FileExists(v)
		if err != nil {
			return err
		}
		if !exist {
			//Logger.Debug("create directory" + v)
			err = os.MkdirAll(v, os.ModePerm)
		}
	}
	return err
}

func CreateDirIfNotExists(path string) {
	exist, err := FileExists(path)
	if nil != err {
		fmt.Errorf(" %s CreateDirIfNotExists error: %v", path, err)
	} else {
		if !exist {
			CreateDir(path)
		}
	}
}

func OpenFileForAppend(filePathName string) *os.File {
	exist, err := FileExists(filePathName)
	var f *os.File

	if nil != err {
		fmt.Errorf(" %s FileExists error: %v", filePathName, err)
		return nil
	} else {
		if !exist {
			f, err = os.OpenFile(filePathName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
		} else {

			f, err = os.OpenFile(filePathName, os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
		}
		if nil != err {
			fmt.Errorf(" %s OpenFile error: %v", filePathName, err)
			return nil
		}
	}

	return f

}
