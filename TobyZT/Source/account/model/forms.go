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

type SignupJsonForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type SignupForm struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Phone    string `json:"phone" bson:"phone"`
	Email    string `json:"email" bson:"email"`
}

type QueryForm struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type JWTInfo struct {
	Secret string `json:"secret"`
	Expire int    `json:"expire"`
}

type SMTPInfo struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     string `json:"port"`
}

// JsonToSignupForm coverts a SignupJsonForm to a SignupForm
func JsonToSignupForm(jsonForm SignupJsonForm) (form SignupForm) {
	form = SignupForm{
		Username: jsonForm.Username,
		Password: jsonForm.Password,
		Phone:    jsonForm.Phone,
		Email:    jsonForm.Email,
	}
	return form
}
