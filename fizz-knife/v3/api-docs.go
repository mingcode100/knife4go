package v3

import (
	"gitee.com/youbeiwuhuan/knife4go/fizz-knife/constant"
	"github.com/gin-gonic/gin"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
)

const (
	// TODO 路径要改
	API_DOCS_RELATIVE_PATH = constant.ROOT_PATH + "/v3/api-docs"
)

func AddApiDocRouter(router *gin.Engine, infos *openapi.Info) *fizz.Fizz {
	fizz := fizz.NewFromEngine(router)

	// Create a new route that serve the OpenAPI spec.
	fizz.GET(API_DOCS_RELATIVE_PATH, nil, fizz.OpenAPI(infos, "json"))

	//utils.GetJson(router, API_DOCS_RELATIVE_PATH, swaggerJson)

	return fizz
}
