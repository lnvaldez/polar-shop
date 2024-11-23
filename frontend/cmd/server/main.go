package main  

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"html/template"
	"net/http"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type page struct {
	Header string 
	Paragraph string 
}

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
	err = coll.FindOne(context.TODO(), bson.D{{Key: "name", Value: name}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("Product not found with name %s\n", name)
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := page{Header: "Wallace", Paragraph: "Ninja"}
		fmt.Println(p)
		templatePath := "C:/Users/lnval/Code/act/me/polar-shop/frontend/template.html"
		t, err := template.ParseFiles(templatePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}