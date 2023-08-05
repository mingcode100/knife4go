package knife

import (
	v2 "gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/v2"
	"github.com/gin-gonic/gin"

	css "gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/v2/knife/webjars/css"

	fonts "gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/v2/knife/webjars/fonts"

	icons "gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/v2/knife/img/icons"

	img "gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/v2/knife/webjars/img"

	js "gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/v2/knife/webjars/js"

	knife "gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/v2/knife"

	oauth "gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/v2/knife/webjars/oauth"
)

//@param content string swag int 命令生成的swagger.json文件里的内容
func InitSwaggerKnife(router *gin.Engine, swaggerJson string) {
	v2.AddApiDocRouter(router, swaggerJson)
	v2.AddSwaggerResourcesRouter(router)

	knife.AddRouterOfDocHtml(router)

	icons.AddRouterOfAndroidChrome192x192Png(router)

	icons.AddRouterOfAndroidChrome512x512Png(router)

	icons.AddRouterOfAppleTouchIcon120x120Png(router)

	icons.AddRouterOfAppleTouchIcon152x152Png(router)

	icons.AddRouterOfAppleTouchIcon180x180Png(router)

	icons.AddRouterOfAppleTouchIcon60x60Png(router)

	icons.AddRouterOfAppleTouchIcon76x76Png(router)

	icons.AddRouterOfAppleTouchIconPng(router)

	icons.AddRouterOfFavicon16x16Png(router)

	icons.AddRouterOfFavicon32x32Png(router)

	icons.AddRouterOfMsapplicationIcon144x144Png(router)

	icons.AddRouterOfMstile150x150Png(router)

	icons.AddRouterOfSafariPinnedTabSvg(router)

	css.AddRouterOfAppB27647f3Css(router)

	css.AddRouterOfAppB27647f3CssGz(router)

	css.AddRouterOfChunk296622eb20e6d994Css(router)

	css.AddRouterOfChunkD7d5f59cA9ffbfcbCss(router)

	css.AddRouterOfChunkVendorsF24a310aCss(router)

	css.AddRouterOfChunkVendorsF24a310aCssGz(router)

	fonts.AddRouterOfFontawesomeWebfont706450d7Ttf(router)

	fonts.AddRouterOfFontawesomeWebfont97493d3fWoff2(router)

	fonts.AddRouterOfFontawesomeWebfontD9ee23d5Woff(router)

	fonts.AddRouterOfFontawesomeWebfontF7c2b4b7Eot(router)

	fonts.AddRouterOfIconfont4ca3d0c0Ttf(router)

	fonts.AddRouterOfIconfontE2d2b98eEot(router)

	img.AddRouterOfEditormdLogo53ea80e2Svg(router)

	img.AddRouterOfFontawesomeWebfont29800836Svg(router)

	img.AddRouterOfIconfont1d48c203Svg(router)

	img.AddRouterOfLoadingC929501eGif(router)

	img.AddRouterOfLoading2x695405a9Gif(router)

	img.AddRouterOfLoading3x65eacf61Gif(router)

	js.AddRouterOfApp06135c03Js(router)

	js.AddRouterOfApp06135c03JsGz(router)

	js.AddRouterOfChunk069eb4372486111cJs(router)

	js.AddRouterOfChunk069eb4372486111cJsLICENSETxt(router)

	js.AddRouterOfChunk069eb4372486111cJsGz(router)

	js.AddRouterOfChunk0d102d5a144e3deeJs(router)

	js.AddRouterOfChunk0d102d5a144e3deeJsGz(router)

	js.AddRouterOfChunk0fd67716D57e2c41Js(router)

	js.AddRouterOfChunk0fd67716D57e2c41JsGz(router)

	js.AddRouterOfChunk296622ebE4e811f0Js(router)

	js.AddRouterOfChunk296622ebE4e811f0JsGz(router)

	js.AddRouterOfChunk2d0af44eEdce7802Js(router)

	js.AddRouterOfChunk2d0bd799D2d95c02Js(router)

	js.AddRouterOfChunk2d0da5320b13c746Js(router)

	js.AddRouterOfChunk3b888a658737ce4fJs(router)

	js.AddRouterOfChunk3b888a658737ce4fJsGz(router)

	js.AddRouterOfChunk3ec4aaa8A79d19f8Js(router)

	js.AddRouterOfChunk3ec4aaa8A79d19f8JsGz(router)

	js.AddRouterOfChunk589faee0F3951d02Js(router)

	js.AddRouterOfChunk589faee0F3951d02JsLICENSETxt(router)

	js.AddRouterOfChunk589faee0F3951d02JsGz(router)

	js.AddRouterOfChunk735c675cEcb8aa58Js(router)

	js.AddRouterOfChunk735c675cEcb8aa58JsGz(router)

	js.AddRouterOfChunkAdb9e94455c41d5bJs(router)

	js.AddRouterOfChunkAdb9e94455c41d5bJsLICENSETxt(router)

	js.AddRouterOfChunkAdb9e94455c41d5bJsGz(router)

	js.AddRouterOfChunkD7d5f59c42190458Js(router)

	js.AddRouterOfChunkD7d5f59c42190458JsLICENSETxt(router)

	js.AddRouterOfChunkD7d5f59c42190458JsGz(router)

	js.AddRouterOfChunkF876db6c4965fec9Js(router)

	js.AddRouterOfChunkF876db6c4965fec9JsGz(router)

	js.AddRouterOfChunkVendorsD59642c5Js(router)

	js.AddRouterOfChunkVendorsD59642c5JsLICENSETxt(router)

	js.AddRouterOfChunkVendorsD59642c5JsGz(router)

	oauth.AddRouterOfAxiosMinJs(router)

	oauth.AddRouterOfOauth2Html(router)
    
}