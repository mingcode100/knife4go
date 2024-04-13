package icons

import (
	"github.com/gin-gonic/gin"
	"github.com/mingcode100/knife4go/fizz-knife/constant"
	"github.com/mingcode100/knife4go/fizz-knife/utils"
)

const (
	APPLE_TOUCH_ICON_PNG_RELATIVE_PATH = constant.ROOT_PATH + "/img/icons/apple-touch-icon.png"
	// 文件内容的16进制表示
	APPLE_TOUCH_ICON_PNG_HEX_CONTENT = `89504e470d0a1a0a0000000d49484452000000b4000000b40802000000b2af91650000000467414d410000b18f0bfc6105000000206348524d00007a26000080840000fa00000080e8000075300000ea6000003a98000017709cba513c00000006624b474400ff00ff00ffa0bda7930000000774494d4507e10113070b31d9657c32000010e74944415478daed9ddb7f13d7b5c7d79ed9238d6ea38b65c996ac11f882aff205c9a6714b42024e81f8f2f9f4adef7d3d0fe76f39afe7fd9c8763d27c1a9352ec00ad4b818634044a52282481a6210162c098e0eb9c87b11d021edb23ada5d148fbfba40f68b6f6defecdccbeacdfdacc300c1008b64272ba0282ca45884360891087c012210e8125421c024b843804960871082c11e2105822c421b044884360891087c012210e8125421c024b843804960871082c11e2105822c421b044884360891087c012210e8125421c024b843804960871082c11e21058c2ed5e7077e101be0d8a81b1b8662caca015680057e486fa3ace6d37d0162b2b2bf7ee3f5c595e0586d71901cebc12607732639009c46d5d62bbeffeef8b0b1f3db885d71566c561edf1cad2f477c6fc0a562f7b3cca7ffce6d7c3853edc9abec4a54ffefe5ffffd3f4b4bcb38c519c042dc7324218539ae380c8042bce53f7b466d5d655b1cc79a06aecedd79f0fc0966dd01c007cb69b6fcd153c42227a7a67bda5bb55000b9aa1b3c995f989c9abeff700eb14ca523bae45b062cb56d1057b5634d0376afb23de6e888a44752bdb855070060c03b4252dc8378c75cbdfecf0f672fe15775830f672f5dbdfe4fb4e20c90e21ede11427c436d3292eaed88a4ed5e55cc807424dddb1c4a2257df001690792e0c325adfacadadbd77eaecd7f7ee2357150000bebe77ffbd5367d7d6d6d04a9419cf855940461f6d34879223e962eee762c41157b5d14c9e4b3272230ce07b0372c687d83b77bfbef7fee973e899040cc378fff4b9bb5fdf436cbb9cf1f1bd01746570491ecde4e3aa56c4b5454e658793ed7db12c723b0040613c17662ae658fdf4b90bd76fdcc6ade6f51bb74f9fbb80569c014c95782e0c0afe1ba52f961d4eb617776d91e2f073ef843e14e02a72530c901b55b9358858e4a3c7f39353d38b8b4b58052e2e2e4d4e4d3f7a3c8f5849b9352837aae88f8d005727f4213ff7167779f18b60b9987eb0a103b935002001efd618ea5ceee2e5ab173efe14abb40b1f7f7af1f255b4ca19c0c29c776b14eb91071b3a7231bde8cb8baf91cca47732f9842f8cdc2003a4a8c2bb34c441fbe2d2f2e4fb33738f11a6df738f9f4cbe3fb38838d564c0bb3429aaa03f3612bef03b99bccc8aff139724d76cb0fe68ba9f60e6057c5f504a7a11fbebb39b5fa08c124e9fbbf0d9cd2fd0aa658094f4f27d98af5113067034dd9f0dd6975248a9cfb2c3a95c5bb811b96506309face4c2c0d1846718c6effe70eeab7f7d534a215ffdeb9bdffd0175eec399920b331ffef4b52ddc7838952bb19052c511f506c7f4418f84bd8561809cf5cb593f62af7df3ed83f74e9d5d5d2d7265627575edbd5367bff9f64125b7d1c423f1317d30ea2df58184300a3a50df968f3723b70f48eeaa33b37ffdf4b31bc55dfbe96737ceccfe15ad2a044fc74df2f1e603f56da59783200e5556c6f5414df1213791e07d3cff7461726ae6d90fcfed5ef8ec87e7935333f34f17102b833eae32d114dfb83ea8ca4ae945e1cc9f3a234d6f347623b7124846f297af5cfff3a54fec5ef5e74b9f5cbe721dad120433b24dde68ecee8c34a11485230e89b1e399fd697f0cb9a1046b00cbcb2b274ece3cf8fed1ee2f79f0fda31327679697f1c24d08d6724cd2fed8f1cc7e89e1880eadd7d3fed8b1cc0056b55e047df5f0e6ed3b1f7c38bbfbef7ff0e1eccddb77d07e9e6015d84462ec586600f116c55c963bd4d883f540fb119a7d8793d3b3b7bebcbb9b6fdefaf2eec9691b4ada1982fd2393ce48d3a1c61ec40231c5a129be71bd803214fa09043b96f71fcebdfbc19995951dde142b2b2bef7e7006339c8760e7d944959571bd803b2d405ed0cfc75b863026512f4310ebf0c7bf5cfef8eae7db7fe7e3ab9ffff12f97d17e723d6645438c59d964a8be2d1f6fc12d13591c1e898feb85a8073b326f334a0a8f673f3c9f9c9a7ebaf0ccea0b4f179e4d4e4d1731efdd06de1192e2f8d3d7a82730ae17d09722f1b702dbc2a9d2176eb7c08c23acc38c23bc72edc6d9f31f59fdefd9f31f5db956e48ad9161820d55145011e4ee5dac229f462f1c5c1007ed9d4af07ed45c1ef8c012cc4790fe633796575f5b7bf3f73effec357ffebdefd87bffdfd9995d555b4facb8cf7682c843f7dd583f15f36916c7f92989a92be48899bc55b6380dc12905398d3da2feffc7b6afa4f2feda519863135fda72feffc1bb3e629556ec18f0234032792be0872b90040e7783b98ecec89161f666205f348bc37ccbc98d53e75e6fc3f6e7df5e2bffce3d657a7ce9cc7acb657e2bd61e6c1efed9ea87e30d9895eac099538828a3a911d2c3a40cd1203e4b44f6ec6bc05bf9f7b7ce2e4ccd2f27afcced2f2f2899333dfcf3dc6ac7373404ee34f5ffddc3b911d0c2ad8c19a1b107a65fb637b86134586b66e5f65f497f7f94b572efded9af9f9d2dfae9dbf7405ade8cda112414f0f27dafb637bf0cbdd80501c5c9247f52283e2b7c31cf677620efb9f2f2e9e989a7932bff0647ee1c4d4ccf3c545b4a219f04ee44996495cd546750283c80bd0baec8bb6d3ec086f0f49f5980b06d73ebf35337b7166f6e2b5cf6fa1156a8054efe5ed98cb339b9058cb7e0aad091d0046527d17bfbb797bfe5bcc4237961a97ce3d80551c81acadadfdefbba7cc0f68f5e48ce7342a135b8ad6230e65c8cf115743e37a81c41eb7077993e2e1dca387738f106b28677c7c0f89896d4c2fc4559207d28b942379cb6b44e326852934db9b0818c05449a131b1518df45fa11ce2f0710fc98ccb00a94195dbf003235090db825203be89cd5c23f0714f199a50a6b44f546b351228dd1a0be33b824ac20016561422131bcdeae29694491c329346f57c03fa2aaf012ca228dd24bb59c5c340e90eb108be641b7c91519d605fc282f2258ccb04e2479bfa19451c615b504ae23fc08bc4002949f2b26300479bfaede6f52a85b266137c2b95dba7d1d8e37a358aa15f31284ce9d5284c6cfbc2a9b7286221ac29ab38229ec038913d4e27b18e155393ac5fd6494c6ce3fa60043d8a6a5bca9d8774a8beb5801dcd06b0618ff3e3dfaf363080f9a94c6c8578cb507d6b991b546e717865653c3b18f6f891cb35404a90d8d56dc1f705a5047e1460d8e31fcf0e7ad123b777c2810cc69de1f49ba811f4eb9025bad8159426b6371b7b3ac3b67301968e03e2608c1d6d1a480708ec711ad5e6f8ce9881041a81892d103bda344031cbdb459b9c20e58f1e6f4273edfd8861dae3f0c36a76fedd469fdc1a44ff5d89b1e34dfb53fe6879dbb3f1eb8efc2a001c6aecee8a64d08b655e89f76a401090b71d1e89f76ab8c18b265d91cc210a87faee704c1ca1757b1cf61e810172938fef2de3b4d600bed72f375198d83ce37a21849edb62d73879a4c6fe78f3cf1264f6b86059a6b506b02072e2e54d7e9668db4f911567d738290e8fc4c7f442e9d9895e86d23ef43204562b93a83738466062b385c387f1b46a8d475c628fdb024a151e49e55ad1b71a6ce2b0384c7b5c891911b7c00016e4bc87e469ff2332e33d6116c49fbe6683f54426365b387f8c57420d936c431bc09bfd729a6cb7d60039adf266fc91af19de9050b1b3ffdac7797100c02f921dbd1469f609ec719b9826368a39736f2cfb8b2441de70fb548438029ccc1e97f2513854375cbb542636fc13078aa222c401007db13d3fa7b85d24e03d61e455edf575fa3045e7fd3cd9d1476962b345a5884366d268265f4f618f8b29bc13753f8c01efd4a418fe0e5fbdaa8d52642728964aa90700ec0d25de264ab3df8eb7936ec606b4934401be9deedf1b4a1074409154903800e0483ad7a23520178a1b83431655d4a2351c4997350a70472a4b1c75ded0985e5008ec7172d68f608f3340cef828e21115491ed30b755e72139b2d2a4b1c00f05a62df401dc18682c294de30f395608f3380f924a597c4c43650b7f7b5c43efc569746c58943953d137a2144618f2bd93140e4810829ea843e88bf415d3215270e00e88eeaaf3710043148a07469457a8d4cf754174998d9eb0d5ddde532b1d9a212c561a6d927b3c715f507367d973426b6e3993c45cef8d2a9447100402650772c43123829b7da7f3598af248254f68cb16399814ca00ebd64142a541c00f05663ae1d3df16a7183cad207b316b487536f3556d6f4f5452a571c618f7f4227306b98d3d1dde7543140c6ce1263e29595099dc0c28347e58a03000af116327b9cb6ab85acf505348dc8c446d23a3c2a5a1c54f7969d3c6ee899e94ca89e8ba854b43880eeadcc807785a4d8b671840648310fef2289022419516153e9e2a01acfef26772c41365c13bab9182e952e0ed85809204bb36f31d2240b14a25ac521c005e20080d71bbabaa244f6b8adf3d55364e037a15aff25c01de2a0da7d30404eab5ba4d95f4f658fbf8da2ca9e097d107fe78806778803e8f62dcd33725eb417acdb1a484e627b2db16fa06e2f694721e21a7150453cbc9a669f2c953d55b40a19ae1107d0c54a99f6b8b8078c8d9306db694c6c14716e94b8491c545196669a7dd31eb76e62c38f02a48b90a5c34de200baf86c338d425a95d32a45fa06aad87a625c260ea073767824251f55f2510a131b952b8718f78983ca13668094f052e402a4f2f3d1e33e714025b94977039513981e578aa3727ce83b429543a02cb8b2d25031192cb6872afb48b970ab38a03272df6c0f55dea272e162715442d6ac6da0ca7856465c2c0ea8807c7bdb40952bb18cb85b1c8e67eab4822acb6a7971b738c0e91cbf5610e5672e33ae1787b3d9c1b7842ab37bd971bd38c0d173055e85ea4c0827a8067180732792bc0ad569324e5025e270ea2ca397a03a87ca21aa441ce0d029682f4175829d43548f381c393ff145a8cebe748eea11073871f2ea8b909c9aeb2855258ef29fd9bc09dd79db0e5255e280b29ff66e628610b8c2c4668b6a1307001c4c76f69437c5564f543f98ec74baddf854a138828a3a911df4f132ed6bf8b867223b18748989cd1655280e00e88fed194eb497e7b78613edfd1593ca1e97ea140797e431bd1057c91302c7d5d0985ee0ee31b1d9a23ac50100cda1e448aa8ffa5746527dcda1a4d36da5a26ac5010023e95ed2bf5c73283992ee75ba958454b338e2aa36aae7899ef95c9247f57cdc6d26365b54b3388072b458ce31af5354b9384c7b1cfa3cd39c2dbbd1c4668b2a1707d0ac50957f9dcd11aa5f1c3293dec9e493786bdb495fe49d4a3a898d8eea6f2100e8c138963dce34b1e94107f6f6ca4f4d8803000ea7726d184961dbc2a9c36e36b1d9a256c411f504c64bb6c779243eae17a2cec51395995a1107000cd5b7e54b8be1cbc75b86eadd6d62b3450d8943959571bda0156b8fd3d64d6c159dca1e971a12070074469a0e15eb1b38d4d8d3196972ba0565a5b6c42131762c3390f6db761ca5fdb16399812a30b1d9a2b6c40100697fec78c69e57d14c655f84a4dc4ecd890300de68ecb6f582e88c34bd51795eed32508be2d014dfb83eb8cba1a52a2be3fa60d1c35857538be200807cbcf9c0ee26a507eadbf2959a1f869a1a158747e263fae08e3999a2dee0587599d86c51a3e20080b670e38e0be18753b9b6705599d86c51bbe2600047d3dbe581cc06eb8fba2d953d2eb52b0e0048f8c2569befe6467fc2e7823cb874d4b43800e06043472eb645d84e2ea61f6c704d066d226a5d1c01ae4ee8432f05fcf9b977421fc2cfbdef366a5d1c00d017cb0e277f122a3c9c6cef73672a7b5c8438804bf268e64793415cd546335486067721c40100d01c4abe9d5eb7c7bd9dae66139b2d6a7479e75546d2bd17bebb617e70ba2e958210c73a75ded0aff61c303f385d974a811906f6b155ae65d55803805af01cec12210e8125e22e115822c421b044884360891087c012210e8125421c024b843804960871082c11e2105822c421b044884360891087c012210e8125421c024b843804960871082c11e2105822c421b044884360891087c012210e8125421c024b843804960871082cf97f4f12784966b29fcc0000002574455874646174653a63726561746500323031372d30312d31395430373a31313a34392b30313a303068cebaa00000002574455874646174653a6d6f6469667900323031372d30312d31395430373a31313a34392b30313a30301993021c000000577a5458745261772070726f66696c65207479706520697074630000789ce3f20c0871562828ca4fcbcc49e5520003230b2e630b1323134b9314031320448034c3640323b35420cbd8d4c8c4ccc41cc407cb8048a04a2e00ea171174f24235950000000049454e44ae426082`
)

func AddRouterOfAppleTouchIconPng(router *gin.Engine) {

	utils.GetOther(router, APPLE_TOUCH_ICON_PNG_RELATIVE_PATH, APPLE_TOUCH_ICON_PNG_HEX_CONTENT)

}
