package v3

import (
	"github.com/gin-gonic/gin"
	"github.com/mingcode100/knife4go/fizz-knife/constant"
	"github.com/mingcode100/knife4go/fizz-knife/utils"
)

const (
	// TODO 路径要改
	SWAGGER_RESOURCES_CONTENT     = `{"configUrl": "` + constant.ROOT_PATH + `/v3/api-docs/swagger-config","oauth2RedirectUrl": "` + constant.ROOT_PATH + `/swagger-ui/oauth2-redirect.html","url": "` + constant.ROOT_PATH + `/v3/api-docs","validatorUrl": ""}`
	SWAGGER_RESOURCES_CONFIG_PATH = constant.ROOT_PATH + "/v3/api-docs/swagger-config"
)

func AddSwaggerResourcesRouter(router *gin.Engine) {
	utils.GetJson(router, SWAGGER_RESOURCES_CONFIG_PATH, SWAGGER_RESOURCES_CONTENT)
}
