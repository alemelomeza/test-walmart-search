package main

// Product defines the properties of a product
type Product struct {
	ID          int    `bson:"id" json:"id"`
	Brand       string `bson:"brand" json:"brand"`
	Description string `bson:"description" json:"description"`
	Image       string `bson:"image" json:"image"`
	Price       int    `bson:"price" json:"price"`
}
