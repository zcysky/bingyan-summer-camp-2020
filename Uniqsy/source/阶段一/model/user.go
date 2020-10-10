package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserID   primitive.ObjectID `json:"user_id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Phone    string             `json:"phone" bson:"phone"`
	Email    string             `json:"email" bson:"email"`
}