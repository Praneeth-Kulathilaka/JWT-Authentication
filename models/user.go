package models
import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)
type User struct{
	ID primitive.ObjectID `json:"id" bson:"id"`
	Username string             `json:"username" bson:"username"`
    Password string             `json:"password" bson:"password"`
}