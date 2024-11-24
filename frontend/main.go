package main

import (
	"context"
	"log"
	"net/http"
	"os"

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

	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler).Methods(http.MethodGet)

	log.Println("⭐ Server running on port 5001")
	http.ListenAndServe(":5001", r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running..."))
}