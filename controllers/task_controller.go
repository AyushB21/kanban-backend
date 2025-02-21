package controllers

import (
	"net/http"
	"taskmanager/database"
	"taskmanager/models"

	"github.com/gin-gonic/gin"
)

// Create Task
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save task in database
	if err := database.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully"})
}

// Get All Tasks
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	if err := database.DB.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// Move Task to Another Column
func MoveTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update task status
	if err := database.DB.Model(&models.Task{}).Where("id = ?", id).Update("column", task.Column).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to move task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task moved successfully"})
}
