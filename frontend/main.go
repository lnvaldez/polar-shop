package main

import (
	"context"
	"log"
	"net/http"
	"html/template"
	"os"
	"polar-shop/service"
	"polar-shop/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		panic(err)
	}

	err = mongoClient.Ping(context.Background(), readpref.Primary())

	if err != nil {
		panic(err)
	}

	log.Println("✅ Connected to MongoDB")
}

func main() {
	defer mongoClient.Disconnect(context.Background())

    tmpl := template.Must(template.ParseGlob("templates/*.html"))

    coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("PRODUCTS_COLLECTION"))

    productService := &service.ProductService{
        MongoCollection: coll,
    }

    r := mux.NewRouter()

    r.HandleFunc("/products", handlers.ProductListHandler(productService, tmpl)).Methods("GET")
	r.HandleFunc("/register", handlers.RenderRegisterPage(tmpl)).Methods("GET")
	r.HandleFunc("/login", handlers.RenderLoginPage(tmpl)).Methods("GET")

	log.Println("⭐ Server running on port 5001")
	http.ListenAndServe(":5001", r)
}

