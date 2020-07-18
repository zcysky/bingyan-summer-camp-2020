/* This file contains a series of functions to interact with database */
package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// SetupDatabase connects to local database and
func SetupDatabase() (err error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err = mongo.NewClient(clientOptions)
	if err != nil {
		log.Println(err)
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Database connected successfully!")
	return err
}

// VerifyLogin accepts a LoginForm and verify its validity
func VerifyLogin(form LoginForm) (valid bool, admin bool, id string, err error) {
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
		return false, false, "", err
	}
	// first check if it is an admin account
	col := client.Database("users").Collection("admins")
	filter := bson.M{"email": form.Email}
	var res User
	col.FindOne(context.TODO(), filter).Decode(&res)
	if res.Email != "" && res.Password == Encrypt(form.Password) {
		return true, true, res.UserID.Hex(), nil
	}

	col = client.Database("users").Collection("users")
	filter = bson.M{"email": form.Email}
	col.FindOne(context.TODO(), filter).Decode(&res)
	if res.Email == "" {
		return false, false, "", nil
	}
	if res.Password != Encrypt(form.Password) {
		return false, false, "", err
	}
	return true, false, res.UserID.Hex(), nil
}

// SignupNew accepts a SignupForm and detects if the account already exists
// If not, it sign up a new account in the database
func SignupNew(form SignupForm) (exist bool, err error) {
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
		return false, err
	}
	col := client.Database("users").Collection("users")
	filter := bson.M{"email": form.Email}
	var res User
	err = col.FindOne(context.TODO(), filter).Decode(&res)
	if res.Email != "" {
		return true, nil
	}
	newUser := bson.M{
		"username": form.Username,
		"password": Encrypt(form.Password),
		"phone":    form.Phone,
		"email":    form.Email,
	}
	_, err = col.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return false, nil
}

// Update accepts a UpdateForm and detects if the account already exists
func Update(form SignupForm) (err error) {

}