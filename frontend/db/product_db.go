package db

import (
	"fmt"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"

	"polar-shop/models"
)

type ProductDB struct {
	MongoCollection *mongo.Collection
}

func (db *ProductDB) GetAvailableProducts() ([]model.Product, error) {
	filter := bson.D{{Key: "isAvailable", Value: true}}

	result, err := db.MongoCollection.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	var prods []model.Product
	err = result.All(context.Background(), &prods)
	
	if err != nil {
		return nil, fmt.Errorf("error %s", err.Error())
	}

	return prods, nil
}