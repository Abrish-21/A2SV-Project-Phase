package controller

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

// CONTROLLER FOR GETTING ALL TASKS
func GetAllTasks(ctx *gin.Context) {
	tasks := data.GetAllTasks()
	ctx.JSON(http.StatusOK, tasks)
}
// GETTING  A TASK BY ID 
func GetTaskByID(ctx *gin.Context){
	id:= ctx.Param("id")
	task, found := data.GetTaskByID(id)
	if !found {
		ctx.JSON(http.StatusNotFound,gin.H{ "Error": " Task not Found",
		})
		return 
	}
	ctx.JSON(http.StatusOK,task )
}

// CREATING A NEW TASK 
func AddNewTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var newTask  models.Task
	if err:= ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "error": err.Error()})
		return 
		
	}
	data.AddNewTask(id, newTask)
	ctx.JSON(http.StatusCreated, gin.H{ "Message:": "Task created"})

}

// UPDATING TASK 
func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask models.Task

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if data.UpdateTask(id, updatedTask) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	}
}

// DELETING TASK 
func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	if data.DeleteTask(id) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	}
}

