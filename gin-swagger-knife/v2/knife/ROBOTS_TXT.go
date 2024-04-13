package knife

import (
	"github.com/gin-gonic/gin"
	"github.com/mingcode100/knife4go/gin-swagger-knife/constant"
	"github.com/mingcode100/knife4go/gin-swagger-knife/utils"
)

const (
	ROBOTS_TXT_RELATIVE_PATH = constant.ROOT_PATH + "/robots.txt"
	// 文件内容的16进制表示
	ROBOTS_TXT_HEX_CONTENT = `557365722d6167656e743a202a0d0a446973616c6c6f773a0d0a`
)

func AddRouterOfRobotsTxt(router *gin.Engine) {

	utils.GetOther(router, ROBOTS_TXT_RELATIVE_PATH, ROBOTS_TXT_HEX_CONTENT)

}
