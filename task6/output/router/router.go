package router

import (
	"task_manager/controller"
	Middleware "task_manager/middleware"

	"github.com/gin-gonic/gin"
)


func RunRouter() {
	router := gin.Default()

	// Handle path routes
	router.GET("/tasks", controller.GetTaskByID)
	router.GET("/tasks/:id", controller.GetTaskByID)
	router.PUT("/tasks/:id", Middleware.AuthMiddleWare(),controller.UpdateTask)
	router.POST("/tasks", Middleware.AuthMiddleWare() ,controller.AddNewTask)
	router.POST("/register", controller.UserRegisterController)
	router.POST("/login", controller.UserLoginController)
	router.GET("/admin_page", Middleware.AuthMiddleWare(), Middleware.AuthRoleMiddleWare(), controller.GetAdminPageController)
	router.DELETE("/tasks/:id", Middleware.AuthMiddleWare(),controller.DeleteTask) 
	router.GET("/user_profile", Middleware.AuthMiddleWare(),controller.GetUserProfileController)

	// Start router
	router.Run()
}