package dashboard

import (
	"marketplace/config"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, appConfig *config.AppConfig) {
	dahboardGroup := r.Group("/dashboard")
	handler := NewDashboardHandler(appConfig)
	dahboardGroup.GET("/collections", handler.HandleCollections)
	dahboardGroup.POST("/settings", handler.HandleSettings)
	dahboardGroup.GET("/", handler.HandleDashboard)

}
