# knife4go

#### 项目调整中... （今年加班太多，维护缓慢，大家谅解）
- 1. 将openapi3.0和swagger2.0的支持单独列分支,master将保持swagger2.0支持。解决现存的。
- 2. 解决现存openapi3.0支持bug。
- 3. 拉新分支用https://github.com/wI2L/fizz 来适配knife4j的ui,不用再生成swagger.json。





#### 介绍
优化gin-swagger的界面,让knif4j适配go语言。感谢萧明。
knif4j地址：https://gitee.com/xiaoym/knife4j


#### 效果图
![image](https://gitee.com/youbeiwuhuan/knife4go/raw/master/img/knife4go.png)


#### 模块说明
| 模块                         | 描述                                                                                                                                                                                                                |  
|----------------------------|:------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------| 
| gin-swagger-knife          | 替换gin-swagger的界面实现(并非在原项目代码上修改)，供gin项目引用 (适用于swagger2)                                                                                                                                                            |  
| gin-openapi3-knife         | 替换gin-swagger的界面实现(并非在原项目代码上修改)，供gin项目引用 (适用于openapi3)                                                                                                                                                            |  
| fizz-knife                 | 替换fizz的界面实现(并非在原项目代码上修改)，供gin项目引用 (基于fizz展示openapi3,不需要用swagger.json) ,使用样例参考源码下的example,还可参考：[wI2L/fizz: Gin wrapper with OpenAPI 3 spec generation (github.com)](https://github.com/wI2L/fizz) 项目说明，暂时没时间整理详细文档 |  
| knife-vue-dist             | knife4j-vue项目打包后的dist目录内容，用于knife-vue-2-go-code中生成代码的逻辑                                                                                                                                                           |    
| knife-vue-2-go-code        | 用于生成gin-swagger-knife中的前端文件接口代码（knife目录下代码），方便根据最新版本的knif4j升级界面                                                                                                                                                   |      
| knife-openapi3-go-code     | 用于生成gin-openapi3-knife中的前端文件接口代码（knife目录下代码），方便根据最新版本的knif4j升级界面                                                                                                                                                  |      
| gin-swagger-knife-example  | gin-swagger-knife 使用示例                                                                                                                                                                                            |
| gin-openapi3-knife-example | gin-openapi3-knife 使用示例                                                                                                                                                                                           |
| fizz-knife-example         | fizz-knife 使用示例                                                                                                                                                                                                   |


#### gin-swagger-knife  使用步骤 （gin-openapi3-knife 与之相同，只换依赖包即可）
- 1. 按照传统gin-swagger方式生成swagger.json,将内容拷贝到一个字符串中。关于gin-swagger的使用参考[gin-swagger生成API文档](https://www.cnblogs.com/zhzhlong/p/11800787.html) 或  [swaggo项目](https://github.com/swaggo/swag)
	- 1.1 先下载cmd包，才能执行相关命令
		go get -u github.com/swaggo/swag/cmd/swag
		我开始没成功，后来进入$GOPATH/bin/ 目录执行go get github.com/swaggo/swag/cmd/swag ，在bin目录下生成一个swag.exe文件，把$GOPATH/bin/ 添加到Path环境变量才算成功
	- 1.2 执行初始化命令（相关的swagger注释写法参照gin-swagger官网）
		swag init  // 注意，一定要和main.go处于同一级目录
		初始化命令，在根目录生成一个docs文件夹，swagger.json就在该文件夹中
			docs/swagger.json
	- 1.3 将swagger.json的内容拷贝到下面示例代码的swaggerJson变量中
	
	
	
- 2. 引入gin-swagger-knife项目
	go get gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife
	
	
- 3.  gin 中初始化相关访问路径
	![image](https://gitee.com/youbeiwuhuan/knife4go/raw/master/img/example.png)

	``` go
		
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




- 4. 访问
	http://127.0.0.1:8080/knife/doc.html



#### fizz-knife （不用再生成swagger.json，用硬编码替代swagger注解）

- 1. 效果图
![image](https://gitee.com/youbeiwuhuan/knife4go/raw/master/img/fizz-knife.png)	


- 2. 引入gin-swagger-knife项目
	 go get gitee.com/youbeiwuhuan/knife4go/fizz-knife

- 3. 代码样例


``` go
    
    package main

	import (
		"fmt"
		fizz_knife "gitee.com/youbeiwuhuan/knife4go/fizz-knife"
		"github.com/gin-contrib/cors"
		"github.com/gin-gonic/gin"
		"github.com/juju/errors"
		"github.com/loopfz/gadgeto/tonic"
		"github.com/wI2L/fizz"
		"github.com/wI2L/fizz/openapi"
		"log"
		"net/http"
		"strconv"
		"time"
	)
	
	// Fruit represents a sweet, fresh fruit.
	type Fruit struct {
		Name    string    `json:"name" validate:"required" description:"名称" example:"banana"`
		Origin  string    `json:"origin" validate:"required" description:"水果的出产国家" enum:"ecuador,france,senegal,china,spain"`
		Price   float64   `json:"price" validate:"required" description:"价格" example:"5.13"`
		AddedAt time.Time `json:"-" binding:"-" description:"Date of addition of the fruit to the market"`
	}
	
	// TypeName implements openapi.Typer interface for Fruit.
	func (f *Fruit) TypeName() string { return "RottenFruit" }
	
	func main() {
		router, err := NewRouter()
		if err != nil {
			log.Fatal(err)
		}
		srv := &http.Server{
			Addr:    ":8080",
			Handler: router,
		}
		srv.ListenAndServe()
	}
	
	// NewRouter returns a new router for the
	// Pet Store.
	func NewRouter() (*fizz.Fizz, error) {
		engine := gin.New()
		engine.Use(cors.Default())
	
		// Override type names.
		// fizz.Generator().OverrideTypeName(reflect.TypeOf(Fruit{}), "SweetFruit")
	
		// Initialize the informations of
		// the API that will be served with
		// the specification.
		infos := &openapi.Info{
			Title:       "Fruits Market水果超市",
			Description: `This is a sample Fruits market server.  水果超市接口`,
			Version:     "1.0.0",
		}
	
		fizz := fizz_knife.InitSwaggerKnife(engine, infos)
	
		// Setup routes.
		routes(fizz.Group("/market", "超市接口", "Your daily dose of freshness，你的描述"))
	
		if len(fizz.Errors()) != 0 {
			return nil, fmt.Errorf("fizz errors: %v", fizz.Errors())
		}
		return fizz, nil
	}
	
	func routes(grp *fizz.RouterGroup) {
		// Add a new fruit to the market.
		grp.POST("", []fizz.OperationOption{
			fizz.Summary("给市场添加一个水果"),
			fizz.Response("400", "Bad request", nil, nil,
				map[string]interface{}{"error": "fruit already exists"},
			),
		}, tonic.Handler(CreateFruit, 200))
	
		// Remove a fruit from the market,
		// probably because it rotted.
		grp.DELETE("/:name", []fizz.OperationOption{
			fizz.Summary("从市场删除一个水果"),
			fizz.ResponseWithExamples("400", "Bad request", nil, nil, map[string]interface{}{
				"fruitNotFound": map[string]interface{}{"error": "fruit not found"},
				"invalidApiKey": map[string]interface{}{"error": "invalid api key"},
			}),
		}, tonic.Handler(DeleteFruit, 204))
	
		// List all available fruits.
		grp.GET("", []fizz.OperationOption{
			fizz.Summary("水果列表"),
			fizz.Response("400", "Bad request", nil, nil, nil),
			fizz.Header("X-Market-Listing-Size", "Listing size", fizz.Long),
		}, tonic.Handler(ListFruits, 200))
	}
	
	// FruitIdentityParams represents the parameters that
	// are required to identity a unique fruit in the market.
	type FruitIdentityParams struct {
		Name   string `path:"name"`
		APIKey string `header:"X-Api-Key" validate:"required"`
	}
	
	// ListFruitsParams represents the parameters that can
	// be used to filter the fruit's market listing.
	type ListFruitsParams struct {
		Origin   *string  `query:"origin" description:"filter by fruit origin"`
		PriceMin *float64 `query:"price_min" description:"filter by minimum inclusive price" validate:"omitempty,min=1"`
		PriceMax *float64 `query:"price_max" description:"filter by maximum inclusive price" validate:"omitempty,max=15"`
	}
	
	// CreateFruit add a new fruit to the market.
	func CreateFruit(c *gin.Context, fruit *Fruit) (*Fruit, error) {
	
		return fruit, nil
	}
	
	// DeleteFruit removes a fruit from the market.
	func DeleteFruit(c *gin.Context, params *FruitIdentityParams) error {
		if params.APIKey == "" {
			return errors.Forbiddenf("invalid api key")
		}
	
		return nil
	}
	
	// ListFruits lists the fruits of the market.
	// Parameters can be used to filter the fruits.
	func ListFruits(c *gin.Context, params *ListFruitsParams) ([]*Fruit, error) {
		basket := make([]*Fruit, 0)
	
		c.Header("X-Market-Listing-Size", strconv.Itoa(len(basket)))
	
		return basket, nil
	}
```


