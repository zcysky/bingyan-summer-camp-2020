package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserID   primitive.ObjectID `json:"userid" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Phone    string             `json:"phone" bson:"phone"`
	Email    string             `json:"email" bson:"email"`
}

type LoginForm struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type SignupForm struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Phone    string `json:"phone" bson:"phone"`
	Email    string `json:"email" bson:"email"`
}

type TokenForm struct {
	UserID   string
	Email    string
	Password string
}

type JsonInfo struct {
	Secret string `json:"secret"`
	Expire int    `json:"expire"`
}

type UpdateForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
