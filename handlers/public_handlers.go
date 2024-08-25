package handlers

import (
	"marketplace/config"
	"marketplace/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PublicHandler struct {
	AppConfig *config.AppConfig
}

func NewPublicHandler(appConfig *config.AppConfig) *PublicHandler {
	return &PublicHandler{AppConfig: appConfig}
}

func (h *PublicHandler) HandleShop(c *gin.Context) {
	shop := models.Shop{}
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
}

func (h *PublicHandler) ProductHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "product.html", nil)
}

func (h *PublicHandler) TermsHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "terms.html", nil)
}
