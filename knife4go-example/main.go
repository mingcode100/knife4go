package main

import (
	"fmt"
	gin_swagger_knife "gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()
	// 获取 swag int 命令生成的swagger.json文件里的内容
	swaggerJson := getFileContent("./swagger.json")
	fmt.Println(swaggerJson)
	gin_swagger_knife.InitSwaggerKnife(router,swaggerJson)
	router.GET("/hello",func(ctx *gin.Context) {
		rs := []byte("dafafdasfaf")

		ctx.Status(http.StatusOK)
		ctx.Header("content-type", "application/json;charset=UTF-8")
		ctx.Header("content-length", strconv.Itoa(len(rs)))
		ctx.Header("connection", "keep-alive")

		_, err := ctx.Writer.Write(rs)
		if nil != err {
			log.Fatal(err)
			return
		}

		ctx.Writer.Flush()
	})

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