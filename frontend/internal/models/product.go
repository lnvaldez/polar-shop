package models 

type Product struct {
	ID string `bson:"_id,omitempty"`
	Name string `bson:"name"`
	Category string `bson:"category"`
	Price int32 `bson:"price"`
	Quantity string `bson:"quantity"`
	Image string `bson:"image"`
	Status bool `bson:"isAvailable"`
	Description string `bson:"description"`
}