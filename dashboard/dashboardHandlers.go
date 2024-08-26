package dashboard

import (
	"marketplace/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
)

type DashboardHandler struct {
	AppConfig *config.AppConfig
}

func NewDashboardHandler(appConfig *config.AppConfig) *DashboardHandler {
	return &DashboardHandler{AppConfig: appConfig}
}

func (h *DashboardHandler) HandleDashboard(ctx *gin.Context) {
	// ctx.HTML(200, "dashboard.html", nil)
	if err := h.AppConfig.Render.HTML(ctx.Writer, http.StatusOK, "dashboard/dashboard", gin.H{"Title": "Dashboard"}, render.HTMLOptions{
		Layout: "", // Disable layout
	}); err != nil {
		ctx.String(http.StatusInternalServerError, "Error rendering template: %v", err)
	}
}

func (h *DashboardHandler) HandleCollections(ctx *gin.Context) {
	// ctx.HTML(200, "collections.html", nil)
	if err := h.AppConfig.Render.HTML(ctx.Writer, http.StatusOK, "dashboard/collections", gin.H{"Title": "Dashboard"}, render.HTMLOptions{
		Layout: "", // Disable layout
	}); err != nil {
		ctx.String(http.StatusInternalServerError, "Error rendering template: %v", err)
	}
}

func (h *DashboardHandler) HandleSettings(ctx *gin.Context) {
	// ctx.HTML(200, "settings.html", nil)
	if err := h.AppConfig.Render.HTML(ctx.Writer, http.StatusOK, "dashboard/settings", gin.H{"Title": "Dashboard"}, render.HTMLOptions{
		Layout: "", // Disable layout
	}); err != nil {
		ctx.String(http.StatusInternalServerError, "Error rendering template: %v", err)
	}
}
