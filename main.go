package main

import (
	"log"
	"os"
	"taskmanager/database"
	"taskmanager/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	database.InitDB(os.Getenv("DATABASE_URL")) // Read DB URL from env variable

	// Create a new Gin router
	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup API routes
	routes.SetupRoutes(r)

	// Get the PORT from the environment (Render assigns it dynamically)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to 8080 if PORT is not set
	}

	// Start the server
	log.Println("🚀 Server running on port " + port)
	r.Run(":" + port)
}
