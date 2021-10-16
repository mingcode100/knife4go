package knife

import (
	v2 "gin-swagger-knife/v2"
	"github.com/gin-gonic/gin"
	
	css "gin-swagger-knife/v2/knife/webjars/css"
	
	fonts "gin-swagger-knife/v2/knife/webjars/fonts"
	
	img "gin-swagger-knife/v2/knife/webjars/img"
	
	js "gin-swagger-knife/v2/knife/webjars/js"
	
	knife "gin-swagger-knife/v2/knife"
	
)

//@param content string swag int 命令生成的swagger.json文件里的内容
func InitSwaggerKnife(router *gin.Engine, swaggerJson string) {
	v2.AddApiDocRouter(router, swaggerJson)
	v2.AddSwaggerResourcesRouter(router)
    
        knife.AddRouterOfREADMEMd(router)
    
        knife.AddRouterOfDocHtml(router)
    
        knife.AddRouterOfFaviconIco(router)
    
        knife.AddRouterOfRobotsTxt(router)
    
        css.AddRouterOfApp5a1f33deCss(router)
    
        css.AddRouterOfApp5a1f33deCssGz(router)
    
        css.AddRouterOfChunkE8f6d0a483c9229fCss(router)
    
        css.AddRouterOfChunkVendorsF24a310aCss(router)
    
        css.AddRouterOfChunkVendorsF24a310aCssGz(router)
    
        fonts.AddRouterOfFontawesomeWebfont706450d7Ttf(router)
    
        fonts.AddRouterOfFontawesomeWebfont97493d3fWoff2(router)
    
        fonts.AddRouterOfFontawesomeWebfontD9ee23d5Woff(router)
    
        fonts.AddRouterOfFontawesomeWebfontF7c2b4b7Eot(router)
    
        fonts.AddRouterOfIconfont4ca3d0c0Ttf(router)
    
        fonts.AddRouterOfIconfontE2d2b98eEot(router)
    
        img.AddRouterOfEditormdLogo84b6c2a9Svg(router)
    
        img.AddRouterOfFontawesomeWebfont139e74e2Svg(router)
    
        img.AddRouterOfIconfontDd63dc33Svg(router)
    
        img.AddRouterOfLoadingC929501eGif(router)
    
        img.AddRouterOfLoading2x695405a9Gif(router)
    
        img.AddRouterOfLoading3x65eacf61Gif(router)
    
        js.AddRouterOfApp4a0199ecJs(router)
    
        js.AddRouterOfApp4a0199ecJsGz(router)
    
        js.AddRouterOfChunk069eb4373e4f260aJs(router)
    
        js.AddRouterOfChunk069eb4373e4f260aJsGz(router)
    
        js.AddRouterOfChunk0fd67716Dedd1f18Js(router)
    
        js.AddRouterOfChunk0fd67716Dedd1f18JsGz(router)
    
        js.AddRouterOfChunk2d0af44eA1c2260bJs(router)
    
        js.AddRouterOfChunk2d0bd799Ad1cba1cJs(router)
    
        js.AddRouterOfChunk2d0d0b98B11399b3Js(router)
    
        js.AddRouterOfChunk2d0da53248668907Js(router)
    
        js.AddRouterOfChunk2d22269d2f6cb00eJs(router)
    
        js.AddRouterOfChunk3b888a6574410f20Js(router)
    
        js.AddRouterOfChunk3b888a6574410f20JsGz(router)
    
        js.AddRouterOfChunk3ec4aaa808f27524Js(router)
    
        js.AddRouterOfChunk3ec4aaa808f27524JsGz(router)
    
        js.AddRouterOfChunk589faee0Cb8ab011Js(router)
    
        js.AddRouterOfChunk589faee0Cb8ab011JsGz(router)
    
        js.AddRouterOfChunk735c675cEda1bcaeJs(router)
    
        js.AddRouterOfChunk735c675cEda1bcaeJsGz(router)
    
        js.AddRouterOfChunkAdb9e944Fa692c2aJs(router)
    
        js.AddRouterOfChunkAdb9e944Fa692c2aJsGz(router)
    
        js.AddRouterOfChunkE8f6d0a4Fa7c0223Js(router)
    
        js.AddRouterOfChunkE8f6d0a4Fa7c0223JsGz(router)
    
        js.AddRouterOfChunkVendors15427cfaJs(router)
    
        js.AddRouterOfChunkVendors15427cfaJsGz(router)
    
}