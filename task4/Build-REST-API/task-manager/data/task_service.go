package data

import (
	"task_manager/models"
	"time"
)

var tasks =  []models.Task {
	{ID: "1", Title: "task 1", Description:"First task", DueDate: time.Now(), Status:"Pending"},
	{ID: "2", Title: "task 2", Description:"Second task", DueDate: time.Now().AddDate(0,0,1), Status:"In progress"},
	{ID: "3", Title: "task 3", Description:"Third task", DueDate: time.Now().AddDate(0,0,2), Status:"Completed"},

	
}

// TO GET ALL TASKS 
func GetAllTasks() []models.Task {
	return tasks 
}
// GET TASK BY ID 
func GetTaskByID(id string) (models.Task, bool) {
	for _, task := range tasks {
		if task.ID == id {
			return  task, true 
		}
	}
	return models.Task{}, false
}

// ADDING NEW TASK TO THE BUILT-IN DATABASE 
func AddNewTask( id string , newTask models.Task) {
	tasks = append(tasks, newTask)

}

// UPDATING AN EXISTING TASK  
func UpdateTask(id string, updated models.Task) bool {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].ID = updated.ID
			tasks[i].Title = updated.Title
			tasks[i].Description = updated.Description
			tasks[i].DueDate = updated.DueDate
			tasks[i].Status = updated.Status
			return true
		}
	}
	return false
}

// DELETING TASK 
func DeleteTask(id string) bool {
	for i,task := range tasks {

		if task.ID == id {
			tasks = append(tasks[:i],tasks[i+1:]...)

			return true
		}
	}
	return false 
}