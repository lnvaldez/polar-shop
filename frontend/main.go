package main

import (
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var mongoClient *mongo.Client

func init() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	
}

func main() {
	
}