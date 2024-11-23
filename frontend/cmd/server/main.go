package main  

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"html/template"
	"net/http"

	"frontend/internal/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type page struct {
	Header string 
	Paragraph string 
}

func main() {
	cfg := config.Load()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

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