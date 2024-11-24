package db 

import (
	"go.mongodb.org/mongo-driver/mongo"
	
	"polar-shop/models"
)

type ProductDB struct {
	MongoCollection *mongo.Collection
}