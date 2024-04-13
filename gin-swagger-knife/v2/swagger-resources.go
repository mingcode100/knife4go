package v2

import (
	"github.com/gin-gonic/gin"
	"github.com/mingcode100/knife4go/gin-swagger-knife/constant"
	"github.com/mingcode100/knife4go/gin-swagger-knife/utils"
)

const (
	SWAGGER_RESOURCES_CONTENT       = `[{"location":"/v2/api-docs?group=2.X版本","name":"2.X版本","swaggerVersion":"2.0","url":"/v2/api-docs?group=2.X版本"}]`
	SWAGGER_RESOURCES_RELATIVE_PATH = constant.ROOT_PATH + "/swagger-resources"
)

func AddSwaggerResourcesRouter(router *gin.Engine) {
	utils.GetJson(router, SWAGGER_RESOURCES_RELATIVE_PATH, SWAGGER_RESOURCES_CONTENT)
}
