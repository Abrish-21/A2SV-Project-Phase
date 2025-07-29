package router

import (
	"task_manager/controller"
	Middleware "task_manager/middleware"

	"github.com/gin-gonic/gin"
)


func RunRouter() {
	router := gin.Default()

	// Handle path routes
	router.GET("/tasks", controller.GetAllTasks)
	router.GET("/tasks/:id", controller.GetTaskByID)
	router.PUT("/tasks/:id", Middleware.AuthMiddleWare(),controller.UpdateTask)
	router.POST("/tasks", middleware.AuthMiddleWare() ,controller.AddNewTask)
	router.POST("/register", controllers.UserRegisterController)
	router.POST("/login", controllers.UserLoginController)
	router.GET("/admin_page", middleware.AuthMiddleWare(), middleware.AuthRoleMiddleWare(), controllers.GetAdminPageController)
	router.DELETE("/tasks/:id", middleware.AuthMiddleWare(),controllers.DeleteTaskController) 
	router.GET("/user_profile", middleware.AuthMiddleWare(), controllers.GetUserProfileController)

	// Start router
	router.Run()
}