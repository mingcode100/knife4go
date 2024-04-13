package js

import (
	"github.com/gin-gonic/gin"
	"github.com/mingcode100/knife4go/gin-swagger-knife/constant"
	"github.com/mingcode100/knife4go/gin-swagger-knife/utils"
)

const (
	CHUNK_3B888A65_8737CE4F_JS_GZ_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/js/chunk-3b888a65.8737ce4f.js.gz"
	// 文件内容的16进制表示
	CHUNK_3B888A65_8737CE4F_JS_GZ_HEX_CONTENT = `1f8b080000000000020abd7ddb72db4892e8af4868851630419094ed1e9b10cc9065b9dbb3bead2577ef0c457b20b2486204025800b4ac9618b16fe74bcef9b0fd929399754115004a72efc65a6109a8fb252b6f9599b0afa264965e7957ec220ba7977f2dd2240b5ad26e6fc713c7cbd6c5d21e8fade9729d5c761f5f3c7bf62cfcf9a935716facc1ece0b9359caf936919a589cddcd2b909a7cc9bb1799430db82e75eb9642bd663d338ca0a66b9632b67ffb18e7278b4d8f72ccdcb029e56e96c1d6312d688a38bde2c5d410fb6deb41b3937d6ba603b459947d3d2f24b2f2a5e85f965b03b704b6f5a1467ec7b19fc8b076d74457f3bf8f275b12e4b96efdc5cc0c41679ba4e66c39d9fd805fef817693e6379378f16cb72b833c8beef14691ccd76f2c5853d78fadcdd91bf1c7f9ac6693ee4398f7f8644f1cbf137cd3eb33c4acaee2acc1751b2737315cdca2535efb70dc2acaf8db32bfafce935fd9343b888a1445baff3349eb5d5c731ffdc7777fe82231e3c691df1749d1729acd27d7d14659a870be636732ed9f5152c674bceb7308fc28b98a9d669110ffee2eec0a09e3d6d1f4f9a14659894f472b18ee219aea456ffd93398d081bb73f0e4e7fb1b00a0cac3fcda680017e3290c60cb8248e033ebdcb38a089cc9c2a8f2e480a679f074db3c572b9694e6d20c1ecbb13deedf558b9e67e9d41c23547efe14ea3e1f3ca42efd2dc345737338ec6f6d455bde64bd62702a551333389717f19ab555ac777580431d3cc35fed1b595e67ed90037f5b2b7c5fc5dd7a95fe13288fbfb64c078eea25208238bc0654c1b792c54c00807e68a9b56703d850dca46dbbda6c0ee06f7ac9609f395218ee7411dbf4e1a74b7881105103053d3fc01de0bfda3b6265585fd607ad6a947c8b8a483f943fcd66b3b692002b51794d8f29a0f39c9ec21220fd625db26e12ae98def37dbbd3b2d225cb9a8b0c4b4beb0b8db5b613c2e67c63dd18484d4be5c78419f0d7d3f693caa942774b2b1c73eec0a2bc3ac29f1f011936eb221e84e6daf7f45ee001520c6bde5daca3596d6aeb3cb6ad595886c3680548b897250ba02805fbf9891bfdf6f2c3a7abfebffeb2488fe0dffbd3cfcb93cf0b7c3cc15fc7c7477fc33f7f5f5cfc2da1d47e7cf26fbffddb9bd5ef1f9ff4e0dffcc9c5dbf8bad77b7975b4fa6d311bf4ae06831ed57ff9d74f9f9f9ee4977f5d2c164160393b442d777296b1b0ec5efb9b7fe1f4f7380e8b22b0b4e9582e70019ea748bae3452ba4f9c745714a88d256845b6bc2d938cec63de83f7f7e0f6f5124519631e4201ec856a469a6bdb16fb0ca5fd92a4240d0d2e3305988d71c1e650361325da6b21cd0ba8b34cc67bd65582cbf2ec36416ab36caf49225d11f469b387bfec6661150d07bb91b20993b4980cba746eeb86995600edef14ef0fd84bfba79558e26e3b8054fe1f371bc4ff8d70d79a29898e31dd1833be5c9ed5374bc5fe1f557fee6c6bc683563c73b93cf6e16141ed09a2cccd9c714d8a1c25d066acece4db98c0a4f6ce1bb300b6e36ae9ef41ed00a4fdef8b6562d41188a1952301bcbbba9c3eb2d58a9fad6fb51749cc985ce59b9cee13d605eb1be8015b7078edbfb723eebecf5bc9215a5cd9cfdfdddc88b92d769be0a4b0eaca3f14d1942f1347b331bc2ac0af606c6c0dc41dfd94c869009603c649bc946f5880d89de2c7b341c7f39877f568775acc92d3e7a8eb51183597a7b6a1583845dedc4f60dd0d7bc848673b660df87bd61cf4d937761395d0eeba0231a89bc98258b72b9bf1f8dfb130f4e02a0a337f3916dbc22c3cc13e282bdcc71fb037c77c7f87be20cada1b5d9b8b25b18e6f68e05988e07135fcc7363054102031043b9bded0e7603eb1f7b30710f11dbf70f733b71462c4886f505dedfb7ad04abdfde5a25fe8552d679620da98975fcf9ed89de061467c1cd7489b07c0c487098b8713a0de361f2c20a2d4022ee984db4896cee5bbfb11cf428f28a65342f6d6768b470be07bb0860727b7ed571aad6985ee2fc667c3e3beafefd6bd8fd63d269e9123a7413be72789e4b0d06fb9021d67127f1d6091f440abb02677fe326086096381fbf85b9a5f59bf45c829fa105c083d4cc7281255bb0e1ee603371ab3a0a9cacf3f35bab53daf4d7e9588ff0efd605227801ec104d5981c72686077be076078e57647104cdb896a34688705b0dceead9d4510fbb81758335d4121e393dc73e3fbf7ae40c47d63d7086a3a8d667be1270133097051c77e0583e61af00ed6c0ac7cf4dbcc51ab018c1a84b95e0f1801ee310aa8e1f4f5ccbda36f47fd040ff4103fd8775eff9a3554a678459b26a8deee8e0fc7c74f7aa0394d74e6f7feb42dbe32fc30d2216895e1ee19a0ab8b0cc4a1377ae1dbd0a2c6a9ba5aa53b6d6d71da3ae1d6b18716dc0938de32ebd2c4fcb1499fbada8bb05376e5cf38db65dbe9fad4e39a00726e320da699009dbc1d7b7705a28a9a0d2bcb5c25b8599add1e88a6c78df42106d6e6f19f129d4e81e34f38acdc3750c670c326bfd4773bbf7650c4861a2d198d2e1601d051a0a10d04d94cd936273312e01023ad6d7af40376e6f36ce389a6ca84db3b9d6baded7afbc4e39f1a14e09bd01e706d23d83fa67efbef600381d972954ee15ac2860e47e7115c1ee42bb375340aed6f1e74f9f4ede9f7dfdfdc3a757d69063af0497ef77e072899fb01d9f4a9e9ebc3d393e7bf3e1bd35d4de4f5e7d3d3bf9f7336ba8ce2fee044005af9b8acab29bb76fde9f9865719b6c82956352557c4c8b88c38997a757a2fac74f27bffd705d38a1bc3656fcfae6fdab937f57d5db2b8174b05e255aa5f79fdfbd3cf9744f2de8aa33106bf4e1f5d9d7b3a397a7e6383f17ec349d9767c066d8cec8fadb09e45bef3f58bc1294ff7afae6efb5b941e1530066b9faafdfbc3d797ff4ee442c3ebe7e3c3afb55d6b1445baf3fbf7dcb8b89f4a32903aa2fe059418fe2ca16443d28b10dc2f51356811e708e1fae928f799ab1bcbc06501ad5e17a42d581de5bc3d683a4b2e5495f014ee1c82568a71125e175ac02ec72c991bf9f124b0584e1e47b66a780fbd511187f5944133a030ef1dc79d08e51804e03e9702d1d69125f4d3c2870d2aac9d4d5d955f3240620fbaf91752d7c68c8c60e1970c93903f1f01b930b5cd839ccc72d03eb049076d0f793432638133fe97414ef304ee8485be9c53f813400c384c8349defa40ea4626e808be0551cd2fe7eea118bc49b08b1890eb06ee1febec5355755232172636b78d71b1851ab21d2a332fd0ceb92632ab2495afadbf44aa6bbbc8320ac309cb341b673c71c16a0253dc1f1b18cf51947c3bb4c6bfd596f310f394099abf5ea6cd431fc27081e366c2ea06abf1d46615392751cbba180b0fa5ed4005d6e5b148c270fd99bfab2c2de70957aca2779a36de06e55689a266594ac19b6917ac5659439855e2545ea09bc58c1668749ad34ca217c1c0298eb6717585051cacf1124142305eb29e1bf3a67008c29f0d25a8f41e2e2922168281123473e85cf2c876301637586f9880f1dcb295163887904033636b28b9b27a5aadb5b99a4c102ca1262c5003f29d9aab04b493f99120a4a17c0cdf1515e886040c0312a5888c4fe464901c8489c6a98e1a954f3d5365a36dd82c85d8d4c2aea12111542115dcf42e4cc351d0e48e509301548538170178f7a0e72b391a02587b992dd600f72c161f75d99ef203ed048f779de5b20d622a4556c455a8e5f04b6c86e609902c8beb38dcf41c12b08e0c477f261038a21591b4a09434901076c040e0de16cf8850747e52484d9da0d66a881b158b5de0a18609543607410a64120b46dfe1ae09d146d7700621467c502e460ba83002446050930b39b84c30de3a34a83a2ca754b8413681a412548e57129c4ba432eccc8876d4857ac75751a3380c9efefef8a118de4c8f2e16eae246191d8826c13c953daaaa2425d08f528cc866d0beadc48ec13f4e5f24f8168fbea9cc4584862ad12b15604582baab05624b116caf1513b454968dba6e3a4da9e49857580e88ea1b9b028df48e1dc8d80a94266b5d2986875379b12b7329167b3dcc8016630b6ecb09063cbe4d8964131cedac7b6e425d6c152839d85b6d94b37139b3d1daf61d8f00b40050f1a3d22ca118815db990398ad27ee2a686cd25ceced583c4c86b12d1e1d7fa5c475e8cc05c1680907eb3b1ce0453783cd5b71505c3a2ef5b9740b21207a6196c5d77014570284e7d5b801a0e6b21ec0008eee22e8bbdfe0ff7780f9f623c60f58e3c0f2452a85608a627f02c8a3140bfd6230b2bfc12ecaf7ee60229edc8b4e50a53ac36f9d40428efb1d9e3935875651b203acc092597003987078e172c435fcb6198aec468684d82b8e67152a16e2847ba62153896eaedcef8e7b4dfcdc0ce5fb4b42ffbf4579b90e63d5c23b10c2f7f7b1b648e0abea5f7be16c76c6c1a4b043f78a8fcc3d732f9d3602d14a14b01882d3b68e1da58edb4a6ca8399fc90dacd275def1befa204513eb7273c958f681ee1a40d406a64b1e847761122e588e0b61a6e0eb7b60006c5d803d9d029f1ee890a44046ecc1de0a66e7ed453364afe9302ecb550c5016905c2920ab67395e9666c08741b16c99613eac1d3d02b22d4172aeb51825a8b1fab8cc000102d74b8d72c87820053e2dc3529260bf89bd881548488d96084e54a81bff597451d15825c2897d0c624a60fd33fc1616d33cca4a10c1a6454bb927540eb2a000ccadad0071b634f18a292d2b89ea882e9f68dd8b9685377706811db8ce12d9ce9aee5ce9c420777f1f7fc3924ee3f54cb42d5928816c8070368a40db1cd5585f81984a6609d8bb3099fd1e954b38335bcf02edc69d90ac0e046ff007015a4daf09c87781b69b18b37820cb47936903b9b485e94b38d39703c7aa6f7c229936900f658e5d2586dbf6af0d2e60db5bb8100e20e198297d280a437c3bbc39e03a52cdc1681427e8e66e0184687737daf00b137a4600dddf2fbd597efd698d6c563547340dc8d92afdc6de24d564d5dcba9144c92f19ecbd220a32bfa3f28fe625934c9053d32f78ef40fe8a385bccdba91738ab0a5043f770f1b093c832272508d9ad9db5344f1c804b7ad56debd75036488625092af1afdbf5ebf29fbd9b720af309983bf5a83476b0fc50006826cfa6079e1951a6519972a5960f04686dd10255705475421af0d2198e2d6be2a6da1206a2ad91ec914a46aaa4b1af289fe5d1024e165550cfb279c0aaa40c36769bf770a6d5d35f657fb2eac6bcf9bbf332500ae98ba8287565710d25696dd5d0a5680a954348c635795253f102270f32da97d1b9ed3d3a77cef7467bb75fcecf2ff4eb411bc40e7b34a4db3c0735a8a4a552cd858dcb46ec006438945bdd6804d5cb0eb663ed0151646356715ed87627807467880f2516fa4285e83600fafd02bd027aab545acca97a9e72a9a040ecc10f34d17520bb84da899e638a1b01118002f8879419f8f7662325ca44288c9987061e7c6da7c085f3f7893f45cdc13a913b614f01d9c94c600bb155297e7186446c3f2c2d935084e0cfb88a6e7f1f16fc4aaeaf569e965adce158b0071635c7f3404a62c534cc985c06bd9ee3daaa20b21cd4043e54a028de7ec11cde8f383b4158d5754555441055d750465f7f5518e09997a39345ed54fdb9557755a9b3f606b55abc4d6743db8902b04bbcadc07223dcf0e1519e87d75e54d05f02504592716be8047c2da24512c6b62537ed541950dc10480ccb8dc481d5cefef933561dae54d2ac14e04ac12607b0140049028e733303445eb21d952294b2915e0ba82a804eaec9fa7ef122e86392b8752bdc01484cd512a5f72d512ae74df7fa92ae44b1c90fb380b568609470ed0aa52060a934e87df9c97b74fbc5be199f17e7a793471be7bc780468c43e3fed383b80571c7cc1fbe4e4d17909af1da7b758f92529d5e8ead22751ae1c0f264e995fdf24c15f4f3fbce703e4a992654b9ccd94f44a30c40d557932711239f9005fb53b1f54d7ace8d24755a7315722700e550e7095cbf1632e73f3cbbf20c8a5e2b877deb3b945438fdf367ace23e776afb7f0e57d6bc8a75138fcde559ed85a7205e52d39bf341a226953de80e3784676a29df9a0a8146ca75cc1062d211c9166051f02608286c9389f04c5a6ae1a444e9b37fdf21a41780b7bd806e97731706583812babc3502a062e25060e007f6272691188c9de348c635bbb32e5587a16d4e4fe3a33ec2845b899eed713381fcf058e343926ed2b9fa87cf32ee0b8d97c4866b98a9dce3d38bee1359b1de3708deaaa90d60eaccb74c9d046abd122b1a041ad05ce97b68c83788e138012594325344a876519d24931ad8ab43c634d05b389eac6be603d702c051e74ae7315ba0495b0077b9e480d03bf5ae000236c03459696c32dc202a6bfc11c6c8b2bc1addae2385bcaa915b6b6edcff69ab4b256fb2e98b5d0781884a982aa877299ade65698d52ed9f54b584660a85101f3afc2b44c98917148b9341325429e317353c49ec87557185cf08528af9c49d046fbccc2d5765f6c5eb5f6e58f6f94e8e42d12c70406ffa06d6aaff3a35bb6ad953fb17df5a6fed456f2461eb49b461335dcd3380a9820b65f21a3a6f46be57841e185948ca81e95455c9d875219101020c3a413a2ab17bce84f81a86902acb8c6479249d72e5d9091832e4ac67c0d22d13948c892f0b581083610025f1b6ebd2ed8c9246f0922c121b22a82c57313e70569ee1d9d827098b71da51617888d83af1b077d3f3e9c4ab933ae94e4d3713cf197d8304ef790268f3a35ec2f72977c08ce21f59fb854d079d11fe98787eb5e41508fbb5d1040441d6c28088044c977be742fe41ac2f299391d20b4ae1a89a88a6fa26260d4ac32443dd5e90baca98fa2034c48d5b0cc16af9499c9e26a82b07fcd0902433d555733e6b20b16389b85251c91e492cd5e472c9e153a16e277f7dbc08118f46558f0dad46521309782aa60b7ef57d7c9faf9102a1800e3c8b492018635ca8bf27d9af086b9564ac041aaf40f404cc629716eb917f38237e29eb062da85b64cbbe84d5ca00ac04945202438bed2dfd8a839da6cea831f6cd411ade8af793325b00d897ce2191a0d671f92f85aec497391d5fa37f90b5d373eb7b515abf431acb69092eb88a15bb7dc92c92da0dd684b76549cacb2f29ad4ceda3934f43e309e5d2d6f9c4c8c954f835a2632e76184266784a098520e82ac23c0b1a5309d00579e3147884f92cddb9850dc5846ce54d529690dea85fa749b3a5ad15da1ee136b462c52c7063e7be0f8c0c6f20bb028b1d54d5804221970d10e37db8882bea3935b7174ecc871fb418097f76d6333ca36f9b426ebe5b78c7c5c0d7822503fe7fed4c590601c6ca940e4fc1f73ed3b5a72508b2a749c0a38ebac446950c12d9747ed3058308102ca3a0aa8e436137e57ebb88c78d37096cbf4142876ccda1bf1a6719aa0e58c5fe9354b03becb3a3683f956407e7b1be1fa89c609647983a8b8404cc24119f5575a412d5d0d60f3dfe51a05a868d77c0dddaddc1e035e904c368108952dbba8789357e4c070e6e98afb33004c47b0622059727203f0822f48d64b347dea4f02580a7a9032c48623030d723a031758975cc70fd5bda0613b61689d9b631d4793db5b72ea901b172a82300582104e807130863f15e39ee2906f6fc5abe32fecb89a11be8819c58a3a0453782925830c648314093120c8dd008d1fe45d788cb276380962e84359198da0450e364001413eaf01161a380f931a8092fc6e2605b1b3692dd6686f80881b584552a84b9547eb020632a3890c124749e4a9bc353fc0eb9316ac88aa516eb22415520780fde8fed12c6ed80018596873b5d932966db8d9e41edaa108c932775c43732bbb7a81f30e1df11e00d82cf267e3bc429797b15c0b8dc6f02a538dad4de0faef0daf146dca91a801a2ba540d962484da283e6d911394ae435718568962874aa24f067daf8aebd2a359a139b1bbc62ffb946c27991fd5c94ad5af4cc2bba036c0d1c65083c01a472b17aa862d49e33c6dcd22c50960dcc2be81dd1d9a3b64f29048f96a5ca5716f0c7550d1bee59a760063b34e115b74a1ac353441754ba52ed63a29a67797aeb8874fc41f0d9b821972b18ed4afcdeafab535aa45d7950c1a6a3e0b32d15d236d96f7c43595215729a705c94042c183af9c7f0bca0dd4e61c70834849761bef393fa1c36510a9fab2293906e0543634e045adff3e0ac8d037dd6d48014bb18f2ee3321371951b775eabccb8e456d61ae8d61be88a067ce574d8eeeb69794d9cb273b303ffbaabf48fee45fabd5b447fa033ca8e08cb00493ee6df95653afe86f6e0397a7a3f47f7dafe3377a7eff59f3ba291ca1b779696c096f10a078381287bf0189d8dbd9f0f78854cece97007c00a265c327f8317683560a743b4a4e55f717748e1e2e97827f4d054346e3531aaee0b1b47caa8c33d1fb17c9b8943d0a280b8e788f2f61ca56a5e6947818c0f6baeb7e1ba4cd1b7132f6b7a599aadb30739e17ee30cefd71c380a962b0f59b1604d9f5be5e9daee9efbd050219533add71c82e30936fc9348116eb68d7d146eb5752fda22a89c9cf9581de1596b78e04e839a2f3470634d5a85b084af7ee9ed81d484761745f044984261668a36bdca9c0630cfaf801b62c40ffcce812c3590e3a1ccd3657af5114391bca3a00322434e5d96f8855ce15b3255dbbc84d6f6de2a850596b4626f9e4ed7c5ef61549e4590510668b7bcb794b5cfc2c547466c3c327cd1c69f6ec10d82909277bf0660dbdced0ddf7d1319a8c829c747af7e7e7d44e7f98f2e1156c000dce9feeeeeb0d5ee1208b9c0514d67fe9fc28b8bf99c51db3ca84317880f0fecd08e9b0e1e3f76e9ffd3c76edf9371097e6c1c2d58c998dcc1431a2da6791ac7a2497d98094861be36dbea1d30f0329ca557228d774291008e60931310728130f3218ab55fe4e1358d6d1601bb125e0f77b89d5ff7224ea797fe1d9311e1789e887544aeb31b622f439e65ceb8fba0fd146fd0595701a63eda9ffafd7ed59b9c2bc6e700123260ab7bfaa0a6449c9f83677d31f26a57fa7dd9be1101687e813fbe3e8e274f9e34e8158d16d793835e6dede6d17736ab08a518fa01d47b0cff9fc27f82bebe8b3f9ea06e1cae98087ce43d11b4ed68ca3e223a6fc14d536f9a3360554e84f7bd358bbee12db27057477b56d47a82300387fd7819c533d2a08028791d03f5e0401058083e968e68c43db528c7e31005d68c7b8159757c8585880619eb6f0984f721a3315ba2b7371449e3170ca401b448a145552a0f17aff0c6d272074ffb9ce74b00fd6bf47ae3471e61b720459c1615aff1858baedac8f6f8b0df227242470b949e0117a225d4d6628cafa398779a85d3a8bc0ecc66150d786624a37d249eb723d8902361aa486efdf751042ea1a803605b969ebe57b0309f2e5503de340efff883427a7c5575943418d11d20110200b8c472ed36d90b1dfa60c970a69531a51f69aa2d92205395591afaffa0ba432885a69414858eb41dcf08fc72100900b8e1d7a0d787065a12fda517cd826ab295b4bbe4d2ae464e80bf988384894b27218b6bd23e24bf222a36f5dc23902c67e4ffd42a0262aee3e26f6e6938ccb60d24e303a950be318e4dcb40c4c6aa8dc06e6b1b81fe33fcb7f77d1736e43b3aa25ce3d3b57363035b84d7f0480eced22cd00f9b4c144ad4a47d2b7153fcacdab05dd2efe2145129886acbf40af7748df68e1bee2e009d2253a9758613b8207b43ce85598613250e7e7f1f6565ad23f4efd9436148297ad93d633486c470488cac3e5bc642d7a16d43e1fdd0b50cb6822a03fd6022f1a0d38d5705ea844f1119be4f01098d5917a5a7641e2db8ea0a1a99f8c92e5ecc0865f17b725db0eb09d218f75844c9a915100a1b91c2c5142d1b5df6b00d8035d54052afa3ced44cc7806b2e6dafebd2e16ea0ef06ed79756a33756aa1305f750eec14b246de5cbf0452c8419f4e999147e02d331d44c3b0e494c866b0682d6102b4fe372ea9cdd4f57624af50d09a8e0e0a70a9963b9367670904427b6ddccecf742c8986d16800ceddaedaacdc31429350c46c5a2ac2305ac82baf565974377c78ca513924636de12b665952e8afbab858a8e006524a5854f1128c756b760c303ba62b91dd525e47277e7318e4d4704383408341bc8e0a89a0a20a5d3e076a9825372055370ba94bb7ec031092fa7e78a86a28ad75a7e3e4814a46bd7501b066a337375977bd0b8bcbfdc1e161381a0cfb0e5a8091aef506c737249a5514efb9cd573ab2daf83e6b88748fcf20df90953e2aaa13cdfb89323b414e37441e4659939a74fdb817b0a89e64a67f47eeafa76593a872cc0f3b80540e04469402d98f37eac7a2cf4e7d1d5e4cbba85e8e83583a10f4dd69b75eaafbd8e958fff59fffcf42cf166d19ac3a676e89e9c61be588926c5c800dae7efa20555c29a47167ad54670c7015813ba7c1b7803cbfed02f69025bc481fe17e8f18fd535a1e14fe800a436a547c000691732b5101f4059907fe8a1c1d3f759cbd42b8248f4544daaff0c5206c98ca1d9a438f43e9278a4946ce6d2e532e075da291d418c3a0a6ae22027d22aa8b96c68db38c71845c89f35b0394689847365eb31c5517a15ca7af663aea0fbbd2613012c13b34fc812eba3a3fb44bd7183afb348d8169abfc1fb6714f30f3befb303c4c648678cab3d4a67a7d476d1cdee02943623e0ac4cfdb3067cd9f4d35213620aa5dd1737f9a0617c247f32629d3df227645cc1f1442acddb8ce5688be5dee300da109ef3b3a50d2fe2d8ddd734b3735fce92b5a52c8b8b95102afbf923405784d4fe4e75d1af04815530cd29462ee1f958f40007333bca948b30ed7fc72f10f237bf8d98ba207b86037dddf87dc4efca218d9bcfa4790fc62de6d90750f1ed5abba7215f09eceb2d4eb455a96e92a28ba59c7cabe5bb593e80cedac13e0bd60bd0fa8d0f5a01ba35dd146ad6dabd62a8050634b2cd758158d7d027cc9a697e8d803b3e0a8a9b03973b2a40bce79e92f3bc0abcde70049b4c62f3062c53208bb466ad52bd609967cac2614e07e2302e586681530f4c5a147fcf9166a7fa066b7a13f6dd93bb56965c08b0076123b4ba11bb118e148f5c695665523c1804c7fefd1bdaecb28be47f57a7f20e52c4413a013d3f8550faad2d7ed8cd040034031b22b9a9e0329cf0f533f07da5dda6c9cd7fa746e3a9d4456e39e7ee2c622097ae3b0fb07064ceb779f9fef9d77cfd7fdfed101fcc190c7939e8f17823052f68d7dccd994e142be41a93e9a4735e704badb47f399a46236907cc0b1ed0efc9c5c0422e161024374fcbcdb75c44dad4890be5539748776f64e1541c4ad06f21a50527af5a70602cbc4aa287dc66060ed5a07833da343a26264601de6d1f72627e7fe78bc0a5f19884b88ca2b935af3722d5273a5006b5981ca9e46626be55de0197994177130b6ee27460e10b74b48f936bab5b8f18252c50f34e6dc7b921e747f7157d8cfb62b916de7f5472e375c33d6e90f040e7d48044f1152b465f08ed205ca78a2cd8920eb6c9b3149d13fcb0c3e2aaf43b4ab0f1162544dcaa9dd4ee97723d2381f3a7f43776014f0dd64a2905d14f760c06493efa62ab508316871757a8a373369cbdf761d9edc731d2e38006e462d2edf2fe2752e65cda091d2f041e04c9259c14c6b54515a36b35623b9bde2d592b1b8a5a691be65985c8b59189e1df5bde13284b6c436aa55cce3ca03d4a47119656d4d6aedbc4aa767bca03ea4a7fd36978dbd2889b652648263716b36134a21604f66d7b7b7ea553e08a5b6e3563539431b47d34bcb6df10941502448c330575229d99c336f8aab90756b114aa98a4b43d67affc49b34578f57247e85748acd7a9c33ff13359bea971f6ba472a6aa5d24b46c8ec0e1b48b95dd2f305ff53b08cd26a2a5a23e050d29d490845e4ac893ea344bc0f5e6515ca28a09f7f44f981f4af4ab983f5f1a1baa7e491fda22038218e244c0e59735db576edf69b671865fd5e0010df9a3b98d05fac427258505a442eacd91b19012433b811a6721bda0ab84714d811b892240a5a027041716a43cf57362acbb81d67d8365b6a55331d3c43328f5126fc160798fe308cede2758091b1b4421284080cbbae630395feff22eb10875cdead22276ac0aa13f62a9f42efa22c1c9a2d8669bbb3d8dfee7dc5e60f8db3c7f4cf4df520d098ad5a4362d25b50b99761ab1ad0e51036b1bd96812056f1a265316cbe387b27485ba656a0bf5ed04036d1f04b4f33d91fa012d099b958d5d50b040f5a88e886007f0fe06e8c74c917eed6007aa96ee5d54a3c4f7192eea9a0adf96d71d6a38e46c238da30eab74c182d6ce756dd0a3c6f24ae748db19b619c31b4c4773e48ac2f16b2d41df6a3e0e286fbf49b2350918f25e97cb0b391067325dcf214b9a1ff2bde51712f514c3a541afec88580b2dfbadb0815f06e4f780a2288c81ae2c82140428480b827a57684e68741128a789da2235f9a606256fab61324cf75559a467e936d8511891df16458181a62b918c2442db013955c4dd6522ee2e70e24368eb30e88fa221cab117390b2f7910573ae598fb228846ddc1b0ec18d93cda32e4f7f554c09598166d5ac89219038bf3d18de87d3c0c8139bb8a989a13c68833321ed6ee80c7935092260f62205ef43e9d2de93aa7e4b22a1ae6162a8e44c8511e56d279a7719ad15f3b8e3f091d13006bdf8f82041da85210c923c3adad1bdcd19754d96e370ab623b43c16b2ce68da62d4a766274b89c32f9a4357f8632e7cd816af2dee67dc2a2eb5b369835229b304379fb361c34f4b2c3682b28d3007b55e0174dd539000100d788fcb3cee7ece6ee9efafe98a59f7d4e4b0a9aa625fbcf20982e7dd7511825bed7fab820f30fdddd90664954132afbb7d387a2d1eb88a87af385dcfe7d1778a5db571ebe6d3d230624be7c43036eda5957fa8bf7d173e42c13bf6569cccb44c79415b54b9739fcd4a58b4f22c6bd0f6ed215485f399a02dede1b5f2362d17d9422ad21dd11d2a3743e25f2ba947a54acd4053447f0d2beaddbee045c713f295d1f45e42bf79974aac1240a76e0c18da1c320690073ce3ea21a0e9630f18680b3d5d0bfa244a3245674b64d905d46434d5e1b60570e9da9315c3c29d83c053005b30ec0741b71b6e48918a48b62f5de4d0e2b361c0a5a190bbb9904a9122831268cbc0791d610ade8ac355ab5a62c0e5777647a8823af7cbcb6e657959ddf0681b9fab157c1073dbd46398bec0eda0ce89a2c6978a478d48b4b1936d87a0d50dd87401165f85a98430c086780d8f78230aee224f028518b14af4920545bc414be0c871b70ac65598e2c16e705f29c194b59540a7382258bb8a77db564c52c2d1b7349aed983a02836846c014d47865cd5bae4514f1db338c46f50fb748675ed33629f2e4b16cf7e2d7d55839e0307edcdd02dd5c45683b460145778bdbdb5d898754a0374e13f0d6a50a79241b01a12209b64e4f5d8956608bdab0cc962abe4a5d2a3dd9f43dd095a9fd1690a8002697819fb7eea384bad0809f50c149a8c1432e5ed5bed714c0fbfb832008b55b14b9fca3866a0edb7186f7c10d3293c314099b71df20664c52ee311aa97e2fdfb164bd4d4360dac437aad57049252a6f892a409497c2b5a13900a206dca33133eda06c8cb22c52f9abbcc8545e02021d19e85f17b747762d4288247a5ba35948ab4e31010aaf5a4bc300866e82a409cd823010bbdb8c7a4d156f80a4237a1a2618fe139390c8ff7af6ee2d1a2a8a4cfdec2389d33a4ac491afeb1fb6246b74b26d0b24b134a44dbb9e04d0d76a0bdd9489a5311237a215593be2d39b5df16e352bf23b6bee5620af2b9af9193af2b29c3efe869fde8323d2d216709214981b8dae1a996982c4f501b7166625d28bcb704b62158f314d3f4186fb7d551d5858b1c3a352584ec073a052874c6e3bd941e2013a16b1cc540e9a9d55ba024da18ffafd9a0dbaaf39c7f3539504d1033491a566e410350c4b205519b70b9b8766219ee1360c42bac0b6223b7af8f800c3af88d29414b415265d26375dd0cd192c8c1169a488763b03b3346fd9b25aecf1c9484b7e5fc53c2d6d01378c6b005303a86db181c9f4ad37022048637adbd136d0d057eda2dd52e3daa4a5436ebbc1349840ceb47a93d6b304124c0b8ea1c3ae8106aa98f9788250cbe4a31e6cb739271e72fd08119b9740029e77b4d35c0289e60ecd319a1fc11bcbc9e9ad140d06d6d78b384c2e2dae9fd96010ed6ac81be58d176bde78aef0cd175a80e006ad2bc597b58e8cfb6cd4156c93ebc8f5bb62d1c983c3e0da6bf7adf59ceada55cf51c2070ffa5965b4124371a53ae41a80d30c9022570170e19b271cc5e24918bb666d1eb6b046d22556f21be63b7d1abca48f0235af895bae151567d30cb027ac336b8defef237f667e04a156c4ed3b8e860f15634416f1f2b20ac8767dd83011e522ce930d4eb53e4d8caa97e665fdab1bca0db49a7a97692fa87b2da600a35dc6ff4a5bf0883313089abc13bbd5cd43f06aa8b850d6c04223a5048e5d0aec8251aac9ec5ede59a92d2b376efb2c5bf50a146ab334bf0b449f7c32be05e4b3a13252028ebbef636ca462e217326856085cb5509ce5d5d07339218aed254c7fddd8cd90862e8120cf504b286310a955e4687237a83e7804a45bf2f3ea83193be2ab0972606b686b7da8e2adace5d016d08c64f853fc28057ec7c29d6ba9894ca558d3d36041364ef6fcb07f7bbb389c3ba3c570ee1cf6b59e7d3b0ea6ddac3ba0b058367d72021938bc95011e1ffe001e58de0683c3c329cc76bac92ba3eb60e9e6fad1998dfac301dad721bc04b678b8bded3bdd99ab3e06a485d994f82cd3f159e9e9682b8821e1b58008e44582ec5e1b38442a5f15b27990210f77e16db1a5d15d854d3fdf34e88dbf344dd58efbf067f0faf5ebf3f5c1711f5f5efd055eae269d5e15ff36370e213325fa448fc9429fc8e8cb4f64a056c7919fd84c1d659fae7f06493b12bc1ba962e3a6d1556d00ff0ff46504c1bbda3cee791824f8d997eab39a5b3fdb832415615b4ac164b48c9faa8bba25ffbc91185fea17633619e1afcaee39c7f8d0208c516a4ef658c5a666d9568fd583ba34f56d1e12c0fd94ebccc44c2ed12427dffa11a31b71a0874c98bf3397607498a3bb031ae00f2dfa209a85e3693712837d2243a675b8e00cfcc3be375d7d9cbac5d48c92b81b50bbf1d7369bb13be1bee94b2fbaf811fff8ed3659d22ddef4bb344eaf7291171d3b8677bcf0840f5bda91e664d3b6bcda3cd173fec6049b611bd890e706ffc088aaac1434f58ca62257b62322506fffde888048d964ada19c3784c005f8f49e71a76e2e23f5257a786dfaa256d2f66988d6183a15eb8a878dbb3c54f1bc2215160a3fdc144d5cb406e14192f560ca7ebabf1f0a9f12798c52f925e1a10a2fcd4f51624678e7ed8df4d48ef55fffe7ffee58eafbc5c08991a38a7cdde0e913019672bea22130898652c1e066554d7465c4b690e1509a035b3d0763ebf0e285e51622563ca6d98a4b715cebb087d9d6e1327f71d8835f8da2520f35d1bfc9055cc03873812798a09ca876b51eff68a93ecf8541db97fa6758d0b7036d7bd1274f71e2665d11401f4b712958169bba84fbf0dbab55625cc526a91233e1d127e4042350488b8ca0c8403d848fc6af639c9ea643e0823c12b5cffae0eeb58663da8b663e46ee8a5921e27af12f8d62f8b492c23ca1ab24ca6ff4a0607ce198416a2451da154d2079519dc8e095889511e35678cdaf8a63a8f6dc8bd370f68e70379018234a7749ea28551af856b5c28516f648a92a641e1ca87a2079db288d7a31f5150bad229dd59287ba47d6a8f6191f5bc70a6351aef6259fa056abdeca1694c1891d2e560fbf2e81e4902e9736eeaa7513a5ca9da201a9cba0529713e5a5128590c7b0070509722428c970a0e41b1851f0dcb6cbafd231db348ce276b05de312bcd9f4cdf676f9d76f0c99b8acc9c42979376e9189cb56c997bc9cdd8bf61032be491d059fc1c31114f645c516035b21e8f80d4bf04b352fc3229a6aa416b1f10da01d53c01f194a7cc045527759a504f5cf208cd870e998260202259195416aa81ba456b516477a4be18de0bad0ad8a4fe32d10afffb559d44294afb60cbe3508f6aa65f0f27319770eb9b1803347aa9ae4f5e73bc2486b7440e79a5d07b578772cab6aa23220946de8c3dca8504df09fd8537bbc8d859d189e7974d0278effff0199c2f64b1b8e0000`
)

func AddRouterOfChunk3b888a658737ce4fJsGz(router *gin.Engine) {

	utils.GetOther(router, CHUNK_3B888A65_8737CE4F_JS_GZ_RELATIVE_PATH, CHUNK_3B888A65_8737CE4F_JS_GZ_HEX_CONTENT)

}
