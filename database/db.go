package database

import (
	"log"
	"taskmanager/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dbURL string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	log.Println("✅ Successfully connected to PostgreSQL!")

	// Auto-migrate database tables
	err = DB.AutoMigrate(&models.Task{}, &models.User{})
	if err != nil {
		log.Fatal("❌ Failed to migrate database tables:", err)
	}
}
