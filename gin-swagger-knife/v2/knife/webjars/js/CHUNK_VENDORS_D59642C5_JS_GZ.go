package js


import (
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	CHUNK_VENDORS_D59642C5_JS_GZ_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/js/chunk-vendors.d59642c5.js.gz"
	// 文件内容的16进制表示
)

func AddRouterOfChunkVendorsD59642c5JsGz(router *gin.Engine) {
    
	utils.GetOther(router, CHUNK_VENDORS_D59642C5_JS_GZ_RELATIVE_PATH, CHUNK_VENDORS_D59642C5_JS_GZ_HEX_CONTENT)
	
}






