package js


import (
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	CHUNK_260D712A_390177FE_JS_LICENSE_TXT_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/js/chunk-260d712a.390177fe.js.LICENSE.txt"
	// 文件内容的16进制表示
	CHUNK_260D712A_390177FE_JS_LICENSE_TXT_HEX_CONTENT = `2f2a2120466f72206c6963656e736520696e666f726d6174696f6e20706c6561736520736565206d65726d6169642e65736d2e6d696e2e6d6a732e4c4943454e53452e747874202a2f0d0a`
)

func AddRouterOfChunk260d712a390177feJsLICENSETxt(router *gin.Engine) {
    
	utils.GetOther(router, CHUNK_260D712A_390177FE_JS_LICENSE_TXT_RELATIVE_PATH, CHUNK_260D712A_390177FE_JS_LICENSE_TXT_HEX_CONTENT)
	
}







