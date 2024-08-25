package main

import (

	// If using Redis for session storage

	"html/template"
	"log"
	"marketplace/config"
	"marketplace/handlers"
	"marketplace/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Shop{}, &models.Product{}, &models.CarouselItem{}, &models.Contact{}, &models.Address{})

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
	publicHandler := handlers.NewPublicHandler(appConfig)

	router.GET("/", publicHandler.HandleShop)
	router.GET("/product", publicHandler.ProductHandler)
	router.GET("/terms", publicHandler.TermsHandler)

	// // // dashboard routes
	dashboardHandler := handlers.NewDashboardHandler(appConfig)
	router.Any("/dashboard", dashboardHandler.HandleDashboard)
	router.Any("/collections", dashboardHandler.HandleCollections)
	router.Any("/settings", dashboardHandler.HandleSettings)

	// Initialize the product handler
	productHandler := handlers.NewProductHandler(appConfig)
	// Set up routes
	router.POST("/products", productHandler.CreateProduct)
	router.GET("/products", productHandler.GetProducts)
	router.GET("/products/:id", productHandler.GetProduct)
	router.PUT("/products/:id", productHandler.UpdateProduct)
	router.DELETE("/products/:id", productHandler.DeleteProduct)

	
	// router.Any("/createshop", CreateShop)
	// router.Any("/createshop", CreateProduct)
	// router.Any("/createshop", CreateCategory)
	// router.Any("/createshop", CreateCollection)

	// // auth handlers
	authHandler := handlers.NewAuthHandler(appConfig)

	router.Any("/login", authHandler.HandleLogin)
	router.Any("/signup", authHandler.HandleSignup)
	router.Any("/logout", authHandler.HandleLogout)

	// Start the server
	err := router.Run(":8000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
