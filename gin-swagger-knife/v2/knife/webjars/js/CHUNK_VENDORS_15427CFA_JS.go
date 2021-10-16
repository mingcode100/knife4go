package js


import (
	"gin-swagger-knife/constant"
	"gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)

func AddRouterOfChunkVendors15427cfaJs(router *gin.Engine) {
    
    utils.GetJs(router, CHUNK_VENDORS_15427CFA_JS_BASE64_OR_CONTENT, CHUNK_VENDORS_15427CFA_JS_RELATIVE_PATH)
    
}

const (
	CHUNK_VENDORS_15427CFA_JS_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/js/chunk-vendors.15427cfa.js}"
)





