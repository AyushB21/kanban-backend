package routes

import (
	"taskmanager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/signup", controllers.RegisterUser) // ✅ Add signup route
	router.POST("/signin", controllers.LoginUser) // ✅ Add signin route

	router.POST("/tasks", controllers.CreateTask)
	router.GET("/tasks", controllers.GetTasks)
	router.PUT("/tasks/:id/move", controllers.MoveTask)
}
