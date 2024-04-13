package icons

import (
	"github.com/gin-gonic/gin"
	"github.com/mingcode100/knife4go/gin-swagger-knife/constant"
	"github.com/mingcode100/knife4go/gin-swagger-knife/utils"
)

const (
	APPLE_TOUCH_ICON_60X60_PNG_RELATIVE_PATH = constant.ROOT_PATH + "/img/icons/apple-touch-icon-60x60.png"
	// 文件内容的16进制表示
	APPLE_TOUCH_ICON_60X60_PNG_HEX_CONTENT = `89504e470d0a1a0a0000000d494844520000003c0000003c08030000000d2229400000000467414d410000b18f0bfc6105000000206348524d00007a26000080840000fa00000080e8000075300000ea6000003a98000017709cba513c000001d4504c5445fffffffcfefdfcfdfdfcfcfdfdfdfd9cd9bc83cfac85d0ad84c8aa7f95997e89967e8a967d89958a94a0e7e9ebb8e3cf54bd8c4dba874886764354669aa3adfefefef7fcfa88d1af4dbb874caf8445676c435366546374d6daddd6efe35cc09149947b435667838e9af8f9fa9dd9bd4db5854672704a5b6cc2c7cde6f5ee69c59a4aa07e445a686e7c89f0f1f268c599b3e2cc4fbb884db987477e74445567abb3bbf2faf679cba54ba98244616a5e6c7ce2e5e8c8eada55bd8d488b78949ea8fbfbfbfafdfc8dd3b24cb184456a6d536373a4acb5dbf1e75fc1934a977c43576748596afeffffa3dbc14db686467571e9f7f06cc69c4ba27f445c69b8e4cf50bb89478175f5fbf87ecda84cac8244636bcdecde57be8e498e7943556693d5b64cb385456d6edff3e962c2954a9a7d435867a8ddc54eba874db786477872edf8f371c89f4ba580445e69bee6d352bc8a488576f7fcf94cae8345666cd2eee15abf9049917afdfefe98d7ba4cb48546706f435165e3f4ec66c4984a9d7e445f6aafe0c94ebb884a9c7df0f9f575caa2c4e8d7f9fdfb89d1afd7f0e45dc0929fdabee7f6ef6ac59ab4e2cdf2faf77acca6c9eadb56be8d8fd3b360c194a4dcc2eaf7f16ec79dbae4d051bc8a80cea9ceecde63c39636aa09e100000001624b47440088051d480000000774494d4507e10113070b2cba6310eb000002994944415448c7edd5e55f5a611407f0a353a6ce42c4ba763b6376c730863abb30b1bb3b66ceeecefdb353383f061384cbcbcdf38a7bcef3fdfceee5792e10bdd77f543628db0f5c767a533b346d75ebf4b1bde4a3b61c1c9d3ebd94b38bab6ee8eae2ace939393af02a89bd3e7673977279c8b83ce53c937ba2e58145ee6efad8cb9bdb3ebe7e8276a57f00cf02fcb50dc1cfd78757797b193c746010747008e78486692661a17c1d120c1b1468f88dc9c3714b11911c1d15ad994447717064049684cb0d31c57c46746c1c47c5273cf713e2f92a2e16c15f625eed566212cf9253383a35edb99d96cac129c93c4f4a7cbdd5e91988cecc629d9d4339d96cb332119c916ee4a0e4e641e7636f0abe16e0633e6c5eaeb153a628042e2ae6e8926f251c5c5c045ca8307a4695a5d065e5fcd4dff989cbcb604b95c60f784525f6a2aa5a90e997505d8551658589d7a3a616d175f5fa5aa8af43706d8dc997aba111014d06b809edc606939654cd886e69fda385d6160437ab4c636a6b87ee50ebb0ba03b6bded0d4b369dc05ddd8816babb803b6ddec2d4d30bddd7cfb8bf0fb6b787deae81415e3934ac8d168687b8333860c6d2c828a2c7c65fb4303e86e0d11173982626b12f531a3c85cbc909b396a625889e991564c2ec0c8225d3e631cdcd432fa865ea05d8f9390b2c2d2ee14697577eac2ce36269d1124cab6b885edf5847f0daaa459668730bdbf513dbb4b569a1a5ed1ddc2b72a53bdb9662dad5fd8170b9ef5a6c696fff2fbcbf6739a68343037b7820c292fcc8001fc9c5603a3ed1b327c7a22cd1a9ee1749da782ad2d2d9b90e9f9f89c57471c9f6f242b425c515e32b85784cca6b8dbd565a61896e34f8c62a4bb777cff6eed63a4cf70fd2877b2b2da91ea58f2a6b313dfd7ab2dad2de9398b7e9bdfed1fa0d9a2598641f2212790000002574455874646174653a63726561746500323031372d30312d31395430373a31313a34342b30313a30300919db600000002574455874646174653a6d6f6469667900323031372d30312d31395430373a31313a34342b30313a3030784463dc000000577a5458745261772070726f66696c65207479706520697074630000789ce3f20c0871562828ca4fcbcc49e5520003230b2e630b1323134b9314031320448034c3640323b35420cbd8d4c8c4ccc41cc407cb8048a04a2e00ea171174f24235950000000049454e44ae426082`
)

func AddRouterOfAppleTouchIcon60x60Png(router *gin.Engine) {

	utils.GetOther(router, APPLE_TOUCH_ICON_60X60_PNG_RELATIVE_PATH, APPLE_TOUCH_ICON_60X60_PNG_HEX_CONTENT)

}
