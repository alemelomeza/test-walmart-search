package main

// Product defines the properties of a product
type Product struct {
	// Bad order struct can cause more memory usage https://link.medium.com/hiTPDoHZ0ab
	Brand       string `bson:"brand" json:"brand"`
	Description string `bson:"description" json:"description"`
	Image       string `bson:"image" json:"image"`
	Price       int    `bson:"price" json:"price"`
	ID          int    `bson:"id" json:"id"`
}
