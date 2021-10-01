package main

import (
	"knife-vue-2-go-code/templ"
)

const (
	KNIFE_VUE_DIST_PATH = "../knife-vue-dist"
	OUTPUT_PATH         = "../gin-swagger-knife/webjars"
)

func main() {
	//knifeArgsList := make([]templ.KnifeArgs, 20, 20)

	templ.ScanKnifeVueDist(KNIFE_VUE_DIST_PATH, "", "webjars", OUTPUT_PATH)

}
