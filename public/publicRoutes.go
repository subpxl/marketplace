package public

import (
	"marketplace/config"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, appConfig *config.AppConfig) {
	publicGroup := r.Group("/")
	handler := NewPublicHandler(appConfig)
	publicGroup.GET("/product", handler.ProductHandler)
	publicGroup.GET("/terms", handler.TermsHandler)
	publicGroup.GET("/", handler.HandleShop)

}
