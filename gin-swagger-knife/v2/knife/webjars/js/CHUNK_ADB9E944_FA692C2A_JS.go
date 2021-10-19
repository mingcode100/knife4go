package js


import (
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	CHUNK_ADB9E944_FA692C2A_JS_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/js/chunk-adb9e944.fa692c2a.js"
	// 文件内容的16进制表示
	CHUNK_ADB9E944_FA692C2A_JS_HEX_CONTENT = `2877696e646f772e7765627061636b4a736f6e703d77696e646f772e7765627061636b4a736f6e707c7c5b5d292e70757368285b5b226368756e6b2d6164623965393434225d2c7b223139616130223a66756e6374696f6e28742c652c6e297b2275736520737472696374223b6e2e722865293b766172206f3d286e28226433623722292c6e28223363613322292c6e28226464623022292c6e28226230633022292c6e28223262336422292c6e28223938363122292c6e2822623163372229292c723d6e28226233313122292c693d6e2e6e2872292c613d7b6e616d653a22446f63756d656e74222c636f6d706f6e656e74733a7b656469746f723a6e28223763396522292c456469746f7253686f773a66756e6374696f6e28297b72657475726e2050726f6d6973652e616c6c285b6e2e6528226368756e6b2d336238383861363522292c6e2e6528226368756e6b2d336563346161613822292c6e2e6528226368756e6b2d326430616634346522295d292e7468656e286e2e62696e64286e756c6c2c22306533362229297d7d2c70726f70733a7b6170693a7b747970653a4f626a6563742c72657175697265643a21307d2c73776167676572496e7374616e63653a7b747970653a4f626a6563742c72657175697265643a21307d7d2c636f6d70757465643a7b6c616e67756167653a66756e6374696f6e28297b72657475726e20746869732e2473746f72652e73746174652e676c6f62616c732e6c616e67756167657d7d2c646174613a66756e6374696f6e28297b72657475726e7b6f70656e4170695261773a22222c6e616d653a224f70656e4150492e6a736f6e227d7d2c637265617465643a66756e6374696f6e28297b76617220743d746869733b746869732e6f70656e4170695261773d6f2e612e6a736f6e35737472696e6769667928746869732e6170692e6f70656e417069526177292c746869732e6e616d653d746869732e6170692e73756d6d6172792b225f4f70656e4150492e6a736f6e222c73657454696d656f7574282866756e6374696f6e28297b742e636f70794f70656e41706928297d292c353030297d2c6d6574686f64733a7b6368616e67653a66756e6374696f6e2874297b746869732e6f70656e4170695261773d747d2c67657443757272656e744931386e496e7374616e63653a66756e6374696f6e28297b72657475726e20746869732e246931386e2e6d657373616765735b746869732e6c616e67756167655d7d2c74726967676572446f776e6c6f61644f70656e3a66756e6374696f6e28297b76617220743d746869732e6f70656e4170695261772c653d646f63756d656e742e637265617465456c656d656e7428226122292c6e3d7b7d2c6f3d746869732e6e616d652c723d77696e646f772e55524c2e6372656174654f626a65637455524c286e657720426c6f62285b745d2c7b747970653a286e2e747970657c7c22746578742f706c61696e22292b223b636861727365743d222b286e2e656e636f64696e677c7c227574662d3822297d29293b652e687265663d722c652e646f776e6c6f61643d6f7c7c2266696c65222c652e636c69636b28292c77696e646f772e55524c2e7265766f6b654f626a65637455524c2872297d2c636f70794f70656e4170693a66756e6374696f6e28297b76617220743d746869732c653d2262746e436f70794f70656e417069222b746869732e6170692e69642c6e3d6e657720692e61282223222b652c7b746578743a66756e6374696f6e28297b72657475726e20742e6f70656e4170695261777d7d293b6e2e6f6e282273756363657373222c2866756e6374696f6e28297b76617220653d742e67657443757272656e744931386e496e7374616e636528292e6d6573736167652e636f70792e6f70656e2e737563636573733b742e246d6573736167652e696e666f2865297d29292c6e2e6f6e28226572726f72222c2866756e6374696f6e2865297b766172206e3d742e67657443757272656e744931386e496e7374616e636528292e6d6573736167652e636f70792e6f70656e2e6661696c3b742e246d6573736167652e696e666f286e297d29297d7d7d2c633d6e28223238373722292c753d4f626a65637428632e612928612c2866756e6374696f6e28297b76617220743d746869732c653d742e24637265617465456c656d656e742c6e3d742e5f73656c662e5f637c7c653b72657475726e206e2822646976222c7b737461746963436c6173733a22646f63756d656e74227d2c5b6e2822612d726f77222c7b7374617469635374796c653a7b226d617267696e2d746f70223a2231307078227d7d2c5b6e2822612d627574746f6e222c7b61747472733a7b747970653a227072696d617279222c69643a2262746e436f70794f70656e417069222b742e6170692e69647d7d2c5b6e2822612d69636f6e222c7b61747472733a7b747970653a22636f7079227d7d292c6e28227370616e222c7b646f6d50726f70733a7b696e6e657248544d4c3a742e5f7328742e247428226f70656e2e636f70792229297d7d2c5b742e5f76282220e5a48d20e588b62022295d295d2c31292c6e2822612d627574746f6e222c7b7374617469635374796c653a7b226d617267696e2d6c656674223a2231307078227d2c6f6e3a7b636c69636b3a742e74726967676572446f776e6c6f61644f70656e7d7d2c5b6e2822612d69636f6e222c7b61747472733a7b747970653a22646f776e6c6f6164227d7d292c742e5f7628222022292c6e28227370616e222c7b646f6d50726f70733a7b696e6e657248544d4c3a742e5f7328742e247428226f70656e2e646f776e6c6f61642229297d7d2c5b742e5f76282220e4b88b20e8bdbd2022295d295d2c31295d2c31292c6e2822612d726f77222c7b7374617469635374796c653a7b226d617267696e2d746f70223a2231307078227d2c61747472733a7b69643a226b6e696665346a446f63756d656e744f70656e41706953686f77456469746f72227d7d2c5b6e2822656469746f722d73686f77222c7b61747472733a7b76616c75653a742e6f70656e4170695261777d2c6f6e3a7b6368616e67653a742e6368616e67657d7d295d2c31295d2c31297d292c5b5d2c21312c6e756c6c2c6e756c6c2c6e756c6c293b652e64656661756c743d752e6578706f7274737d2c623331313a66756e6374696f6e28742c652c6e297b0d0a2f2a210d0a202a20636c6970626f6172642e6a732076322e302e380d0a202a2068747470733a2f2f636c6970626f6172646a732e636f6d2f0d0a202a0d0a202a204c6963656e736564204d495420c2a9205a656e6f20526f6368610d0a202a2f0d0a742e6578706f7274733d66756e6374696f6e28297b76617220743d7b3133343a66756e6374696f6e28742c652c6e297b2275736520737472696374223b6e2e6428652c7b64656661756c743a66756e6374696f6e28297b72657475726e20627d7d293b766172206f3d6e28323739292c723d6e2e6e286f292c693d6e28333730292c613d6e2e6e2869292c633d6e28383137292c753d6e2e6e2863293b66756e6374696f6e20732874297b72657475726e28733d2266756e6374696f6e223d3d747970656f662053796d626f6c26262273796d626f6c223d3d747970656f662053796d626f6c2e6974657261746f723f66756e6374696f6e2874297b72657475726e20747970656f6620747d3a66756e6374696f6e2874297b72657475726e207426262266756e6374696f6e223d3d747970656f662053796d626f6c2626742e636f6e7374727563746f723d3d3d53796d626f6c262674213d3d53796d626f6c2e70726f746f747970653f2273796d626f6c223a747970656f6620747d292874297d66756e6374696f6e206c28742c65297b666f7228766172206e3d303b6e3c652e6c656e6774683b6e2b2b297b766172206f3d655b6e5d3b6f2e656e756d657261626c653d6f2e656e756d657261626c657c7c21312c6f2e636f6e666967757261626c653d21302c2276616c756522696e206f2626286f2e7772697461626c653d2130292c4f626a6563742e646566696e6550726f706572747928742c6f2e6b65792c6f297d7d76617220663d66756e6374696f6e28297b66756e6374696f6e20742865297b2866756e6374696f6e28742c65297b69662821287420696e7374616e63656f66206529297468726f77206e657720547970654572726f72282243616e6e6f742063616c6c206120636c61737320617320612066756e6374696f6e22297d2928746869732c74292c746869732e7265736f6c76654f7074696f6e732865292c746869732e696e697453656c656374696f6e28297d72657475726e2066756e6374696f6e28742c652c6e297b6526266c28742e70726f746f747970652c65292c6e26266c28742c6e297d28742c5b7b6b65793a227265736f6c76654f7074696f6e73222c76616c75653a66756e6374696f6e28297b76617220743d617267756d656e74732e6c656e6774683e302626766f69642030213d3d617267756d656e74735b305d3f617267756d656e74735b305d3a7b7d3b746869732e616374696f6e3d742e616374696f6e2c746869732e636f6e7461696e65723d742e636f6e7461696e65722c746869732e656d69747465723d742e656d69747465722c746869732e7461726765743d742e7461726765742c746869732e746578743d742e746578742c746869732e747269676765723d742e747269676765722c746869732e73656c6563746564546578743d22227d7d2c7b6b65793a22696e697453656c656374696f6e222c76616c75653a66756e6374696f6e28297b746869732e746578743f746869732e73656c65637446616b6528293a746869732e7461726765742626746869732e73656c65637454617267657428297d7d2c7b6b65793a2263726561746546616b65456c656d656e74222c76616c75653a66756e6374696f6e28297b76617220743d2272746c223d3d3d646f63756d656e742e646f63756d656e74456c656d656e742e676574417474726962757465282264697222293b746869732e66616b65456c656d3d646f63756d656e742e637265617465456c656d656e742822746578746172656122292c746869732e66616b65456c656d2e7374796c652e666f6e7453697a653d2231327074222c746869732e66616b65456c656d2e7374796c652e626f726465723d2230222c746869732e66616b65456c656d2e7374796c652e70616464696e673d2230222c746869732e66616b65456c656d2e7374796c652e6d617267696e3d2230222c746869732e66616b65456c656d2e7374796c652e706f736974696f6e3d226162736f6c757465222c746869732e66616b65456c656d2e7374796c655b743f227269676874223a226c656674225d3d222d393939397078223b76617220653d77696e646f772e70616765594f66667365747c7c646f63756d656e742e646f63756d656e74456c656d656e742e7363726f6c6c546f703b72657475726e20746869732e66616b65456c656d2e7374796c652e746f703d22222e636f6e63617428652c22707822292c746869732e66616b65456c656d2e7365744174747269627574652822726561646f6e6c79222c2222292c746869732e66616b65456c656d2e76616c75653d746869732e746578742c746869732e66616b65456c656d7d7d2c7b6b65793a2273656c65637446616b65222c76616c75653a66756e6374696f6e28297b76617220743d746869732c653d746869732e63726561746546616b65456c656d656e7428293b746869732e66616b6548616e646c657243616c6c6261636b3d66756e6374696f6e28297b72657475726e20742e72656d6f766546616b6528297d2c746869732e66616b6548616e646c65723d746869732e636f6e7461696e65722e6164644576656e744c697374656e65722822636c69636b222c746869732e66616b6548616e646c657243616c6c6261636b297c7c21302c746869732e636f6e7461696e65722e617070656e644368696c642865292c746869732e73656c6563746564546578743d7528292865292c746869732e636f70795465787428292c746869732e72656d6f766546616b6528297d7d2c7b6b65793a2272656d6f766546616b65222c76616c75653a66756e6374696f6e28297b746869732e66616b6548616e646c6572262628746869732e636f6e7461696e65722e72656d6f76654576656e744c697374656e65722822636c69636b222c746869732e66616b6548616e646c657243616c6c6261636b292c746869732e66616b6548616e646c65723d6e756c6c2c746869732e66616b6548616e646c657243616c6c6261636b3d6e756c6c292c746869732e66616b65456c656d262628746869732e636f6e7461696e65722e72656d6f76654368696c6428746869732e66616b65456c656d292c746869732e66616b65456c656d3d6e756c6c297d7d2c7b6b65793a2273656c656374546172676574222c76616c75653a66756e6374696f6e28297b746869732e73656c6563746564546578743d75282928746869732e746172676574292c746869732e636f70795465787428297d7d2c7b6b65793a22636f707954657874222c76616c75653a66756e6374696f6e28297b76617220743b7472797b743d646f63756d656e742e65786563436f6d6d616e6428746869732e616374696f6e297d63617463682865297b743d21317d746869732e68616e646c65526573756c742874297d7d2c7b6b65793a2268616e646c65526573756c74222c76616c75653a66756e6374696f6e2874297b746869732e656d69747465722e656d697428743f2273756363657373223a226572726f72222c7b616374696f6e3a746869732e616374696f6e2c746578743a746869732e73656c6563746564546578742c747269676765723a746869732e747269676765722c636c65617253656c656374696f6e3a746869732e636c65617253656c656374696f6e2e62696e642874686973297d297d7d2c7b6b65793a22636c65617253656c656374696f6e222c76616c75653a66756e6374696f6e28297b746869732e747269676765722626746869732e747269676765722e666f63757328292c646f63756d656e742e616374697665456c656d656e742e626c757228292c77696e646f772e67657453656c656374696f6e28292e72656d6f7665416c6c52616e67657328297d7d2c7b6b65793a2264657374726f79222c76616c75653a66756e6374696f6e28297b746869732e72656d6f766546616b6528297d7d2c7b6b65793a22616374696f6e222c7365743a66756e6374696f6e28297b76617220743d617267756d656e74732e6c656e6774683e302626766f69642030213d3d617267756d656e74735b305d3f617267756d656e74735b305d3a22636f7079223b696628746869732e5f616374696f6e3d742c22636f707922213d3d746869732e5f616374696f6e26262263757422213d3d746869732e5f616374696f6e297468726f77206e6577204572726f722827496e76616c69642022616374696f6e222076616c75652c20757365206569746865722022636f707922206f7220226375742227297d2c6765743a66756e6374696f6e28297b72657475726e20746869732e5f616374696f6e7d7d2c7b6b65793a22746172676574222c7365743a66756e6374696f6e2874297b696628766f69642030213d3d74297b69662821747c7c226f626a65637422213d3d732874297c7c31213d3d742e6e6f646554797065297468726f77206e6577204572726f722827496e76616c69642022746172676574222076616c75652c2075736520612076616c696420456c656d656e7427293b69662822636f7079223d3d3d746869732e616374696f6e2626742e686173417474726962757465282264697361626c65642229297468726f77206e6577204572726f722827496e76616c6964202274617267657422206174747269627574652e20506c65617365207573652022726561646f6e6c792220696e7374656164206f66202264697361626c6564222061747472696275746527293b69662822637574223d3d3d746869732e616374696f6e262628742e6861734174747269627574652822726561646f6e6c7922297c7c742e686173417474726962757465282264697361626c6564222929297468726f77206e6577204572726f722827496e76616c6964202274617267657422206174747269627574652e20596f752063616e5c27742063757420746578742066726f6d20656c656d656e747320776974682022726561646f6e6c7922206f72202264697361626c656422206174747269627574657327293b746869732e5f7461726765743d747d7d2c6765743a66756e6374696f6e28297b72657475726e20746869732e5f7461726765747d7d5d292c747d28293b66756e6374696f6e20702874297b72657475726e28703d2266756e6374696f6e223d3d747970656f662053796d626f6c26262273796d626f6c223d3d747970656f662053796d626f6c2e6974657261746f723f66756e6374696f6e2874297b72657475726e20747970656f6620747d3a66756e6374696f6e2874297b72657475726e207426262266756e6374696f6e223d3d747970656f662053796d626f6c2626742e636f6e7374727563746f723d3d3d53796d626f6c262674213d3d53796d626f6c2e70726f746f747970653f2273796d626f6c223a747970656f6620747d292874297d66756e6374696f6e206828742c65297b666f7228766172206e3d303b6e3c652e6c656e6774683b6e2b2b297b766172206f3d655b6e5d3b6f2e656e756d657261626c653d6f2e656e756d657261626c657c7c21312c6f2e636f6e666967757261626c653d21302c2276616c756522696e206f2626286f2e7772697461626c653d2130292c4f626a6563742e646566696e6550726f706572747928742c6f2e6b65792c6f297d7d66756e6374696f6e206428742c65297b72657475726e28643d4f626a6563742e73657450726f746f747970654f667c7c66756e6374696f6e28742c65297b72657475726e20742e5f5f70726f746f5f5f3d652c747d2928742c65297d66756e6374696f6e20792874297b76617220653d66756e6374696f6e28297b69662822756e646566696e6564223d3d747970656f66205265666c6563747c7c215265666c6563742e636f6e7374727563742972657475726e21313b6966285265666c6563742e636f6e7374727563742e7368616d2972657475726e21313b6966282266756e6374696f6e223d3d747970656f662050726f78792972657475726e21303b7472797b72657475726e20446174652e70726f746f747970652e746f537472696e672e63616c6c285265666c6563742e636f6e73747275637428446174652c5b5d2c2866756e6374696f6e28297b7d2929292c21307d63617463682874297b72657475726e21317d7d28293b72657475726e2066756e6374696f6e28297b766172206e2c6f3d672874293b69662865297b76617220723d672874686973292e636f6e7374727563746f723b6e3d5265666c6563742e636f6e737472756374286f2c617267756d656e74732c72297d656c7365206e3d6f2e6170706c7928746869732c617267756d656e7473293b72657475726e206d28746869732c6e297d7d66756e6374696f6e206d28742c65297b72657475726e21657c7c226f626a65637422213d3d7028652926262266756e6374696f6e22213d747970656f6620653f66756e6374696f6e2874297b696628766f696420303d3d3d74297468726f77206e6577205265666572656e63654572726f72282274686973206861736e2774206265656e20696e697469616c69736564202d2073757065722829206861736e2774206265656e2063616c6c656422293b72657475726e20747d2874293a657d66756e6374696f6e20672874297b72657475726e28673d4f626a6563742e73657450726f746f747970654f663f4f626a6563742e67657450726f746f747970654f663a66756e6374696f6e2874297b72657475726e20742e5f5f70726f746f5f5f7c7c4f626a6563742e67657450726f746f747970654f662874297d292874297d66756e6374696f6e207628742c65297b766172206e3d22646174612d636c6970626f6172642d222e636f6e6361742874293b696628652e686173417474726962757465286e292972657475726e20652e676574417474726962757465286e297d76617220623d66756e6374696f6e2874297b2166756e6374696f6e28742c65297b6966282266756e6374696f6e22213d747970656f66206526266e756c6c213d3d65297468726f77206e657720547970654572726f72282253757065722065787072657373696f6e206d75737420656974686572206265206e756c6c206f7220612066756e6374696f6e22293b742e70726f746f747970653d4f626a6563742e63726561746528652626652e70726f746f747970652c7b636f6e7374727563746f723a7b76616c75653a742c7772697461626c653a21302c636f6e666967757261626c653a21307d7d292c6526266428742c65297d286e2c74293b76617220653d79286e293b66756e6374696f6e206e28742c6f297b76617220723b72657475726e2066756e6374696f6e28742c65297b69662821287420696e7374616e63656f66206529297468726f77206e657720547970654572726f72282243616e6e6f742063616c6c206120636c61737320617320612066756e6374696f6e22297d28746869732c6e292c28723d652e63616c6c287468697329292e7265736f6c76654f7074696f6e73286f292c722e6c697374656e436c69636b2874292c727d72657475726e2066756e6374696f6e28742c652c6e297b6526266828742e70726f746f747970652c65292c6e26266828742c6e297d286e2c5b7b6b65793a227265736f6c76654f7074696f6e73222c76616c75653a66756e6374696f6e28297b76617220743d617267756d656e74732e6c656e6774683e302626766f69642030213d3d617267756d656e74735b305d3f617267756d656e74735b305d3a7b7d3b746869732e616374696f6e3d2266756e6374696f6e223d3d747970656f6620742e616374696f6e3f742e616374696f6e3a746869732e64656661756c74416374696f6e2c746869732e7461726765743d2266756e6374696f6e223d3d747970656f6620742e7461726765743f742e7461726765743a746869732e64656661756c745461726765742c746869732e746578743d2266756e6374696f6e223d3d747970656f6620742e746578743f742e746578743a746869732e64656661756c74546578742c746869732e636f6e7461696e65723d226f626a656374223d3d3d7028742e636f6e7461696e6572293f742e636f6e7461696e65723a646f63756d656e742e626f64797d7d2c7b6b65793a226c697374656e436c69636b222c76616c75653a66756e6374696f6e2874297b76617220653d746869733b746869732e6c697374656e65723d61282928742c22636c69636b222c2866756e6374696f6e2874297b72657475726e20652e6f6e436c69636b2874297d29297d7d2c7b6b65793a226f6e436c69636b222c76616c75653a66756e6374696f6e2874297b76617220653d742e64656c65676174655461726765747c7c742e63757272656e745461726765743b746869732e636c6970626f617264416374696f6e262628746869732e636c6970626f617264416374696f6e3d6e756c6c292c746869732e636c6970626f617264416374696f6e3d6e65772066287b616374696f6e3a746869732e616374696f6e2865292c7461726765743a746869732e7461726765742865292c746578743a746869732e746578742865292c636f6e7461696e65723a746869732e636f6e7461696e65722c747269676765723a652c656d69747465723a746869737d297d7d2c7b6b65793a2264656661756c74416374696f6e222c76616c75653a66756e6374696f6e2874297b72657475726e20762822616374696f6e222c74297d7d2c7b6b65793a2264656661756c74546172676574222c76616c75653a66756e6374696f6e2874297b76617220653d762822746172676574222c74293b696628652972657475726e20646f63756d656e742e717565727953656c6563746f722865297d7d2c7b6b65793a2264656661756c7454657874222c76616c75653a66756e6374696f6e2874297b72657475726e2076282274657874222c74297d7d2c7b6b65793a2264657374726f79222c76616c75653a66756e6374696f6e28297b746869732e6c697374656e65722e64657374726f7928292c746869732e636c6970626f617264416374696f6e262628746869732e636c6970626f617264416374696f6e2e64657374726f7928292c746869732e636c6970626f617264416374696f6e3d6e756c6c297d7d5d2c5b7b6b65793a226973537570706f72746564222c76616c75653a66756e6374696f6e28297b76617220743d617267756d656e74732e6c656e6774683e302626766f69642030213d3d617267756d656e74735b305d3f617267756d656e74735b305d3a5b22636f7079222c22637574225d2c653d22737472696e67223d3d747970656f6620743f5b745d3a742c6e3d2121646f63756d656e742e7175657279436f6d6d616e64537570706f727465643b72657475726e20652e666f7245616368282866756e6374696f6e2874297b6e3d6e26262121646f63756d656e742e7175657279436f6d6d616e64537570706f727465642874297d29292c6e7d7d5d292c6e7d28722829297d2c3832383a66756e6374696f6e2874297b69662822756e646566696e656422213d747970656f6620456c656d656e74262621456c656d656e742e70726f746f747970652e6d617463686573297b76617220653d456c656d656e742e70726f746f747970653b652e6d6174636865733d652e6d61746368657353656c6563746f727c7c652e6d6f7a4d61746368657353656c6563746f727c7c652e6d734d61746368657353656c6563746f727c7c652e6f4d61746368657353656c6563746f727c7c652e7765626b69744d61746368657353656c6563746f727d742e6578706f7274733d66756e6374696f6e28742c65297b666f72283b74262639213d3d742e6e6f6465547970653b297b6966282266756e6374696f6e223d3d747970656f6620742e6d6174636865732626742e6d6174636865732865292972657475726e20743b743d742e706172656e744e6f64657d7d7d2c3433383a66756e6374696f6e28742c652c6e297b766172206f3d6e28383238293b66756e6374696f6e207228742c652c6e2c6f2c72297b76617220613d692e6170706c7928746869732c617267756d656e7473293b72657475726e20742e6164644576656e744c697374656e6572286e2c612c72292c7b64657374726f793a66756e6374696f6e28297b742e72656d6f76654576656e744c697374656e6572286e2c612c72297d7d7d66756e6374696f6e206928742c652c6e2c72297b72657475726e2066756e6374696f6e286e297b6e2e64656c65676174655461726765743d6f286e2e7461726765742c65292c6e2e64656c65676174655461726765742626722e63616c6c28742c6e297d7d742e6578706f7274733d66756e6374696f6e28742c652c6e2c6f2c69297b72657475726e2266756e6374696f6e223d3d747970656f6620742e6164644576656e744c697374656e65723f722e6170706c79286e756c6c2c617267756d656e7473293a2266756e6374696f6e223d3d747970656f66206e3f722e62696e64286e756c6c2c646f63756d656e74292e6170706c79286e756c6c2c617267756d656e7473293a2822737472696e67223d3d747970656f662074262628743d646f63756d656e742e717565727953656c6563746f72416c6c287429292c41727261792e70726f746f747970652e6d61702e63616c6c28742c2866756e6374696f6e2874297b72657475726e207228742c652c6e2c6f2c69297d2929297d7d2c3837393a66756e6374696f6e28742c65297b652e6e6f64653d66756e6374696f6e2874297b72657475726e20766f69642030213d3d7426267420696e7374616e63656f662048544d4c456c656d656e742626313d3d3d742e6e6f6465547970657d2c652e6e6f64654c6973743d66756e6374696f6e2874297b766172206e3d4f626a6563742e70726f746f747970652e746f537472696e672e63616c6c2874293b72657475726e20766f69642030213d3d74262628225b6f626a656374204e6f64654c6973745d223d3d3d6e7c7c225b6f626a6563742048544d4c436f6c6c656374696f6e5d223d3d3d6e292626226c656e67746822696e2074262628303d3d3d742e6c656e6774687c7c652e6e6f646528745b305d29297d2c652e737472696e673d66756e6374696f6e2874297b72657475726e22737472696e67223d3d747970656f6620747c7c7420696e7374616e63656f6620537472696e677d2c652e666e3d66756e6374696f6e2874297b72657475726e225b6f626a6563742046756e6374696f6e5d223d3d3d4f626a6563742e70726f746f747970652e746f537472696e672e63616c6c2874297d7d2c3337303a66756e6374696f6e28742c652c6e297b766172206f3d6e28383739292c723d6e28343338293b742e6578706f7274733d66756e6374696f6e28742c652c6e297b6966282174262621652626216e297468726f77206e6577204572726f7228224d697373696e6720726571756972656420617267756d656e747322293b696628216f2e737472696e67286529297468726f77206e657720547970654572726f7228225365636f6e6420617267756d656e74206d757374206265206120537472696e6722293b696628216f2e666e286e29297468726f77206e657720547970654572726f722822546869726420617267756d656e74206d75737420626520612046756e6374696f6e22293b6966286f2e6e6f64652874292972657475726e2066756e6374696f6e28742c652c6e297b72657475726e20742e6164644576656e744c697374656e657228652c6e292c7b64657374726f793a66756e6374696f6e28297b742e72656d6f76654576656e744c697374656e657228652c6e297d7d7d28742c652c6e293b6966286f2e6e6f64654c6973742874292972657475726e2066756e6374696f6e28742c652c6e297b72657475726e2041727261792e70726f746f747970652e666f72456163682e63616c6c28742c2866756e6374696f6e2874297b742e6164644576656e744c697374656e657228652c6e297d29292c7b64657374726f793a66756e6374696f6e28297b41727261792e70726f746f747970652e666f72456163682e63616c6c28742c2866756e6374696f6e2874297b742e72656d6f76654576656e744c697374656e657228652c6e297d29297d7d7d28742c652c6e293b6966286f2e737472696e672874292972657475726e2066756e6374696f6e28742c652c6e297b72657475726e207228646f63756d656e742e626f64792c742c652c6e297d28742c652c6e293b7468726f77206e657720547970654572726f722822466972737420617267756d656e74206d757374206265206120537472696e672c2048544d4c456c656d656e742c2048544d4c436f6c6c656374696f6e2c206f72204e6f64654c69737422297d7d2c3831373a66756e6374696f6e2874297b742e6578706f7274733d66756e6374696f6e2874297b76617220653b6966282253454c454354223d3d3d742e6e6f64654e616d6529742e666f63757328292c653d742e76616c75653b656c73652069662822494e505554223d3d3d742e6e6f64654e616d657c7c225445585441524541223d3d3d742e6e6f64654e616d65297b766172206e3d742e6861734174747269627574652822726561646f6e6c7922293b6e7c7c742e7365744174747269627574652822726561646f6e6c79222c2222292c742e73656c65637428292c742e73657453656c656374696f6e52616e676528302c742e76616c75652e6c656e677468292c6e7c7c742e72656d6f76654174747269627574652822726561646f6e6c7922292c653d742e76616c75657d656c73657b742e6861734174747269627574652822636f6e74656e746564697461626c6522292626742e666f63757328293b766172206f3d77696e646f772e67657453656c656374696f6e28292c723d646f63756d656e742e63726561746552616e676528293b722e73656c6563744e6f6465436f6e74656e74732874292c6f2e72656d6f7665416c6c52616e67657328292c6f2e61646452616e67652872292c653d6f2e746f537472696e6728297d72657475726e20657d7d2c3237393a66756e6374696f6e2874297b66756e6374696f6e206528297b7d652e70726f746f747970653d7b6f6e3a66756e6374696f6e28742c652c6e297b766172206f3d746869732e657c7c28746869732e653d7b7d293b72657475726e286f5b745d7c7c286f5b745d3d5b5d29292e70757368287b666e3a652c6374783a6e7d292c746869737d2c6f6e63653a66756e6374696f6e28742c652c6e297b766172206f3d746869733b66756e6374696f6e207228297b6f2e6f666628742c72292c652e6170706c79286e2c617267756d656e7473297d72657475726e20722e5f3d652c746869732e6f6e28742c722c6e297d2c656d69743a66756e6374696f6e2874297b666f722876617220653d5b5d2e736c6963652e63616c6c28617267756d656e74732c31292c6e3d2828746869732e657c7c28746869732e653d7b7d29295b745d7c7c5b5d292e736c69636528292c6f3d302c723d6e2e6c656e6774683b6f3c723b6f2b2b296e5b6f5d2e666e2e6170706c79286e5b6f5d2e6374782c65293b72657475726e20746869737d2c6f66663a66756e6374696f6e28742c65297b766172206e3d746869732e657c7c28746869732e653d7b7d292c6f3d6e5b745d2c723d5b5d3b6966286f26266529666f722876617220693d302c613d6f2e6c656e6774683b693c613b692b2b296f5b695d2e666e213d3d6526266f5b695d2e666e2e5f213d3d652626722e70757368286f5b695d293b72657475726e20722e6c656e6774683f6e5b745d3d723a64656c657465206e5b745d2c746869737d7d2c742e6578706f7274733d652c742e6578706f7274732e54696e79456d69747465723d657d7d2c653d7b7d3b66756e6374696f6e206e286f297b696628655b6f5d2972657475726e20655b6f5d2e6578706f7274733b76617220723d655b6f5d3d7b6578706f7274733a7b7d7d3b72657475726e20745b6f5d28722c722e6578706f7274732c6e292c722e6578706f7274737d72657475726e206e2e6e3d66756e6374696f6e2874297b76617220653d742626742e5f5f65734d6f64756c653f66756e6374696f6e28297b72657475726e20742e64656661756c747d3a66756e6374696f6e28297b72657475726e20747d3b72657475726e206e2e6428652c7b613a657d292c657d2c6e2e643d66756e6374696f6e28742c65297b666f7228766172206f20696e2065296e2e6f28652c6f292626216e2e6f28742c6f2926264f626a6563742e646566696e6550726f706572747928742c6f2c7b656e756d657261626c653a21302c6765743a655b6f5d7d297d2c6e2e6f3d66756e6374696f6e28742c65297b72657475726e204f626a6563742e70726f746f747970652e6861734f776e50726f70657274792e63616c6c28742c65297d2c6e28313334297d28292e64656661756c747d7d5d293b`
)

func AddRouterOfChunkAdb9e944Fa692c2aJs(router *gin.Engine) {
    
    utils.GetJs(router, CHUNK_ADB9E944_FA692C2A_JS_HEX_CONTENT, CHUNK_ADB9E944_FA692C2A_JS_RELATIVE_PATH)
    
}







