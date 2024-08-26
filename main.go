package main

import (

	// If using Redis for session storage

	"log"
	"marketplace/config"
	"marketplace/dashboard"
	"marketplace/product"
	"marketplace/public"
	"marketplace/user"

	"github.com/unrolled/render"

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

func initRender() *render.Render {
	return render.New(render.Options{
		Directory:  "./templates/backend", // Use full path if necessary
		Extensions: []string{".html"},
		Layout:     "layout", // Remove the .html extension
	})
}

func initPublicRender() *render.Render {
	return render.New(render.Options{
		Directory:  "./templates/frontend",
		Extensions: []string{".html"},
		Layout:     "layout",
		
	})
}

func main() {

	// Initialize Gin router

	db := initDB()
	r := initRender()
	publicRender := initPublicRender()

	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	appConfig := config.NewAppConfig(db, r, publicRender, router)

	// Set up static file serving and HTML templates
	router.Static("/static", "./static")

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
