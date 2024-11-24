package service

type ProductService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data interface{} `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

