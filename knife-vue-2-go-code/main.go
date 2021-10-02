package main

import (
	"knife-vue-2-go-code/templ"
)

const (
	KNIFE_VUE_DIST_PATH = "../knife-vue-dist"
	OUTPUT_PATH         = "../gin-swagger-knife/knifefiles"
)

func main() {

	knifeArgs := &templ.KnifeArgs{
		KnifeImport:make(map[string]string,100),
		KnifeLine:make([]string,0,100),
	}

	templ.ScanKnifeVueDist(KNIFE_VUE_DIST_PATH, "", "knifefiles", OUTPUT_PATH,knifeArgs)

}
