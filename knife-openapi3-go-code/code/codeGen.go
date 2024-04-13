package code

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/mingcode100/knife4go/knife-openapi3-go-code/utils"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/template"
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
			ScanKnifeVueDist(dirPath+"/"+f.Name(), relativePath+"/"+f.Name(), f.Name(), outputPath+"/"+f.Name(), knifeArgs)
			continue
		}

		fname := strings.ReplaceAll(f.Name(), "@", "")

		templArgs := DistFileTemplArgs{
			PackageName: packageName,
			//HexContent:     getFileBase64(dirPath + "/" + fname),
			FileDir:        relativePath,
			FileRelavePath: relativePath + "/" + fname,
			FileType:       getFileType(fname),
			FileName:       fname,
			FileName2:      getFileName2(fname),
			FileName3:      getFileName3(fname),
		}

		if templArgs.FileType == "html" || templArgs.FileType == "css" {
			templArgs.HexContent = getFileContent(dirPath + "/" + f.Name())
		} else {
			templArgs.HexContent = getFileHexContent(dirPath + "/" + f.Name())
		}

		//fmt.Println("***********", templArgs.String())
		makeDistGoFile(templArgs, outputPath)
		//
		add2KnifeArgs(knifeArgs, templArgs)

	}
}

func add2KnifeArgs(args *KnifeArgs, args2 DistFileTemplArgs) {
	importAlian := args2.PackageName
	importPath := IMPORT_ROOT + args2.FileDir

	//fmt.Println(importAlian, " ", importPath)

	//解决包重名问题
	v, ok := args.KnifeImport[importAlian]
	if ok {
		if v != importPath {
			importAlian = importAlian + strconv.Itoa(i)
			i = i + 1
		}
	}

	args.KnifeImport[importAlian] = importPath

	args.Lines = append(args.Lines, KnifeLine{
		FileName3:    args2.FileName3,
		PackageAlian: importAlian,
	})

}

// 生成go文件
func makeDistGoFile(args DistFileTemplArgs, path string) {
	//TODO 生成go文件
	utils.CreateDirIfNotExists(path)

	f, err := os.OpenFile(path+"/"+args.FileName2+".go", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModeAppend|os.ModePerm)
	if nil != err {
		fmt.Errorf(" %s create file error: %v", path+"/"+args.FileName, err)
		return
	}

	defer f.Close()

	tmpl, err := template.ParseFiles("./templ/webfile.go.tmpl")
	if err != nil {
		panic(err)
	}

	//tmpl.Execute(os.Stdout, args)
	tmpl.Execute(f, args)

}

func getFileName2(fName string) string {
	s := strings.Replace(fName, ".", "_", -1)
	s = strings.Replace(s, "-", "_", -1)

	return strings.ToUpper(s)
}

func getFileName3(fName string) string {
	s := strings.Replace(fName, ".", "_", -1)
	s = strings.Replace(s, "-", "_", -1)
	sarr := strings.Split(s, "_")

	s1 := ""
	for _, v := range sarr {
		vv := []rune(v)
		if len(vv) > 0 {
			if vv[0] >= 'a' && vv[0] <= 'z' { //首字母大写
				vv[0] -= 32
			}
			s1 += string(vv)
		}
	}

	return s1
}

func getFileBase64(fpath string) string {
	bytes, err := ioutil.ReadFile(fpath)
	if nil != err {
		fmt.Errorf(" %s getFileBase64 error: %v", fpath, err)
		return ""
	}

	return base64.StdEncoding.EncodeToString(bytes)
}

func getFileContent(fpath string) string {
	bytes, err := ioutil.ReadFile(fpath)
	if nil != err {
		fmt.Errorf(" %s getFileBase64 error: %v", fpath, err)
		return ""
	}

	return string(bytes)
}

func getFileHexContent(fpath string) string {
	bytes, err := ioutil.ReadFile(fpath)
	if nil != err {
		fmt.Errorf(" %s getFileBase64 error: %v", fpath, err)
		return ""
	}

	return hex.EncodeToString(bytes)
}

func getFileType(fName string) string {

	sArr := strings.Split(fName, ".")

	return sArr[len(sArr)-1]
}
