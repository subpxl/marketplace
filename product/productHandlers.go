package product

import (
	"marketplace/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
)

type ProductHandler struct {
	AppConfig *config.AppConfig
}

func NewProductHandler(appConfig *config.AppConfig) *ProductHandler {
	return &ProductHandler{AppConfig: appConfig}
}

// Create a new product
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "product/create.html", gin.H{})
		if err := h.AppConfig.Render.HTML(c.Writer, http.StatusOK, "product/create", gin.H{"Title": "Dashboard"}, render.HTMLOptions{
			Layout: "", // Disable layout
		}); err != nil {
			c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
		}

	} else {

		var product Product

		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := h.AppConfig.DB.Create(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, product)
	}
}

// Get all products
func (h *ProductHandler) GetProducts(c *gin.Context) {
	var products []Product
	if err := h.AppConfig.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// Get a single product by ID
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product
	if err := h.AppConfig.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// Update a product
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product
	if err := h.AppConfig.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.AppConfig.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// Delete a product
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := h.AppConfig.DB.Delete(&Product{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
