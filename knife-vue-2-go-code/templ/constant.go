package templ

const (
	KNIFE_GO_TEMPL = `package knife

import (
	v2 "gin-swagger-knife/v2"
	"github.com/gin-gonic/gin"
	{{}}


)

//@param content string swag int 命令生成的swagger.json文件里的内容
func InitSwaggerKnife(router *gin.Engine, swaggerJson string) {
	v2.AddApiDocRouter(router, swaggerJson)
	v2.AddSwaggerResourcesRouter(router)

}`

	WEBJARS_TEMPL = `package {{.PackageName}}

import (
	"gin-swagger-knife/knife/constant"
	"gin-swagger-knife/knife/utils"
	"github.com/gin-gonic/gin"
)

const (
	{{.FileName2}}_BASE64        = "{{.FileBase64}}"
	{{.FileName2}}_RELATIVE_PATH = constant.ROOT_PATH + "{{.FileRelavePath}}}"
)

func AddRouter(router *gin.Engine) {
	{{if eq .FileType  ''}}
	utils.GetJson(router, {{.FileName2}}_BASE64, {{.FileName2}}_RELATIVE_PATH)
	{{end}}
}
`
)
