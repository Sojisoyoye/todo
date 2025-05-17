package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/Sojisoyoye/todo/internal/repository"
	"github.com/Sojisoyoye/todo/internal/config"
	"github.com/Sojisoyoye/todo/internal/handlers"
	"github.com/Sojisoyoye/todo/internal/models"
	"github.com/Sojisoyoye/todo/internal/service"

)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	db, err := config.ConnectDB()

	// Auto migrate models
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	// Initialize repository, service, and handlers
	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)

	// Set up Gin router
	r := gin.Default()

	// Set up routes
	handlers.RegisterRoutes(r, todoService)

	// Start the server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	log.Printf("Server running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
