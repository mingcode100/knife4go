package v2

import (
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"
)

const (
	SWAGGER_RESOURCES_CONTENT       = `[{"location":"/v3/api-docs?group=2.X版本","name":"2.X版本","swaggerVersion":"2.0","url":"/v2/api-docs?group=2.X版本"}]`
	SWAGGER_RESOURCES_RELATIVE_PATH = constant.ROOT_PATH + "/swagger-resources"
)

func AddSwaggerResourcesRouter(router *gin.Engine) {
	utils.GetJson(router, SWAGGER_RESOURCES_RELATIVE_PATH, SWAGGER_RESOURCES_CONTENT)
}
