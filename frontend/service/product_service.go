package service

import (
	"net/http"
	"encoding/json"
	"log"
	"polar-shop/db"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svc *ProductService) GetAvailableProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	db := db.ProductDB{MongoCollection: svc.MongoCollection}

	prods, err := db.GetAvailableProducts()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error; ", err)
		res.Error = err.Error()
		return
	}

	res.Data = prods 
	w.WriteHeader(http.StatusOK)
}