package main

import (
	"log"
	"os"

	"github.com/dalia3aly/go-sqlite-backend/config"
	"github.com/dalia3aly/go-sqlite-backend/models"
	"github.com/dalia3aly/go-sqlite-backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file in development
	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("Error loading .env file")
		}
	}

	// Initialize database
	config.InitDB()

	// Auto migrate the schema
	if err := config.DB.AutoMigrate(&models.Category{}, &models.Todo{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Create default category if none exists
	var count int64
	config.DB.Model(&models.Category{}).Count(&count)
	if count == 0 {
		defaultCategory := models.Category{
			Name: "Default Category",
		}
		if err := config.DB.Create(&defaultCategory).Error; err != nil {
			log.Printf("Error creating default category: %v", err)
		}

		// Create a sample todo
		todo := models.Todo{
			Title:       "Welcome to Todo App!",
			Description: "This is your first todo item. You can add more todos using the API.",
			CategoryID:  defaultCategory.ID,
		}
		if err := config.DB.Create(&todo).Error; err != nil {
			log.Printf("Error creating sample todo: %v", err)
		}
	}

	// Initialize Gin router
	if os.Getenv("GO_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	// Setup CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Setup routes
	routes.SetupRoutes(router)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
