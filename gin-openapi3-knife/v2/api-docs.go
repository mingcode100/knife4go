package v2

import (
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"
)

const (
	API_DOCS_RELATIVE_PATH = constant.ROOT_PATH + "/v3/api-docs"
)

// @param content string swag int 命令生成的swagger.json文件里的内容
func AddApiDocRouter(router *gin.Engine, swaggerJson string) {
	utils.GetJson(router, API_DOCS_RELATIVE_PATH, swaggerJson)
}
