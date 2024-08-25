package user

import (
	"marketplace/config"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, appConfig *config.AppConfig) {
	userGroup := r.Group("/user")
	handler := NewUserHandler(appConfig)
	userGroup.GET("/login", handler.HandleLogin)
	userGroup.POST("/login", handler.HandleLogin)
	userGroup.GET("/signup", handler.HandleSignup)
	userGroup.POST("/signup", handler.HandleSignup)
	userGroup.POST("/logout", handler.HandleLogout)
}
