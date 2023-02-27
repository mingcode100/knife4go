package css


import (
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	CHUNK_296622EB_20E6D994_CSS_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/css/chunk-296622eb.20e6d994.css"
	// 文件内容的16进制表示
	CHUNK_296622EB_20E6D994_CSS_HEX_CONTENT = `.api-title[data-v-5df78020]{margin-top:10px;margin-bottom:5px;font-size:16px;font-weight:600;height:30px;line-height:30px;border-left:4px solid #00ab6d;text-indent:8px}`
)

func AddRouterOfChunk296622eb20e6d994Css(router *gin.Engine) {
    
    utils.GetCss(router, CHUNK_296622EB_20E6D994_CSS_RELATIVE_PATH, CHUNK_296622EB_20E6D994_CSS_HEX_CONTENT)
	
}







