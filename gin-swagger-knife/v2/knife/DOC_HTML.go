package knife


import (
	"gin-swagger-knife/constant"
	"gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	DOC_HTML_RELATIVE_PATH = constant.ROOT_PATH + "/doc.html"
	DOC_HTML_HEX_CONTENT = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width,initial-scale=1">
    <link rel="icon" href="favicon.ico">
    <title></title>
    <link href="webjars/css/chunk-e8f6d0a4.83c9229f.css" rel="prefetch">
    <link href="webjars/js/chunk-069eb437.3e4f260a.js" rel="prefetch">
    <link href="webjars/js/chunk-0fd67716.dedd1f18.js" rel="prefetch">
    <link href="webjars/js/chunk-2d0af44e.a1c2260b.js" rel="prefetch">
    <link href="webjars/js/chunk-2d0bd799.ad1cba1c.js" rel="prefetch">
    <link href="webjars/js/chunk-2d0d0b98.b11399b3.js" rel="prefetch">
    <link href="webjars/js/chunk-2d0da532.48668907.js" rel="prefetch">
    <link href="webjars/js/chunk-2d22269d.2f6cb00e.js" rel="prefetch">
    <link href="webjars/js/chunk-3b888a65.74410f20.js" rel="prefetch">
    <link href="webjars/js/chunk-3ec4aaa8.08f27524.js" rel="prefetch">
    <link href="webjars/js/chunk-589faee0.cb8ab011.js" rel="prefetch">
    <link href="webjars/js/chunk-735c675c.eda1bcae.js" rel="prefetch">
    <link href="webjars/js/chunk-adb9e944.fa692c2a.js" rel="prefetch">
    <link href="webjars/js/chunk-e8f6d0a4.fa7c0223.js" rel="prefetch">
    <link href="webjars/css/app.5a1f33de.css" rel="preload" as="style">
    <link href="webjars/css/chunk-vendors.f24a310a.css" rel="preload" as="style">
    <link href="webjars/js/app.4a0199ec.js" rel="preload" as="script">
    <link href="webjars/js/chunk-vendors.15427cfa.js" rel="preload" as="script">
    <link href="webjars/css/chunk-vendors.f24a310a.css" rel="stylesheet">
    <link href="webjars/css/app.5a1f33de.css" rel="stylesheet">
</head>
<body>
<noscript><strong>We're sorry but knife4j-vue doesn't work properly without JavaScript enabled. Please enable it to
    continue.</strong></noscript>
<div id="app"></div>
<script src="webjars/js/chunk-vendors.15427cfa.js"></script>
<script src="webjars/js/app.4a0199ec.js"></script>
</body>
</html>
`
)

func AddRouterOfDocHtml(router *gin.Engine) {
    
    utils.GetHtml(router, DOC_HTML_HEX_CONTENT, DOC_HTML_RELATIVE_PATH)
    
}







