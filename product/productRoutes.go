package product

import (
	"marketplace/config"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, appConfig *config.AppConfig) {
	productGroup := r.Group("/product")
	handler := NewProductHandler(appConfig)
	productGroup.GET("/create", handler.CreateProduct)
	productGroup.POST("/delete/:id", handler.DeleteProduct)
	productGroup.GET("/update/:id", handler.UpdateProduct)
	productGroup.POST("/get", handler.GetProducts)
	productGroup.POST("/get/:id", handler.GetProduct)
}
