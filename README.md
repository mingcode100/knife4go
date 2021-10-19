# knife4go

#### 介绍
优化gin-swagger的界面,让knif4j适配go语言。感谢萧明。
knif4j地址：https://gitee.com/xiaoym/knife4j


#### 效果图
![image](https://gitee.com/youbeiwuhuan/knife4go/raw/master/img/knife4go.png)


#### 模块说明
| 模块        | 描述    |  
| --------   | :-----   | 
|  gin-swagger-knife   |  替换gin-swagger的界面实现(并非在原项目代码上修改)，供gin项目引用  |      
|  knife-vue-2-go-code   |  用于生成gin-swagger-knife中的代码，方便根据最新版本的knif4j升级界面  |      
|  knife4go-example |  gin-swagger-knife 使用示例  |


#### 使用步骤
1. 按照传统gin-swagger方式生成swagger.json,将内容拷贝到一个字符串中。
	1.1 先下载cmd包，才能执行相关命令
		go get -u github.com/swaggo/swag/cmd/swag
		我开始没成功，后来进入$GOPATH/bin/ 目录执行go get github.com/swaggo/swag/cmd/swag ，在bin目录下生成一个swag.exe文件，把$GOPATH/bin/ 添加到Path环境变量才算成功
	1.2 执行初始化命令（相关的swagger注释写法参照gin-swagger官网）
		swag init  // 注意，一定要和main.go处于同一级目录
		初始化命令，在根目录生成一个docs文件夹，swagger.json就在该文件夹中
			docs/swagger.json
	1.3 将swagger.json的内容拷贝到下面示例代码的swaggerJson变量中
	
	
	
2. 引入gin-swagger-knife项目
	go get gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife
	
	
3.  gin 中初始化相关访问路径
··· go
	
	import (
		"fmt"
		gin_swagger_knife "gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife"
		"github.com/gin-gonic/gin"
	)
	
	router := gin.Default()
	// 获取 swag int 命令生成的swagger.json文件里的内容
	swaggerJson := "swagger.json的内容"
	//初始化gin路径
	gin_swagger_knife.InitSwaggerKnife(router,swaggerJson)
	router.Run(":8080")
```

![image](https://gitee.com/youbeiwuhuan/knife4go/raw/master/img/example.png)


4. 访问
	http://127.0.0.1:8080/knife/doc.html
	
	