package main

import (
	"fmt"
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
	Name    string    `json:"name" validate:"required" example:"banana"`
	Origin  string    `json:"origin" validate:"required" description:"Country of origin of the fruit" enum:"ecuador,france,senegal,china,spain"`
	Price   float64   `json:"price" validate:"required" description:"Price in euros" example:"5.13"`
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

	fizz := fizz.NewFromEngine(engine)

	// Override type names.
	// fizz.Generator().OverrideTypeName(reflect.TypeOf(Fruit{}), "SweetFruit")

	// Initialize the informations of
	// the API that will be served with
	// the specification.
	infos := &openapi.Info{
		Title:       "Fruits Market",
		Description: `This is a sample Fruits market server.`,
		Version:     "1.0.0",
	}

	// 这里  fizz.OpenAPI(infos, "json") 就是open-api-v3.json 内容。正好可以结合knif4j的open-api-v3界面来适配
	// Create a new route that serve the OpenAPI spec.
	fizz.GET("/openapi.json", nil, fizz.OpenAPI(infos, "json"))

	// Setup routes.
	routes(fizz.Group("/market", "market", "Your daily dose of freshness"))

	if len(fizz.Errors()) != 0 {
		return nil, fmt.Errorf("fizz errors: %v", fizz.Errors())
	}
	return fizz, nil
}

func routes(grp *fizz.RouterGroup) {
	// Add a new fruit to the market.
	grp.POST("", []fizz.OperationOption{
		fizz.Summary("Add a fruit to the market"),
		fizz.Response("400", "Bad request", nil, nil,
			map[string]interface{}{"error": "fruit already exists"},
		),
	}, tonic.Handler(CreateFruit, 200))

	// Remove a fruit from the market,
	// probably because it rotted.
	grp.DELETE("/:name", []fizz.OperationOption{
		fizz.Summary("Remove a fruit from the market"),
		fizz.ResponseWithExamples("400", "Bad request", nil, nil, map[string]interface{}{
			"fruitNotFound": map[string]interface{}{"error": "fruit not found"},
			"invalidApiKey": map[string]interface{}{"error": "invalid api key"},
		}),
	}, tonic.Handler(DeleteFruit, 204))

	// List all available fruits.
	grp.GET("", []fizz.OperationOption{
		fizz.Summary("List the fruits of the market"),
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
