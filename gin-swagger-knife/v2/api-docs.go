package v2

import (
	"github.com/gin-gonic/gin"
	"github.com/mingcode100/knife4go/gin-swagger-knife/constant"
	"github.com/mingcode100/knife4go/gin-swagger-knife/utils"
)

const (
	API_DOCS_RELATIVE_PATH = constant.ROOT_PATH + "/v2/api-docs"
)

// @param content string swag int 命令生成的swagger.json文件里的内容
func AddApiDocRouter(router *gin.Engine, swaggerJson string) {
	utils.GetJson(router, API_DOCS_RELATIVE_PATH, swaggerJson)
}
