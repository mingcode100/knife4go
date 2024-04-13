package js

import (
	"github.com/gin-gonic/gin"
	"github.com/mingcode100/knife4go/gin-openapi3-knife/constant"
	"github.com/mingcode100/knife4go/gin-openapi3-knife/utils"
)

const (
	CHUNK_D7D5F59C_E61130F3_JS_LICENSE_TXT_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/js/chunk-d7d5f59c.e61130f3.js.LICENSE.txt"
	// 文件内容的16进制表示
	CHUNK_D7D5F59C_E61130F3_JS_LICENSE_TXT_HEX_CONTENT = `2f2a210d0a202a20636c6970626f6172642e6a732076322e302e360d0a202a2068747470733a2f2f636c6970626f6172646a732e636f6d2f0d0a202a200d0a202a204c6963656e736564204d495420c2a9205a656e6f20526f6368610d0a202a2f0d0a`
)

func AddRouterOfChunkD7d5f59cE61130f3JsLICENSETxt(router *gin.Engine) {

	utils.GetOther(router, CHUNK_D7D5F59C_E61130F3_JS_LICENSE_TXT_RELATIVE_PATH, CHUNK_D7D5F59C_E61130F3_JS_LICENSE_TXT_HEX_CONTENT)

}
