package handlers

import (
	"encoding/json"
	"polar-shop/db"
	"polar-shop/models"
	// "polar-shop/service"

	// "html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrderService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data interface{} `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func (svc *OrderService) CreateNewOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}

	defer json.NewEncoder(w).Encode(res)

	var ord model.Order

	err := json.NewDecoder(r.Body).Decode(&ord)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid body", err)
		res.Error = err.Error()
		return
	}

	db := db.OrderDB{MongoCollection: svc.MongoCollection}

	createOrd, err := db.CreateNewOrder(&ord)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("order creation error ", err)
		res.Error = err.Error() 
		return 
	}

	res.Data = ord.ID
	w.WriteHeader(http.StatusOK)

	log.Println("order created ", createOrd)
}