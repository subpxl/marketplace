package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "shop.html", gin.H{
			"name":           shop.Name,
			"logo":           shop.Logo,
			"topMessage":     shop.TopMessage,
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
		})
	})

	r.GET("/product", func(c *gin.Context) {
		c.HTML(http.StatusOK, "product.html", nil)
	})

	r.Run(":8000")
}
