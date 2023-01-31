package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

	type Product struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
    Name        string              `json:"name" bson:"name"`
	Price       int                 `josn:"price" bson:"price"`
	Image      string               `json:"image" bson:"image"`
	}