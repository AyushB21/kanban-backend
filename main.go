
package database

import (
	"log"

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
	err = DB.AutoMigrate(&Task{}, &User{})
	if err != nil {
		log.Fatal("❌ Failed to migrate database tables:", err)
	}
}

// Task Model
type Task struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	Status      string `gorm:"default:pending"`
	Column      string `gorm:"default:To Do"`
	Assignee    string
}

// User Model
type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
