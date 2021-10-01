package knife

import (
	v2 "gin-swagger-knife/v2"
	"github.com/gin-gonic/gin"


)

//@param content string swag int 命令生成的swagger.json文件里的内容
func InitSwaggerKnife(router *gin.Engine, swaggerJson string) {
	v2.AddApiDocRouter(router, swaggerJson)
	v2.AddSwaggerResourcesRouter(router)

}
