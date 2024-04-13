package icons

import (
	"github.com/gin-gonic/gin"
	"github.com/mingcode100/knife4go/fizz-knife/constant"
	"github.com/mingcode100/knife4go/fizz-knife/utils"
)

const (
	APPLE_TOUCH_ICON_120X120_PNG_RELATIVE_PATH = constant.ROOT_PATH + "/img/icons/apple-touch-icon-120x120.png"
	// 文件内容的16进制表示
	APPLE_TOUCH_ICON_120X120_PNG_HEX_CONTENT = `89504e470d0a1a0a0000000d4948445200000078000000780802000000b606a1850000000467414d410000b18f0bfc6105000000206348524d00007a26000080840000fa00000080e8000075300000ea6000003a98000017709cba513c00000006624b474400ff00ff00ffa0bda7930000000774494d4507e10113070b2e546d71c700000bca4944415478daed9ddb7314551ec77fe7d2333db7eeb9dfa72797c9652693c9249389464104311548c83cec3eeff3fe33fbbceffbbaa8a51b45d9808052140a584617c1c212175729515458c22d49ef43c3ae9030d3e7373d3d93dafebca64ff739dfe99ccbeff73da789aeebe0d07968b72bf0ff8223b44d3842db8423b44d3842db8423b44d3842db8423b44d3842db8423b44d3842db8423b44d3842db8423b44d3842db8423b44df096577c70fd8bfb1b0f856f4c407fb0b9f1cf3558174b2c504aa72ac56838d84eab7ebcf9cb279f5ddad8d810148330cd4b5c14c473216e26ed4e169bddbbe52d2efdfaafe56be731899875fdfec91b1b57ee00112bf7bbc5fd7ffcc3ef09112cf6185dd7fffab7bfbfb6bc22580c58c1e7de13032efc5c42c862aed65ce8d65d4743ab0f061298164b441a57899789be202ba7cefee3f25798270200c03f2e7fb572eaac58191d889749e32a48985f7730906868f5e6d7b4163aee511773354e99f0f375a071371ff68b96fbf5d6edc3cb2bf7eedf47b4f9defdfb8797577ebd755bb4201ff6d3b81bd16970ca1673b5b8476d7e99a9c170363e520df7219a0d047849a16149b4011f5df8fcccb955c403cf9c5bfde8c2e7626574a061899714d12ecea01aee9b8d8fb4bccc94d01eee6ae4eb7e4916ae850e44e17c4c119ddd3c78f8f0f5b78fddfce59650a99bbfdc7afded630f1e0a0edd14f89842148e789dfd92dcc8d73ddc65e221e62887b4dd89a2c98b9f400756f0b39447b41997ae5c3d7af28c5091a327cf5cba7255b87a290f2bf8112a03c0ee44b11cd2cc5c69566846e842ae96f00411b5216eca2b0ab8c4de6a5dd7978f9eba7aed3b93d75fbdf6ddf2d153c2b32317e51585b831eb898427b890ab3162aaacc003347f743e5bc5f4633ab0ac87f77b45df9aeb3ffcf8e6bbef6f6c6cb6bc726363f3cd77dfbffec38fa215e3fd5e9615fe6f030002309fad6afea8c9ebc57ec97de9f121352d5c290060848fabc42f3cd53b71fadca7172fb7bcecd38b974f9c3e27766b1d889ff1711518e6e51952d3fbd2e3e6af17133ae4f235b4ba8bb65ee66c6d158db8f868407464bf7d67edb5e595b5bbf79a5cb376f7de6bcbc76edf5913bb35013e1aa01117e2757651ded0ea2197cf7c11e1be692656988e0e0a57ad8d869d5fbdf4c1d90b4d2ef8e0ec85f3ab5f88dd14fbc31b4c4707676205a122c242bb99b494afab2eaf70ed74207ececbc2ffaaebebeb6fbc73fcc64f3f6ffbd71b3ffdfcc63bc7d7d7d7c52ac3082fabc48f99d2a92eef52beee66925029cc685b54332fa7c6100541073ee06519e1c1e7cad7d78e1cff70db3f1d39fee195afaf895683653c7c40787036783935565433a2a5304213420e64a732be30a69ad8e9d49163a7b70a7ae5eb6b478e9d16ae3f6aba6990f1850f64a710012f643c3aed0d1dcc4e5144804d0796f6b0419fe8dbb4b58b68dea534abc0a08fa531533a4ac8c1ec54da1b4228860ffcef498d1583594c490abcac2296bc4f0d7ae757bf683e486e83111228abb8761783d93db83eb31da115c9d3d0eab2e09860b49686255e140ee2acddbd777879c598c6ddbeb376b8d5b46f1b08f02226c8050032931a5a5d913c38b9da4a654d45079e8b0de1caf2114c5872f5e297274e7f0c00274e7fbc7af14bb1c246d87644386c6bf05c6c682a3a80d6aa2da15d942fe5eb21b778d5ff1b68174c676c6c6cbef5dec94f3ebbf4d67b27cd2ccd9f80231311001072fb97f2a895da63da4dce1694d42b222bd1ffa103cb7b594e788ef5cdb7dfffe9cf7ff9e6dbef851f97f3b23c724af74a7abca0a4da11aa5da109c07ca69af7c7308525225514e2114b866e6e6e5ebff1d3e6a6c8ebac03f150a9a2e03255797f6c3e838aa6fd060bec0646aecb64b4f0a9f6d384cc8602edd7a1256cc84f1332e27566849ac954b5c41a5fc7aee4e878d854fc7bebf3a552800431d300b3e84082925412cef2188c87b55dc9d1f66b618dd03e2e37b4192f77235510cf750935511a5370bfa597bb1bda8c8f8be7f0b6a985454c84f32f2690bf3cfaffba358f7a27e494eec5c4e844386f49452c13dac8bac76405a10591a954417a2a5a2011a9a21219633e8ac90ad267b11d56fec7f607e27399094c491d58cec3fa840320ad6fdbe763394c580300e63213fd81b85575b1b86bdc9fa9149424a62427d2b8825b4d6ccfa3359182b07801404149eecf542c54c662a123eec0a2362de16c4d31371fb172aac747023486311f49942d6ad311b79595b17eb09f8d0d4f46fa312509f05280863149bca7d181865dbc84cc544d46fa6763c3d6ca62bdd01eee6a68f500ced614e0bc6cc5548f022f2b2480c9540524b9a199321f09d6a8039442b99792254cc936a2f24fdc04955b307829592a8572966bd211a119a10773b524ded6a412549ee9d11d5c9457549cf928e9091ec485135ad1a90559ce1739909bc498c9756019990d60a77a3ab0011fcb60963f8490f9ec64ce17e984201ddcc3b237551e51b0b6a6b282f1023c72342838f3d18892de972e77488d0e0a1d74f990c172c3dd52149f3310e045bcf968295f0f8a988f84e8ecaeac7a74b02ee8e87984616b8a8a48a6038de2cd47f558a18e736099a3b342bb99d4d0b0b6261f13b33519e6231f666da9babc0d4dd87c2444c7f7198ea8e9bd2954c727e4a9c53a830df6a6ca233897ac693a2e3421e44076328b1bca8d895acbd89b0e44a6bca2e2cc47595fe44016354112c18e9db3296f682187b535a56436d83a9acc06fd2c8599d251421672532994f948ec419d7e80c14bc9d25810b5dc3216d34d6c4d8fcc47c885fb5810bb88156e872d04244f235f9799780041071a6aba37cdd86117c2998f5c8d7c3d80351f0961dfa6fbc9c8c06c1c19127be66e4bec9e5183d9f8f064046f3e12c23ea125ca0e69d361b4ada9b29dad8913a982341f85ddfe43b8d0390a5b8f91185492afe2d21686ad497b72f6a603d3f0e6a3573395415c320885ad421380b94cb5cf8f4ac4192fafe7f1cbab03f13ce33537419f3f3ed7b6f94808bb0f4689c9caa286b635b9d96fba6336eca7094ca68a11baa8a112f66dd085136876254627b05bf8a552800625d8d4695092b099aa89707e17d68282a60b427bb9bb91af236d4daac4c714e0948f2944c59a8ff22853557b74e74ca54a388fdcc20fc00a3ea9aab202329eb93b51ac58643e12a23b421bbd645c16b768ea4064264d05898c99d2c565153942b44dd74e09cbfb6373d909e4b84f31e508c05c760269e56e9b6e1ec7b63f5d69d3462fc4a092dc9fb6d27c2444378536d666ed6c0c318f8bf225ad8e59975a44970f187c3e3e84b435093219e97f3e8edc4066095d16da9ef8193e76681ddd3f32b314ccede97044784fb254c245c3ada3fb4253421672b5cee53852ded042ae86c9ef58dbccee3ede20e30b77286b67642c910731584a4f080d007b53e551f143305a32aa66903978abe915a1559777499bb6d659e166d292368d719574805e111a00eab182e84145cd998915903ea90ed043421b6b0aabdc6f41976f0977a05967e821a101605849091d26d7847de9f1611bd7f72de92da10921f3d96ace67f678c46791f345e7b3d54e9b8f84e82da10120e9092ee4a6da89643242177253b8fd069da3e78406c3d6d4c62e9231f40e9a4ed28b42fbdbd81765ec09c31c75dd617a516800a846facd1c33be95d9f848d59670a0283d2ab444d9a15c2d2a8bed5d8dca8143b99a6de623217a5468402544f6a76d351f09d1bb4203c05c6662c0f49749060209e4d90ab6d0d3424765655133756206a76c51ab45ed351f09d1d34203c00be6be4c520df7bd801a3c6da3d785366c4dcd4f35f27119697db2915e171a8c2f93343da76b7772d4e4373abac80e10daf832c9b34e9e8b7b54f3dfe8e822bd5e3f83bc3f369f9ddc1a222200f3d9c96e998f84d8194203c02be9f121f5e9b0e7908a3d19d576768cd0a12d817c235120f48d8e2eb26384068099d850ed372738d7a20333d8d3abed6727092d3369e9f199e48ae459c29dc7de257692d000500c665f4e95a19d2f0c74895ec95d9a8412329fad7e7ff7e703d9c9ae9b8f842098aff676155dd76f3ef877d8e5efa994604b769ed03b941dd647ef5c1ca16dc211da261ca16dc211da261ca16dc211da261ca16dc211da261ca16dc211da261ca16dc211da261ca16dc211da26fe0359d75550675b11bc0000002574455874646174653a63726561746500323031372d30312d31395430373a31313a34362b30313a30309e86ca490000002574455874646174653a6d6f6469667900323031372d30312d31395430373a31313a34362b30313a3030efdb72f5000000577a5458745261772070726f66696c65207479706520697074630000789ce3f20c0871562828ca4fcbcc49e5520003230b2e630b1323134b9314031320448034c3640323b35420cbd8d4c8c4ccc41cc407cb8048a04a2e00ea171174f24235950000000049454e44ae426082`
)

func AddRouterOfAppleTouchIcon120x120Png(router *gin.Engine) {

	utils.GetOther(router, APPLE_TOUCH_ICON_120X120_PNG_RELATIVE_PATH, APPLE_TOUCH_ICON_120X120_PNG_HEX_CONTENT)

}
