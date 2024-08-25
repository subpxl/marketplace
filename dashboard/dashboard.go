package dashboard

import (
	"marketplace/product"

	"gorm.io/gorm"
)

type Address struct {
	gorm.Model

	City     string
	Street   string
	State    string
	Pincode  string
	Country  string
	Landmark string
	Location string
}

type Contact struct {
	gorm.Model

	Instagram      string
	Facebook       string
	Linkedin       string
	Twitter        string
	Youtube        string
	MobileNumber   string
	WhatsappNumber string
	Email          string
}

type CarouselItem struct {
	gorm.Model

	ID       string
	ImageURL string
	AltText  string
	LinkURL  string
}

type Shop struct {
	gorm.Model

	Name          string
	Logo          string
	TopMessage    string
	SecondMessage string
	CarouselItems []CarouselItem
	Products      []product.Product
	Contact       Contact // Embedded Contact struct
	Message       string
	About         string
	Address       Address // Embedded Address struct
	Terms         string
	PrivacyPolicy string
}

var shop = Shop{
	Name:          "My Shop",
	Logo:          "/static/logo.png", // Path or URL to the shop logo
	TopMessage:    "Welcome to My Shop!",
	SecondMessage: `Sale! Up to 50% off on selected items!`,
	Products:      product.Products,
	CarouselItems: []CarouselItem{
		{ID: "slide1", ImageURL: "/static/1.jpg", AltText: "Banner 1", LinkURL: "/2323"},
		{ID: "slide2", ImageURL: "/static/2.jpg", AltText: "Banner 2", LinkURL: "232323"},
		{ID: "slide3", ImageURL: "/static/3.jpg", AltText: "Banner 3", LinkURL: "4r4r"}},
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
