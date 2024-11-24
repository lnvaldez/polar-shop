package service

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svc *ProductService) GetAvailableProducts(w http.ResponseWriter, r *http.Request) {}