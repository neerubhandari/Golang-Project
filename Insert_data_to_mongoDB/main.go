package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)
var ctx context.Context
var err error
var client *mongo.Client


type Product struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Image string `json:"image"`
 }
var products [] Product


func init() {
	ctx = context.Background()
	client, err = mongo.Connect(ctx,
	options.Client().ApplyURI("mongodb://admin:password@localhost:27017/admin?authSource=admin"))
	if err = client.Ping(context.TODO(),
	readpref.Primary()); err != nil {log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	}

	func postProducts(c *gin.Context) {
		var product  Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":
			err.Error()})
			return
			}
		collection := client.Database((
	"MONGO_DATABASE")).Collection("products")
			_, err = collection.InsertOne(ctx, product)
			if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Error while inserting a new products"})
			return
			}

		c.JSON(http.StatusOK, product)
	 }

	func main() {
		router := gin.Default()
		router.POST("/products", postProducts)
		router.Run()}