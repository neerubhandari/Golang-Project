// package main

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type Product struct{
//    ID string `json:"id"`
//    Name string `json:"name"`
//    Price int `json:"price"`
//    Imahe string `json:"image"`
// }

// var product [] Product
// func init(){
//    product =make ([]Product ,0)
//    file,_:=ioutil.ReadFile("back.json")
//    _=json.Unmarshal([]byte(file),&product)
// }

// func getProducts(c *gin.Context) {
//    c.JSON(http.StatusOK, product)
// }
// func main() {
//    router := gin.Default()
//    router.GET("/products", getProducts)
//    router.Run()}

package main

import (
	"context"
	handlers "ecomm-back/handlers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var productsHandler *handlers.ProductsHandler

func init() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI("mongodb://admin:password@localhost:27017/admin?authSource=admin"))
	if err = client.Ping(context.TODO(),
		readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	collection := client.Database(
		"MONGO_DATABASE").Collection("products")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	status := redisClient.Ping()
	fmt.Println(status)
	productsHandler = handlers.NewRecipesHandler(ctx,
		collection, redisClient)
}

func main() {
	router := gin.Default()
	authorized := router.Group("/")
	authorized.Use(handlers.AuthMiddleware())
	{
		authorized.POST("/products", productsHandler.PostProducts)
		authorized.GET("/products", productsHandler.ListRecipesHandler)
		authorized.PUT("/products/:id", productsHandler.UpdateProductHandler)

	}
	router.Run()

}
