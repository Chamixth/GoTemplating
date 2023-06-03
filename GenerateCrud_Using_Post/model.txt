package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents a user entity
type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`            //unique id for the user
	Name     string             `json:"name" bson:"name,omitempty"`         //name of the user
	Email    string             `json:"email" bson:"email,omitempty"`       //email of the user
	Password string             `json:"password" bson:"password,omitempty"` //password of the user
}
