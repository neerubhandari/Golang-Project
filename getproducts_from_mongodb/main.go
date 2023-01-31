package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

	func getProducts(c *gin.Context) {
		var product  Product
	collection := client.Database((
"MONGO_DATABASE")).Collection("products")
		cur, err := collection.Find(ctx, bson.M{})
		if err != nil {
		c.JSON(http.StatusInternalServerError,
		gin.H{"error": err.Error()})
		return
		}
		
		//The defer keyword instructs a function to execute after the surrounding function completes.
        //This method frees the resources your cur consumes in both the client application and the MongoDB server.
		defer cur.Close(ctx)
		

		products := make([]Product, 0)
		for cur.Next(ctx) {
		cur.Decode(&product)
		products = append(products, product)
		}
		c.JSON(http.StatusOK, products)
	}

	func main() {
		router := gin.Default()
		router.GET("/products", getProducts)
		router.Run()
	}