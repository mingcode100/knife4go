package knife


import (
	"gin-swagger-knife/constant"
	"gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	README_MD_RELATIVE_PATH = constant.ROOT_PATH + "/README.md"
	README_MD_HEX_CONTENT = `23206b6e6966652d7675652d646973740d0a0d0a2323232320e4bb8be7bb8d0d0a6b6e6966652d76756520e68993e58c85e79a8464697374e79baee5bd95e69687e4bbb60d0a200d0a0d0a23232323206b6e6966652d767565e59cb0e59d800d0a68747470733a2f2f67697465652e636f6d2f7869616f796d2f6b6e696665346a2f6b6e6966652d767565`
)

func AddRouterOfREADMEMd(router *gin.Engine) {
    
	utils.GetOther(router, README_MD_HEX_CONTENT, README_MD_RELATIVE_PATH)
	
}







