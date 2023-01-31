package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)
var ctx context.Context
var err error
var client *mongo.Client


type Product struct{
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

	func updateProductHandler(c *gin.Context) {
		id :=c.Param("id")
		var product  Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":
			err.Error()})
			return
			}
			objectId, _ := primitive.ObjectIDFromHex(id)
		collection := client.Database((
	"MONGO_DATABASE")).Collection("products")
	_, err = collection.UpdateOne(ctx, bson.M{
		"_id": objectId,
		}, bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: product.Name},
		{Key: "price", Value: product.Price},
		{Key: "image", Value: product.Image},
		}}})
			if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Error while inserting a new products"})
			return
			}

			c.JSON(http.StatusOK, gin.H{"message": "product has been updated"})
	 }

	func main() {
		router := gin.Default()
		router.PUT("/products/:id", updateProductHandler)
		router.Run()}