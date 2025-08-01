package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email" bson:"email"`
	Name     string             `json:"name" bson:"name"`
	Username string             `json:"username" bson:"username"`
	Role     string             `json:"role" bson:"role"`
	Password string             `json:"password" bson:"password"`
} 
