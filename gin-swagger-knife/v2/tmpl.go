package v2

import (
	"gin-swagger-knife/constant"
	"gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"
)

const (
	BASE64        = `{{base64}}`
	RELATIVE_PATH = constant.ROOT_PATH + "{{rea}}}"
)

func AddRouter(router *gin.Engine) {
	utils.GetJson(router, BASE64, RELATIVE_PATH)
}
