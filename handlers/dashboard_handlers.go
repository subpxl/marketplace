package handlers

import (
	"marketplace/config"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	AppConfig *config.AppConfig
}

func NewDashboardHandler(appConfig *config.AppConfig) *DashboardHandler {
	return &DashboardHandler{AppConfig: appConfig}
}

func (h *DashboardHandler) HandleDashboard(ctx *gin.Context) {
	ctx.HTML(200, "dashboard.html", nil)
}

func (h *DashboardHandler) HandleCollections(ctx *gin.Context) {
	ctx.HTML(200, "collections.html", nil)
}

func (h *DashboardHandler) HandleSettings(ctx *gin.Context) {
	ctx.HTML(200, "settings.html", nil)
}
