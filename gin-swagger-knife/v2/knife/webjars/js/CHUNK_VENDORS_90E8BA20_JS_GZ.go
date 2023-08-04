package js


import (
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	CHUNK_VENDORS_90E8BA20_JS_GZ_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/js/chunk-vendors.90e8ba20.js.gz"
	// 文件内容的16进制表示
)

func AddRouterOfChunkVendors90e8ba20JsGz(router *gin.Engine) {
    
	utils.GetOther(router, CHUNK_VENDORS_90E8BA20_JS_GZ_RELATIVE_PATH, CHUNK_VENDORS_90E8BA20_JS_GZ_HEX_CONTENT)
	
}






