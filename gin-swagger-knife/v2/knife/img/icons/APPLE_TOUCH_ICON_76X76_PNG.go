package icons

import (
	"github.com/gin-gonic/gin"
	"github.com/mingcode100/knife4go/gin-swagger-knife/constant"
	"github.com/mingcode100/knife4go/gin-swagger-knife/utils"
)

const (
	APPLE_TOUCH_ICON_76X76_PNG_RELATIVE_PATH = constant.ROOT_PATH + "/img/icons/apple-touch-icon-76x76.png"
	// 文件内容的16进制表示
	APPLE_TOUCH_ICON_76X76_PNG_HEX_CONTENT = `89504e470d0a1a0a0000000d494844520000004c0000004c0803000000f049a1190000000467414d410000b18f0bfc6105000000206348524d00007a26000080840000fa00000080e8000075300000ea6000003a98000017709cba513c0000023d504c5445fffffff9fdfbf9fbfbf9f9faf9fafafefefe98d7b97bcca67ccca77ccda77abca175869074808e74818ebac1c7bde6d355bd8d4dba874db78647797243536643546662707fe6e8eb8dd3b24dbb874ba681445f699ba4addaf1e65ec193488676556474d6daddfeffffa1dbc04caf8345676c848f9bf8f9fae9f6f06cc69c49937a4356674a5b6cc2c8cdb8e3cf50bb894cb58546716f6f7c8af0f1f2f4fbf77dcda74a9f7e445a68445567acb3bbccebdd57be8e4db987477d735e6d7ce2e5e8fbfdfc92d5b64ba98144616a949ea8fdfdfd44606adef2e962c295488a77516171d1d5d9d1d6daa8ddc44cb18445696d7f8a97eef0f1ecf8f271c89f4a967b4357674b5b6d828e9a52bc8a4db686467470f6fbf983cfab4ba27f445c68d2ede05abf8f478174fcfefd4cab8244636be2f4ec65c397498d79435566aee0c84ebb884cb285456c6eaddfc8eff9f475c9a24a9a7d435867c3e8d753bc8b467771f8fcfa88d1af4ba580445d69d7efe45cc0914884769dd9bd4cad8345656be6f5ee69c59a49917ab4e2cc4fbb884cb485466f6ff2faf67acba54a9d7e435968c9eada4db886477b73fafdfc8ed3b34ba781445f6adbf1e75fc193488777a4dcc14caf84435265eaf7f16dc69d49947bb9e4d0498e79f5fbf87fcda97fcda8ceecde94d5b7e0f3ea63c296a9dec54eba87edf8f372c8a0bfe6d4f7fcf984cfacd3eee15abf90fdfefe99d7bae4f4ed66c498afe0c9f0f9f576caa3c4e8d754bd8c8ad2b0d8f0e55dc0929fdabfe7f6ef6ac59bb5e2cdf3faf7c9eadbc7e9d99516437300000001624b47440088051d480000000774494d4507e10113070b2dcd64207d0000037c4944415458c3edd6575b13511006e0410c282010587a13a4480d4de9551045050441a44b17547aef55142c60c182d244a44a6f8abf4d8c61671242d8943b3357e49c99f7f9963dd90d80bad4a52e000d52a7344fb3c513ebe2e186e6293a228e69699f39acb33aba7ae744a5a72fd6a58f1bba3a67d9096d2d71ccc090cf969131637258a666a4c9cc945d678c8d70c0d040e232cd71cfc2d20a356bbc509e355a56961638602e719960638b9b7c3b76cae4bc3ddb627f1e97ed48b7ad8de41d70b840a2393a6134671751878b3306737224c12e381cb99f175d89e6e6ce0e7a788a1a3c3dd835773762b95e94723abc04d8e0ed83d17cfd84db7ebe18ccc71b5b055ed28e9aff2512ed72003b1a1824dc0e0a6457022e936097fca51edce0106c090dc368e111079b11e1182c2c141b4382a57f0b22a348b4e82ba8c5c4426c0c5a57a349b0a8c863be5371d788761dcf41fc0db8198f1faf13eb5a1c1c57b708969088d1926e2761b0c40482dd3ad682e414a2dd496581b4bb69ecdfa9778895927c3c0606e9d8782f03a3e19d6432ee614bba810c0b32b348b4ec1cd4582b279b04cbca9485416e1e391ef7a560f7c9b1c8cb956941be1689565028a9318505249856be6c0c8a8a8956522a81959610abb8084eaa076588953f148fc63c2c47acecd189163cae20d12aabc4b0aa4a12ace2f1c91854d7e080512d8dc6d49267754d35070beaea49b48646d498c60612acbe8e0b064dcd38c26f21580b596f6ee26401af95446b6b3fd498f63612ac95c70d838e4ea2751d1e8fd22e62757670b400ba7b70acb7ef5f34a6af17177bba395bd0ff84441b782ac49e0e90604ffab963f08cbce00787fe4663860671c9f0991c163c7f41a2bd7cc59830af5e92602f9ecb83c1f0083906af0f92bd269f4786e5b2c0e10d89f6d6e99dd35b12ec8d837c188cbe27da87aa0fc47a3f2aa705f091bce07b3f916321f828b705639fc97f893c5ef95fc6e4c7609cbce049858c2b6041fe84546ce2a467b5f49a9c92624d4d2a64014c7f3d627d9d56d082996f47b09419453198ad91b0d26715b6e0fb9c0436f75d710ce617c4ac8579252c585c12c3961695c160f907b17e2c2b6501ace00bbe6c45490b56d7586c6d55590cd63744d6c6bad216686c8ab04d0de531d8da165adb5b2ab080b723c476b8bec265d7eede81b5b7ab120be0a7802ff8a9220b7e69f3b57fa90a83fd917d955950f79bdb2f3b75a9ebbfae3f3bbcf2ef518aca330000002574455874646174653a63726561746500323031372d30312d31395430373a31313a34352b30313a3030af6ed0d40000002574455874646174653a6d6f6469667900323031372d30312d31395430373a31313a34352b30313a3030de336868000000577a5458745261772070726f66696c65207479706520697074630000789ce3f20c0871562828ca4fcbcc49e5520003230b2e630b1323134b9314031320448034c3640323b35420cbd8d4c8c4ccc41cc407cb8048a04a2e00ea171174f24235950000000049454e44ae426082`
)

func AddRouterOfAppleTouchIcon76x76Png(router *gin.Engine) {

	utils.GetOther(router, APPLE_TOUCH_ICON_76X76_PNG_RELATIVE_PATH, APPLE_TOUCH_ICON_76X76_PNG_HEX_CONTENT)

}
