package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type Trainer struct {
	Name string 
	Age int 
	City string
}

func main() {
	// The follwing are steps to connect to mongoDB driver 
	// step one: setting options 
	clientOptions := options.Client().ApplyURI("mongodb+srv://zere28yalem21:xCyARVNSSBO2TQBu@cluster0.w9auirk.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

	// connect to mongoDB 
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	// check the connection 
	err =  client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB sucessfully connected!")

	collection := client.Database("text").Collection("trainer")
// 

// create Trainer structs 
ash := Trainer{"ash", 21, "Pallet Town"}
misty := Trainer{"Misty", 10, "Cerulean City"}
brock := Trainer{"Brock", 15, "Pewter City"}

insertResult, err := collection.InsertOne(context.Background(), ash)

if err != nil {
	log.Fatal(err)
}
fmt.Println("inserted a single document", insertResult.InsertedID)


// adding many trainers to the collection  
trainers := []interface{}{misty, brock, ash}

insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
if err != nil {
    log.Fatal(err)
}

fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
// pwd:  xCyARVNSSBO2TQBu

	// if we want to close the DB instance 
	// err = client.Disconnect(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Database successfully disconnected!")

}