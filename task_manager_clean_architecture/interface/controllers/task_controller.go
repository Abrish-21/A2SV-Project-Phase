package controllers

import (
	"net/http"
	"task_manager_clean_architecture/domain/models"
	"task_manager_clean_architecture/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskController struct {
	Usecase *usecase.TaskUsecase
}

func NewTaskController(u *usecase.TaskUsecase) *TaskController {
	return &TaskController{Usecase: u}
}

func (ctrl *TaskController) Create(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	task.ID = uuid.New().String()
	if err := ctrl.Usecase.Create(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create task"})
		return
	}
	c.JSON(http.StatusCreated, task)
}

func (ctrl *TaskController) GetAll(c *gin.Context) {
	userID := c.Query("user_id")
	tasks, _ := ctrl.Usecase.GetAll(userID)
	c.JSON(http.StatusOK, tasks)
}
