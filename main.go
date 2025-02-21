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
	// Load database URL from environment variables
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("❌ DATABASE_URL is not set. Please configure it in your environment.")
	}

	// Initialize the database
	database.InitDB(dbURL)

	// Create a new Gin router
	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Change "*" to your frontend URL in production
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup API routes
	routes.SetupRoutes(r)

	// Start the server
	log.Println("✅ Server running on port 8080")
	r.Run(":8080")
}
