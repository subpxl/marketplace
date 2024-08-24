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
		})
	})

	r.GET("/product", func(c *gin.Context) {
		c.HTML(http.StatusOK, "product.html", nil)
	})
	r.GET("/terms", func(c *gin.Context) {

		c.HTML(http.StatusOK, "terms.html", nil)
	})


	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	r.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "signup.html", nil)
	})

	r.GET("/createshop", func(c *gin.Context) {
		c.HTML(200, "createshop.html", nil)
	})


	dashboard := r.Group("/dashboard")
	dashboard.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "dashboard.html", nil)
	})

	dashboard.GET("/collections", func(ctx *gin.Context) {
		ctx.HTML(200, "collections.html", nil)
	})

	dashboard.GET("/settings", func(ctx *gin.Context) {
		ctx.HTML(200, "settings.html", nil)
	})


	dashboard.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	dashboard.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "signup.html", nil)
	})

	r.Run(":8000")
}
