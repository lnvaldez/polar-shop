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
	"github.com/rs/cors"
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

    tmpl := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/product_list.html",
		"templates/register.html",
		"templates/login.html",
		"templates/partials/footer.html",
		"templates/partials/navbar.html", 
	))

    productColl := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("PRODUCTS_COLLECTION"))
    orderColl := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("ORDERS_COLLECTION"))


    productService := &service.ProductService{
        MongoCollection: productColl,
    }

	orderService := &handlers.OrderService{
		MongoCollection: orderColl,
	}

    r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    r.HandleFunc("/products", handlers.ProductListHandler(productService, tmpl)).Methods("GET")
	r.HandleFunc("/register", handlers.RenderRegisterPage(tmpl)).Methods("GET")
	r.HandleFunc("/login", handlers.RenderLoginPage(tmpl)).Methods("GET")

	r.HandleFunc("/order", orderService.CreateNewOrder).Methods(http.MethodPost)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5000"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	corsHandler := c.Handler(r)

	log.Println("⭐ Server running on port 5001")
	http.ListenAndServe(":5001", corsHandler)
}

