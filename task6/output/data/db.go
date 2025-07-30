package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TaskCollection *mongo.Collection

func InitDB() {
	// // Load .env file
    // err := godotenv.Load()
    // if err != nil {
    //     log.Fatal("Error loading .env file")
    // }

    // MONGO_URI := os.Getenv("MONGO_URI")
    // if MONGO_URI == "" {
    //     log.Fatal("MONGO_URI not set in environment")
    // }

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	
	ctx, cancel := context.WithTimeout(context.Background(), 20 *time.Second)

	defer cancel()

	// create the database connection 
	client, err := mongo.Connect(ctx, clientOptions) 
	if err != nil {
		log.Fatal(err)
	}

	// now assign to the task collection reference  
	TaskCollection = client.Database("taskdb").Collection("tasks")


}