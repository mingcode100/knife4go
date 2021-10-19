package knife


import (
	"gin-swagger-knife/constant"
	"gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	ROBOTS_TXT_RELATIVE_PATH = constant.ROOT_PATH + "/robots.txt"
	ROBOTS_TXT_HEX_CONTENT = `557365722d6167656e743a202a0d0a446973616c6c6f773a0d0a`
)

func AddRouterOfRobotsTxt(router *gin.Engine) {
    
	utils.GetOther(router, ROBOTS_TXT_HEX_CONTENT, ROBOTS_TXT_RELATIVE_PATH)
	
}







