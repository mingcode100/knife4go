package templ

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"text/template"
	"knife-vue-2-go-code/utils"
	"os"
	"strconv"
	"strings"
)

var i = 0

// 扫描knife-vue-dist目录
func ScanKnifeVueDist(dirPath string, relativePath string, packageName string, outputPath string, knifeArgs *KnifeArgs) {
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
			ScanKnifeVueDist(dirPath+"/"+f.Name(), relativePath+"/"+f.Name(), f.Name(), outputPath+"/"+f.Name(), nil)
			continue
		}

		templArgs := DistFileTemplArgs{
			PackageName: packageName,
			//FileBase64:     getFileBase64(dirPath + "/" + f.Name()),
			FilePath:       relativePath,
			FileRelavePath: relativePath + "/" + f.Name(),
			FileType:       getFileType(f.Name()),
			FileName:       f.Name(),
			FileName2:      getFileName2(f.Name()),
		}

		//fmt.Println("***********", templArgs.String())
		makeDistGoFile(templArgs, outputPath)

		//add2KnifeArgs(knifeArgs, templArgs)

	}
}

func add2KnifeArgs(args *KnifeArgs, args2 DistFileTemplArgs) {
	importAlian := args2.PackageName
	importPath := "gin-swagger-knife/knifefiles" + args2.FilePath

	fmt.Println(importAlian, " ", importPath)
	v, ok := args.KnifeImport[importAlian]
	if ok {
		if v != importPath {
			importAlian = importAlian + strconv.Itoa(i)
			i = i + 1
		}
	}

	args.KnifeImport[importAlian] = importPath

	//append(args.KnifeLine, "")

}

//生成go文件
func makeDistGoFile(args DistFileTemplArgs, path string) {
	//TODO 生成go文件
	utils.CreateDirIfNotExists(path)

	f, err := os.OpenFile(path+"/"+args.FileName2+".go", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModeAppend|os.ModePerm)
	if nil != err {
		fmt.Errorf(" %s create file error: %v", path+"/"+args.FileName, err)
		return
	}

	defer f.Close()

	tmpl, err := template.New("WEBJARS_TEMPL").Parse(WEBJARS_TEMPL)
	if err != nil {
		panic(err)
	}

	tmpl.Execute(os.Stdout,args)
	//tmpl.Execute(f,args)

	//io.WriteString(f, args.FileBase64)

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
