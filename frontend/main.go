package main

import (
	"os"
	"context"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func init() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
}

func main() {
	
}