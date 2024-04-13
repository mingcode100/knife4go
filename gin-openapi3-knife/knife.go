package knife

import (
	"github.com/gin-gonic/gin"
	v3 "github.com/mingcode100/knife4go/gin-openapi3-knife/v3"

	css "github.com/mingcode100/knife4go/gin-openapi3-knife/v3/knife/webjars/css"

	fonts "github.com/mingcode100/knife4go/gin-openapi3-knife/v3/knife/webjars/fonts"

	icons "github.com/mingcode100/knife4go/gin-openapi3-knife/v3/knife/img/icons"

	img "github.com/mingcode100/knife4go/gin-openapi3-knife/v3/knife/webjars/img"

	js "github.com/mingcode100/knife4go/gin-openapi3-knife/v3/knife/webjars/js"

	knife "github.com/mingcode100/knife4go/gin-openapi3-knife/v3/knife"

	oauth "github.com/mingcode100/knife4go/gin-openapi3-knife/v3/knife/webjars/oauth"
)

// @param content string swag int 命令生成的swagger.json文件里的内容
func InitSwaggerKnife(router *gin.Engine, swaggerJson string) {
	v3.AddApiDocRouter(router, swaggerJson)
	v3.AddSwaggerResourcesRouter(router)

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

	css.AddRouterOfAppAc23e017Css(router)

	css.AddRouterOfAppAc23e017CssGz(router)

	css.AddRouterOfChunk75464e7e8fb93ba5Css(router)

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

	js.AddRouterOfApp2fab4ac5Js(router)

	js.AddRouterOfApp2fab4ac5JsGz(router)

	js.AddRouterOfChunk069eb4372cfebf27Js(router)

	js.AddRouterOfChunk069eb4372cfebf27JsLICENSETxt(router)

	js.AddRouterOfChunk069eb4372cfebf27JsGz(router)

	js.AddRouterOfChunk0d102d5aB2bddffcJs(router)

	js.AddRouterOfChunk0d102d5aB2bddffcJsGz(router)

	js.AddRouterOfChunk0fd67716D57e2c41Js(router)

	js.AddRouterOfChunk0fd67716D57e2c41JsGz(router)

	js.AddRouterOfChunk260d712a390177feJs(router)

	js.AddRouterOfChunk260d712a390177feJsLICENSETxt(router)

	js.AddRouterOfChunk260d712a390177feJsGz(router)

	js.AddRouterOfChunk2d0af44e392afcd6Js(router)

	js.AddRouterOfChunk2d0bd799Eb48b7f1Js(router)

	js.AddRouterOfChunk2d0da532591ad7fcJs(router)

	js.AddRouterOfChunk3b888a658737ce4fJs(router)

	js.AddRouterOfChunk3b888a658737ce4fJsGz(router)

	js.AddRouterOfChunk3ec4aaa8A79d19f8Js(router)

	js.AddRouterOfChunk3ec4aaa8A79d19f8JsGz(router)

	js.AddRouterOfChunk589faee05bfd1708Js(router)

	js.AddRouterOfChunk589faee05bfd1708JsLICENSETxt(router)

	js.AddRouterOfChunk589faee05bfd1708JsGz(router)

	js.AddRouterOfChunk735c675c5b409314Js(router)

	js.AddRouterOfChunk735c675c5b409314JsGz(router)

	js.AddRouterOfChunk75464e7eB130271bJs(router)

	js.AddRouterOfChunk75464e7eB130271bJsGz(router)

	js.AddRouterOfChunkAdb9e9442c7f24feJs(router)

	js.AddRouterOfChunkAdb9e9442c7f24feJsLICENSETxt(router)

	js.AddRouterOfChunkAdb9e9442c7f24feJsGz(router)

	js.AddRouterOfChunkD7d5f59cE61130f3Js(router)

	js.AddRouterOfChunkD7d5f59cE61130f3JsLICENSETxt(router)

	js.AddRouterOfChunkD7d5f59cE61130f3JsGz(router)

	js.AddRouterOfChunkVendorsD51cf6f8Js(router)

	js.AddRouterOfChunkVendorsD51cf6f8JsLICENSETxt(router)

	js.AddRouterOfChunkVendorsD51cf6f8JsGz(router)

	oauth.AddRouterOfAxiosMinJs(router)

	oauth.AddRouterOfOauth2Html(router)

}
