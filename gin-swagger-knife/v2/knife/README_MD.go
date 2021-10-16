package knife


import (
	"gin-swagger-knife/constant"
	"gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)

func AddRouterOfREADMEMd(router *gin.Engine) {
    
	utils.GetOther(router, README_MD_BASE64_OR_CONTENT, README_MD_RELATIVE_PATH)
	
}

const (
	README_MD_RELATIVE_PATH = constant.ROOT_PATH + "/README.md}"
	README_MD_BASE64_OR_CONTENT = `# knife-vue-dist

#### 介绍
knife-vue 打包的dist目录文件
 

#### knife-vue地址
https://gitee.com/xiaoym/knife4j/knife-vue`
)






