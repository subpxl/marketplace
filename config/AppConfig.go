package config

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppConfig struct {
	DB        *gorm.DB
	Templates *template.Template
	Router    *gin.Engine
}

func NewAppConfig(db *gorm.DB, templates *template.Template, router *gin.Engine) *AppConfig {
	return &AppConfig{
		DB:        db,
		Templates: templates,
		Router:    router,
	}
}
