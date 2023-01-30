package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)
var ctx context.Context
var err error
var client *mongo.Client
func main() {

ctx = context.Background()
client, err = mongo.Connect(ctx,
options.Client().ApplyURI("mongodb://admin:password@localhost:27017/admin?authSource=admin"))
if err = client.Ping(context.TODO(),
readpref.Primary()); err != nil {
log.Fatal(err)
}
log.Println("Connected to MongoDB")
}

