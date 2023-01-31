package handlers

import (
	"ecomm-back/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)
type ProductsHandler struct {
collection *mongo.Collection
ctx        context.Context
}

func NewRecipesHandler(ctx context.Context, collection *mongo.
Collection) *ProductsHandler {
return &ProductsHandler{
collection: collection,
ctx: ctx,
}
}

func (handler *ProductsHandler) ListRecipesHandler(c *gin.
	Context) {
	cur, err := handler.collection.Find(handler.ctx, bson.M{})
	if err != nil {
	c.JSON(http.StatusInternalServerError,
	gin.H{"error": err.Error()})
	return
	}
	defer cur.Close(handler.ctx)
	products := make([]models.Product, 0)
	for cur.Next(handler.ctx) {
	var product models.Product
	cur.Decode(&product)
	products = append(products, product)
	}
	c.JSON(http.StatusOK, products)
	}


	func (handler *ProductsHandler) PostProducts(c *gin.Context) {
		var product  models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":
			err.Error()})
			return
			}
			_, err := handler.collection.InsertOne(handler.ctx, product)
			if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Error while inserting a new products"})
			return
			}

		c.JSON(http.StatusOK, product)
	 }

	 func (handler *ProductsHandler) UpdateProductHandler(c *gin.Context) {
		id :=c.Param("id")
		var product  models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":
			err.Error()})
			return
			}
			objectId, _ := primitive.ObjectIDFromHex(id)
		
	_, err := handler.collection.UpdateOne(handler.ctx, bson.M{
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