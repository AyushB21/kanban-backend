package routes

import (
	"taskmanager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Auth routes
	api.POST("/register", controllers.RegisterUser)
	api.POST("/login", controllers.LoginUser)

	// Task routes
	api.POST("/tasks", controllers.CreateTask)
	api.GET("/tasks", controllers.GetTasks)
	api.PATCH("/tasks/:id", controllers.MoveTask)
}
