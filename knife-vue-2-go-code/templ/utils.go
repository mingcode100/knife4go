package templ

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"
)
// 扫描knife-vue-dist目录
func ScanKnifeVueDist(dirPath string, relativePath string, packageName string, outputPath string) {
	files, err := ioutil.ReadDir(dirPath)
	if nil != err {
		fmt.Errorf("KNIFE_VUE_DIST_PATH error: %v", err)
		return
	}

	if nil == files || 0 == len(files) {
		return
	}

	for _, f := range files {
		if f.IsDir() {
			ScanKnifeVueDist(dirPath+"/"+f.Name(), relativePath+"/"+f.Name(), f.Name(), outputPath+"/"+f.Name())
		}

		templArgs := DistFileTemplArgs{
			PackageName: packageName,
			//FileBase64:     getFileBase64(dirPath + "/" + f.Name()),
			FileRelavePath: relativePath + "/" + f.Name(),
			FileType:       getFileType(f.Name()),
			FileName:       f.Name(),
			FileName2:      getFileName2(f.Name()),
		}

		fmt.Println("***********", templArgs.String())
		//TODO 生成go文件

	}
}

func getFileName2(fName string) string {
	s := strings.Replace(fName, ".", "_", -1)
	s = strings.Replace(s, "-", "_", -1)

	return strings.ToUpper(s)
}

func getFileBase64(fpath string) string {
	bytes, err := ioutil.ReadFile(fpath)
	if nil != err {
		fmt.Errorf(" %s getFileBase64 error: %v", fpath, err)
		return ""
	}

	return base64.StdEncoding.EncodeToString(bytes)
}

func getFileType(fName string) string {

	sArr := strings.Split(fName, ".")

	return sArr[len(sArr)-1]
}
