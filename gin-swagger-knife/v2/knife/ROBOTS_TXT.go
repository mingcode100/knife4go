package knife


import (
	"gin-swagger-knife/constant"
	"gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)

func AddRouterOfRobotsTxt(router *gin.Engine) {
    
	utils.GetOther(router, ROBOTS_TXT_BASE64_OR_CONTENT, ROBOTS_TXT_RELATIVE_PATH)
	
}

const (
	ROBOTS_TXT_RELATIVE_PATH = constant.ROOT_PATH + "/robots.txt}"
	ROBOTS_TXT_BASE64_OR_CONTENT = `User-agent: *
Disallow:
`
)






