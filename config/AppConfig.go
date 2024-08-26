package config

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
	"gorm.io/gorm"
)

type AppConfig struct {
	DB           *gorm.DB
	PublicRender *render.Render
	Render       *render.Render
	Router       *gin.Engine
}

func NewAppConfig(db *gorm.DB, render *render.Render, publicRender *render.Render, router *gin.Engine) *AppConfig {
	return &AppConfig{
		DB:           db,
		Render:       render,
		PublicRender: publicRender,
		Router:       router,
	}
}
