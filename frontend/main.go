package main  

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found");
	}

	uri := os.Getenv("MONGO_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect((context.TODO())); err != nil {
			panic(err)
		}
	}()
	
	coll := client.Database("polar_shop").Collection("products")
	name := "Bonnet"

	var result bson.M 
	err = coll.FindOne(context.TODO(), bson.D{{"name", name}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("Product not found with name $s\n", name)
		return
	}
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}