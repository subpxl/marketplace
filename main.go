package main

import (

	// If using Redis for session storage

	"html/template"
	"log"
	"marketplace/config"
	"marketplace/dashboard"
	"marketplace/product"
	"marketplace/public"
	"marketplace/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&dashboard.Shop{}, &product.Product{}, &dashboard.CarouselItem{}, &dashboard.Contact{}, &dashboard.Address{})

	return db
}

func initTemplates() *template.Template {
	templates := template.Must(template.ParseGlob("templates/*"))
	return templates
}

func main() {

	// Initialize Gin router
	db := initDB()
	templates := initTemplates()
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	appConfig := config.NewAppConfig(db, templates, router)

	// Set up static file serving and HTML templates
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	// Setup routes

	// public routers

	// // // dashboard routes

	// Initialize the product handler

	// // auth handlers
	public.SetupRoutes(router, appConfig)
	user.SetupRoutes(router, appConfig)
	dashboard.SetupRoutes(router, appConfig)
	product.SetupRoutes(router, appConfig)

	// Start the server
	err := router.Run(":8000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
