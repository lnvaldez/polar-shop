package handlers

import (
	"encoding/json"
	"polar-shop/db"
	"polar-shop/models"
	// "polar-shop/service"

	// "html/template"
	"strconv"
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

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error parsing form data", err)
		res.Error = "Error parsing form data"
		return
	}
	ord := model.Order{
		CustomerName:    r.FormValue("customerName"),
		Email:           r.FormValue("email"),
		Address:         r.FormValue("address"),
		Product:         r.FormValue("productName"), 
		Status:          "Pending",              
	}
	quantity := r.FormValue("productQuantity")
	parsedQuantity, err := strconv.Atoi(quantity)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid quantity", err)
		res.Error = "Invalid quantity"
		return
	}
	ord.ProductQuantity = int32(parsedQuantity)

	db := db.OrderDB{MongoCollection: svc.MongoCollection}
	createOrd, err := db.CreateNewOrder(&ord)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("order creation error", err)
		res.Error = err.Error()
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Println("order created", createOrd)
}
