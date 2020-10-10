package model

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginForm struct {
	IsAdmin		bool	`json:"is_admin"`
	Email    	string 	`json:"email" bson:"email"`
	Password 	string 	`json:"password" bson:"password"`
}

func VerifyLoginForm(loginForm LoginForm) (id string, err error) {
	err = TestConnection()
	if err != nil {
		return "", err
	}

	var resultUser User
	filter := bson.M{"email":	loginForm.Email}
	collection := new(mongo.Collection)
	if loginForm.IsAdmin {
		collection = client.Database("users").Collection("admin")
	} else {
		collection = client.Database("users").Collection("users")
	}
	_ = collection.FindOne(context.TODO(), filter).Decode(&resultUser)
	err = checkResult(resultUser, loginForm)
	if err != nil {
		return "", err
	}
	return resultUser.UserID.Hex(), nil
}

func checkResult(res User, loginForm LoginForm) (err error) {
	if res.Email == "" {
		return errors.New("wrong email address")
	}
	err = CheckPasswd(res.Password, loginForm.Password)
	return err
}