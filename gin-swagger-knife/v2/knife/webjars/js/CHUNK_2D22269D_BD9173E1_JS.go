package js

import (
	"github.com/gin-gonic/gin"
	"github.com/mingcode100/knife4go/gin-swagger-knife/constant"
	"github.com/mingcode100/knife4go/gin-swagger-knife/utils"
)

const (
	CHUNK_2D22269D_BD9173E1_JS_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/js/chunk-2d22269d.bd9173e1.js"
	// 文件内容的16进制表示
	CHUNK_2D22269D_BD9173E1_JS_HEX_CONTENT = `2877696e646f772e7765627061636b4a736f6e703d77696e646f772e7765627061636b4a736f6e707c7c5b5d292e70757368285b5b226368756e6b2d3264323232363964225d2c7b636630343a66756e6374696f6e28742c652c69297b2275736520737472696374223b692e722865293b766172206e3d7b6e616d653a22456469746f7253686f77222c636f6d706f6e656e74733a7b656469746f723a6928223763396522297d2c70726f70733a7b76616c75653a7b747970653a537472696e672c72657175697265643a21302c64656661756c743a22227d7d2c646174613a66756e6374696f6e28297b72657475726e7b6c616e673a226a617661736372697074222c656469746f723a6e756c6c2c656469746f724865696768743a3230307d7d2c6d6574686f64733a7b7265736574456469746f724865696768743a66756e6374696f6e28297b76617220743d746869733b73657454696d656f7574282866756e6374696f6e28297b76617220653d742e656469746f722e73657373696f6e2e6765744c656e67746828293b313d3d65262628653d3130293b76617220693d31362a653b742e656469746f724865696768743d697d292c333030297d2c6368616e67653a66756e6374696f6e2874297b746869732e24656d697428226368616e6765222c74297d2c656469746f72496e69743a66756e6374696f6e2874297b76617220653d746869733b746869732e656469746f723d742c6928223230393922292c6928226262333622292c6928223164323922292c746869732e7265736574456469746f7248656967687428292c746869732e656469746f722e72656e64657265722e6f6e2822616674657252656e646572222c2866756e6374696f6e28297b652e24656d6974282273686f774465736372697074696f6e222c2231323322297d29297d7d7d2c723d6928223238373722292c6f3d4f626a65637428722e6129286e2c2866756e6374696f6e28297b76617220743d746869732c653d742e24637265617465456c656d656e742c693d742e5f73656c662e5f637c7c653b72657475726e20692822646976222c5b692822656469746f72222c7b61747472733a7b76616c75653a742e76616c75652c6c616e673a742e6c616e672c7468656d653a2265636c69707365222c77696474683a2231303025222c6865696768743a742e656469746f724865696768747d2c6f6e3a7b696e69743a742e656469746f72496e69742c696e7075743a742e6368616e67657d7d295d2c31297d292c5b5d2c21312c6e756c6c2c6e756c6c2c6e756c6c293b652e64656661756c743d6f2e6578706f7274737d7d5d293b`
)

func AddRouterOfChunk2d22269dBd9173e1Js(router *gin.Engine) {

	utils.GetJs(router, CHUNK_2D22269D_BD9173E1_JS_RELATIVE_PATH, CHUNK_2D22269D_BD9173E1_JS_HEX_CONTENT)

}
