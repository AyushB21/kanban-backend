package models

import "gorm.io/gorm"

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
