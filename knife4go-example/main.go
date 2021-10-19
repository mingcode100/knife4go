package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	gin_swagger_knife "gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife"
)

func main() {
	router := gin.Default()

	swaggerJson := getFileContent("./swagger.json")


	router.Run(":8080")
}


func getFileContent(fpath string) string {
	bytes, err := ioutil.ReadFile(fpath)
	if nil != err {
		fmt.Errorf(" %s getFileBase64 error: %v", fpath, err)
		return ""
	}

	return string(bytes)
}