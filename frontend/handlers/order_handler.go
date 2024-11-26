package handlers

import (
	"polar-shop/models"
	"polar-shop/service"

	"html/template"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data interface{} `json:"data,omitempty"`
	Error string `json:"error, omitempty"`
}

func (svc *service.OrderService) 