package data

import (
	"context"
	"log"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TO GET ALL TASKS
func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	
	cursor, err := TaskCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO()) 

	// lopping through each documents and attack to the array  
	for cursor.Next(context.TODO()) {

		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	    
		if err := cursor.Err(); err != nil {
			return nil, err
		}
	}
	return tasks, nil
}
// GET TASK BY ID  
func GetTaskByID(id string) (*models.Task, error) {
	var task models.Task
	filter := bson.M {"id":id}
	// find that task with give id 
	err := TaskCollection.FindOne(context.TODO(),filter).Decode(&task)
	if err != nil {
		return nil, err
	}
	
	return &task, nil
}

// ADDING NEW TASK TO THE BUILT-IN DATABASE 
func AddNewTask(task models.Task) error {
	task.ID = primitive.NewObjectID()
	_, err := TaskCollection.InsertOne(context.TODO(), task)

	return err
	}
	

// UPDATING AN EXISTING TASK  
func UpdateTask(id string, updated models.Task) error {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}
	filter := bson.M {"_id":objectID}
	update := bson.M{
		"$set": bson.M{
			"title":       updated.Title,
			"description": updated.Description,
			"due_date":    updated.DueDate,
			"status":      updated.Status,
		},
	}

	_, err = TaskCollection.UpdateOne(context.TODO(),filter,  update)
    return err

}
// DELETING TASK 
func DeleteTask(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err // invalid ObjectID format
	}

	filter := bson.M{"_id": objectID}
	result, err := TaskCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err // DB error
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments // nothing matched the ID
	}

	return nil
}
