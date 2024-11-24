package db 

import "go.mongodb.org/mongo-driver/mongo"

type ProductDB struct {
	MongoCollection *mongo.Collection
}