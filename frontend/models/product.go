package model 

type Product struct {
	ID string `bson:"_id,omitempty"`
	Name string `bson:"name"`
	Category string `bson:"category"`
	Price float64 `bson:"price"`
	Quantity int32 `bson:"quantity"`
	Image string `bson:"image"`
	Status bool `bson:"isAvailable"`
	Description string `bson:"description"`
}