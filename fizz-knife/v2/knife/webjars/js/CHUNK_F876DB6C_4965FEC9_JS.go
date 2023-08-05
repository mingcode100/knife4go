package js


import (
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	CHUNK_F876DB6C_4965FEC9_JS_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/js/chunk-f876db6c.4965fec9.js"
	// 文件内容的16进制表示
)

func AddRouterOfChunkF876db6c4965fec9Js(router *gin.Engine) {
    
    utils.GetJs(router, CHUNK_F876DB6C_4965FEC9_JS_RELATIVE_PATH, CHUNK_F876DB6C_4965FEC9_JS_HEX_CONTENT)
    
}






