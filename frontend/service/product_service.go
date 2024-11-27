package service

import (
	"polar-shop/db"

	"polar-shop/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	MongoCollection *mongo.Collection
}

func (svc *ProductService) GetAvailableProducts() ([]model.Product, error) {
    db := db.ProductDB{MongoCollection: svc.MongoCollection}
    return db.GetAvailableProducts()
}

func (svc *ProductService) GetProductByName(productId string) (*model.Product, error) {
	db := db.ProductDB{MongoCollection: svc.MongoCollection}
	return db.GetProductByName(productId)
}