package public

import (
	"marketplace/config"
	"marketplace/dashboard"
	"marketplace/product"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
)

type PublicHandler struct {
	AppConfig *config.AppConfig
}

func NewPublicHandler(appConfig *config.AppConfig) *PublicHandler {
	return &PublicHandler{AppConfig: appConfig}
}

var shop = dashboard.Shop{
	Name:          "My Shop",
	Logo:          "/static/logo.png", // Path or URL to the shop logo
	TopMessage:    "Welcome to My Shop!",
	SecondMessage: `Sale! Up to 50% off on selected items!`,
	Products:      product.Products,
	CarouselItems: []dashboard.CarouselItem{
		{ImageURL: "/static/1.jpg", AltText: "Banner 1", LinkURL: "/2323"},
		{ImageURL: "/static/2.jpg", AltText: "Banner 2", LinkURL: "232323"},
		{ImageURL: "/static/3.jpg", AltText: "Banner 3", LinkURL: "4r4r"}},
	Contact: dashboard.Contact{
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
	Address: dashboard.Address{
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

func (h *PublicHandler) HandleShop(c *gin.Context) {

	if err := h.AppConfig.PublicRender.HTML(c.Writer, http.StatusOK, "shop", gin.H{
		"name":           shop.Name,
		"logo":           shop.Logo,
		"topMessage":     shop.TopMessage,
		"SecondMessage":  shop.SecondMessage,
		"products":       shop.Products,
		"ContactNumber":  shop.Contact.MobileNumber,
		"WhatsappNumber": shop.Contact.WhatsappNumber,
		"Contact":        shop.Contact,
		"Message":        shop.Message,
		"ShopName":       shop.Name,
		"ShopLogo":       shop.Logo,
		"About":          shop.About,
		"Address":        shop.Address,
		"Terms":          shop.Terms,
		"PrivacyPolicy":  shop.PrivacyPolicy,
		"SocialMedia":    shop.Contact, // Passing the entire contact information
		"CarouselItems":  shop.CarouselItems,
	}, render.HTMLOptions{
		Layout: "", // Disable layout
	}); err != nil {
		c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
	}
}

func (h *PublicHandler) ProductHandler(c *gin.Context) {
	// c.HTML(http.StatusOK, "product.html", nil)
	if err := h.AppConfig.PublicRender.HTML(c.Writer, http.StatusOK, "product", gin.H{"Title": "Dashboard"}, render.HTMLOptions{
		Layout: "", // Disable layout
	}); err != nil {
		c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
	}
}

func (h *PublicHandler) TermsHandler(c *gin.Context) {
	// c.HTML(http.StatusOK, "terms.html", nil)
	if err := h.AppConfig.PublicRender.HTML(c.Writer, http.StatusOK, "terms", gin.H{"Title": "Dashboard"}, render.HTMLOptions{
		Layout: "", // Disable layout
	}); err != nil {
		c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
	}
}
