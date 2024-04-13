package main

import (
	"fmt"
	"github.com/mingcode100/knife4go/knife-openapi3-go-code/code"
	"github.com/mingcode100/knife4go/knife-openapi3-go-code/utils"
	"os"
	"text/template"
)

func main() {

	knifeArgs := &code.KnifeArgs{
		KnifeImport: make(map[string]string, 100),
		Lines:       make([]code.KnifeLine, 0, 100),
	}

	fmt.Println("==================================开始生成knife-openapi3代码")
	code.ScanKnifeVueDist(code.KNIFE_VUE_DIST_PATH, code.ROOT_RELATE_PATH, code.ROOT_PACKAGE, code.OUTPUT_PATH, knifeArgs)
	makeKnifeFile(knifeArgs, code.KNIFE_GO_PATH)
	fmt.Println("==================================生成完成")

}

// 生成go文件
func makeKnifeFile(args *code.KnifeArgs, path string) {
	utils.CreateDirIfNotExists(path)

	f, err := os.OpenFile(path+"/"+code.KNIFE_GO_NAME, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModeAppend|os.ModePerm)
	if nil != err {
		fmt.Errorf(" %s create file error: %v", path+"/"+code.KNIFE_GO_NAME, err)
		return
	}

	defer f.Close()

	tmpl, err := template.ParseFiles("./templ/knife.go.tmpl")
	if err != nil {
		panic(err)
	}

	//tmpl.Execute(os.Stdout, args)
	tmpl.Execute(f, args)

}
