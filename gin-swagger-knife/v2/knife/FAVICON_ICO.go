package knife


import (
	"gin-swagger-knife/constant"
	"gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)

func AddRouterOfFaviconIco(router *gin.Engine) {
    
	utils.GetOther(router, FAVICON_ICO_BASE64_OR_CONTENT, FAVICON_ICO_RELATIVE_PATH)
	
}

const (
	FAVICON_ICO_RELATIVE_PATH = constant.ROOT_PATH + "/favicon.ico}"
	FAVICON_ICO_BASE64_OR_CONTENT = `AAABAAEAICAAAAEAIACoEAAAFgAAACgAAAAgAAAAQAAAAAEAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAzakCwN3PKMDl30jA4thOwOPbPr/XwBIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMHp6wDA6OYmwOjnksDo6OjB6uv/werr/8Hq6//B6uv/wOnq/MDn5cbA491evsSbBAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADA6usOwerrksHp6/jB6uv/werr/8Hq6//A5N//v9S8/77Lp/++y6n/v9jD/8Dn5f/A6ObOwN/TLgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwerrJsHq69jB6uv/werr/8Dp6v++yKH/u5xA/7mCB/+4gAD/uYAA/7mAAP+4gAD/uYUM/7uhS/++y6jwwOTXSAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMHp6ibA6uvmwerr/8Hp6v+/3M3/upg4/7iAAP+5gQD/uYEA/7mBAP+5gQD/uYEA/7mBAP+5gQD/uYEA/7iAAP+6mDf0vKZYOgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADA6uoQwOnq2MHq6//B6er/v9O5/7mGEf+4gAL/uIAC/7iAAv+4gAH/uYEA/7mBAP+4gAH/uIAC/7iAAv+4gAL/uIAB/7mBAP+5gADguYAADgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAv+rqAMDq6pLB6uv/werr/7/by/+5hxL/qHUP/ygcd/8LCIr/DgmJ/2pKQv+5gQD/uYEA/4hfK/8RDIf/CwiK/xYPhP+CWzD/uYEA/7h/AP+8ql6OAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADA6eoiwOnq+MHq6//B6uv/upg3/7mBAP+5gQD/l2kc/wAAkf86KWf/uYEA/7mBAP+5gQD/r3oJ/woHiv8DAo//kmYi/7mBAP+5gwf/vLFw/7/VvvS/3s8yAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMDq6ojB6uv/werr/77Lp/+4gAD/uYEA/7mBAP+mdA//AACS/0kzWv+5gQD/uYEA/7mBAP9bQE3/AACR/1c8Uf+5gQD/uYEA/7mAAP+5hhD/vLBt/7iAAnIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADC6+0Awerq2MHq6//B6uv/u6FM/7mBAP+5gQD/uYEA/6h1Dv8AAJP/TDVY/7mBAP+5gQD/oHAV/wYEjf8XEIH/snwH/7mBAP+5gQD/uYEA/7iAAP+5hg3/uYAAwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMHo6RTA6uv/werr/8Do6P+5hg//uYEA/7mBAP+5gQD/qHUO/wAAk/9MNVj/uYEA/7mBAP84J2j/AACR/4JaLv+4gAD/uYEC/7qXNP+9tXj/vsii/77Psv++xp/0wOjoTMDq6gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwerrOsHq6//B6uv/wN3P/7iAAP+5gQD/uYEA/7mBAP+odQ7/AACT/0w1WP+4gQD/hFwt/wAAkf87KWb/uYIF/7unWP+/18P/werr/8Hq6v/B6uv/werr/8Hq6//B6uv/wenrfAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADB6utGwerr/8Hq6/+/1sL/uIAA/7mBAP+5gQD/uYEA/6h1Dv8AAJP/SjRa/6JxFf8VD4L/CAWL/6WJSf+/1sH/werr/8Hq6//B6uv/wers/8Hq7P/B6ur/werr/8Hq6//A6er4wOvrFgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMDq60bB6uv/werr/8Dg1f+5gQL/uYEA/7mBAP+5gQD/qHUO/wAAk/8SDYT/Ew2F/wUGkv95ibL/wOnp/8Hq6//A6er/vsii/7udQv+5hAr/uYYP/72/j//B6uv/werr/8Hq6//B6utCAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAwerrKsHq6//B6uv/werr/7uiTv+5gQD/uYEA/7mBAP+odQ7/AACT/0w1WP+2lkL/pMTY/8Hq6//B6uv/v9fD/7qXNv+4gAD/uYEA/7mBAP+5gQD/uYIG/8Dg1//B6uv/werr/8Dq60QAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAC/6eoCwerr5MHq6//B6uv/wOTf/7umWP+5ggf/uH8A/6l5F/86Rq3/lrHK/8Hp6v/B6uv/v+fq/5mgnv+5gwf/uYEA/7mBAP+5gQD/uYEA/7mBAP+4gAD/v9O7/8Hq6//A6er/werrHAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADA6etywerr/8Hq6//B6uv/werr/8Dm4v/A4df/wOfl/8Hq6//B6uv/werr/7/Xwv+liUn/EAuH/5BlI/+4gAD/uYEA/7mBAP+5gQD/uYEA/7mKGP/A5+T/werr/8Hp67gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAL7OsQS+yqbewerq/8Hq6//B6uv/werr/8Hq6//B6uv/werr/7PNyf+7plf/uYEC/7iBAP91UTn/Ew2E/6h1D/+5gQD/uo8i/7qSKP+7mj3/v9W//8Hq6v/A6eviwOnqIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALiAAHq6mTr/vb2K/77NrP++z7D/vsSa/6ehe/8hKJ//QS5h/7mBAP+5gQD/uYAA/7Z/A/8NCYj/KRx0/7N9Bf+9wZL/wers/8Hq6//A6uv6wenqpsHq6xwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuYEAEriAAOi5gQD/uYAA/7mAAP+teAv/NyZs/wAAk/8KB43/fVc0/7mCBf+5jBv/flgz/wMCkf8AAJP/Uk+H/8Hn5P+/1b//u6dZ/7qVMo4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuIEAYLiBAP+5gQD/uYEA/7V+BP+hcBb/oXAW/6FwFv+odQ//uYAA/728h/+spX7/pJNm/6WqmP+74uf/werq/7ufR/+5gQDguYEADgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAC7gAAAuIAAmrmBAP+5gQD/uYEA/7mBAP+5gQD/uYEA/7mBAP+5gQD/uYgS/7/SuP/B6uv/werr/8Dm4v+8rWb/uIAA8riBADgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAC4gAAEuYEAoLiAAP+5gQD/uYEA/7mBAP+5gQD/uYEA/7mBAP+5gQD/uYEC/7qWMv+7nD7/uYYQ/7iAAPC4gABEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADDlAACuIAAbrmBAPK5gQD/uYEA/7mBAP+5gQD/uYEA/7mBAP+5gQD/uYEA/7mBAP+4gADKuYAAKgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuIEAHLmBAIq4gQDiuYEA/7mBAP+5gQD/uYEA/7mBAPq4gADAuIAAWriCAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAALmBAAK5gAAiuYEARLiAAEy5gQA6uYAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA///////////////////////gH///gAf//wAD//4AAf/8AAD/+AAAf/gAAH/wAAB/8AAAP/AAAD/wAAAf8AAAD/AAAA/wAAAP8AAAD/gAAA/4AAAf/AAAP/wAAH/+AAD//gAB//8AA///wAf//+Af//////////////////////8=`
)






