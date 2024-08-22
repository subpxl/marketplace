package main

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Image       string
	Slug        string
}

// Define products
var Products = []Product{
	{ID: 1, Name: "Elegant Tote", Slug: "Elegant-Tote", Description: "Luxurious leather bag", Price: 199.99, Image: "/static/1.jpg"},
	{ID: 2, Name: "Stylish Backpack", Slug: "Stylish-Backpack", Description: "Comfortable and spacious", Price: 149.99, Image: "/static/2.jpg"},
	{ID: 3, Name: "Classic Wallet", Slug: "Classic-Wallet", Description: "Premium quality wallet", Price: 89.99, Image: "/static/3.jpg"},
	{ID: 4, Name: "Sporty Watch", Slug: "Sporty-Watch", Description: "Durable and stylish", Price: 299.99, Image: "/static/4.jpg"},
}

type Address struct {
	City     string
	Street   string
	State    string
	Pincode  string
	Country  string
	Landmark string
	Location string
}

type Contact struct {
	Instagram      string
	Facebook       string
	Linkedin       string
	Twitter        string
	Youtube        string
	MobileNumber   string
	WhatsappNumber string
	Email          string
}

type Shop struct {
	Name          string
	Logo          string
	TopMessage    string
	Products      []Product
	Contact       Contact // Embedded Contact struct
	Message       string
	About         string
	Address       Address // Embedded Address struct
	Terms         string
	PrivacyPolicy string
}

var shop = Shop{
	Name:       "My Shop",
	Logo:       "/static/logo.png", // Path or URL to the shop logo
	TopMessage: "Welcome to My Shop!",
	Products:   Products,
	Contact: Contact{
		Instagram:      "https://instagram.com/fb",
		Facebook:       "https://facebook.com/fb",
		Linkedin:       "https://linkedin.com/fb",
		Twitter:        "https://twitter.com/fb",
		Youtube:        "https://youtube.com/fb",
		MobileNumber:   "919109517579",
		WhatsappNumber: "919109517579",
		Email:          "contact@myshop.com",
	},
	Message: "I'm interested in your product.",
	About: `We offer a variety of high-quality products.
			We offer a variety of high-quality products.
			We offer a variety of high-quality products.
			We offer a variety of high-quality products.
			We offer a variety of high-quality products.`,
	Address: Address{
		City:     "Anytown",
		Street:   "123 Main St",
		State:    "StateName",
		Pincode:  "123456",
		Country:  "CountryName",
		Landmark: "Near the central park",
		Location: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3022.1422937950147!2d-73.98731968482413!3d40.75889497932681!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x89c25855c6480299%3A0x55194ec5a1ae072e!2sTimes+Square!5e0!3m2!1sen!2sus!4v1560412335569!5m2!1sen!2sus",
	},
	Terms:         "All terms apply.",
	PrivacyPolicy: "Your privacy is important to us.",
}
