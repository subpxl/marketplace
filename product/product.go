package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	ID          int
	Name        string
	Description string
	Price       float64
	Image       string
	Slug        string
	ShopID      uint // Foreign key referencing the Shop model

}

// Define products
var Products = []Product{
	{ID: 1, Name: "Elegant Tote", Slug: "Elegant-Tote", Description: "Luxurious leather bag", Price: 199.99, Image: "/static/1.jpg"},
	{ID: 2, Name: "Stylish Backpack", Slug: "Stylish-Backpack", Description: "Comfortable and spacious", Price: 149.99, Image: "/static/2.jpg"},
	{ID: 3, Name: "Classic Wallet", Slug: "Classic-Wallet", Description: "Premium quality wallet", Price: 89.99, Image: "/static/3.jpg"},
	{ID: 4, Name: "Sporty Watch", Slug: "Sporty-Watch", Description: "Durable and stylish", Price: 299.99, Image: "/static/4.jpg"},
}
