package controllers

import (
	"taskmanager/database"
	"taskmanager/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Task
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := database.DB.Exec("INSERT INTO tasks (title, description, column, assignee) VALUES (?, ?, ?, ?)",
		task.Title, task.Description, task.Column, task.Assignee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully"})
}

// Get All Tasks
func GetTasks(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, title, description, status, column, assignee FROM tasks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Column, &task.Assignee); err != nil {
			continue
		}
		tasks = append(tasks, task)
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

	_, err := database.DB.Exec("UPDATE tasks SET column=? WHERE id=?", task.Column, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to move task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task moved successfully"})
}
