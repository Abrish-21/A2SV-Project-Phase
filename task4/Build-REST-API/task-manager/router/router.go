package router

import (
	"task_manager/controller"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/tasks", controller.GetAllTasks)
	r.GET("/tasks/:id", controller.GetTaskByID)
	r.POST("/tasks", controller.AddNewTask)
	r.PUT("/tasks/:id", controller.UpdateTask)
	r.DELETE("/task/:id", controller.DeleteTask)

	return r
}