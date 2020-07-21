/* This file contains a series of functions to interact with database. */
package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	if res.Email != "" && Compare(res.Password, form.Password) {
		return true, true, res.UserID.Hex(), nil
	}

	col = client.Database("users").Collection("users")
	filter = bson.M{"email": form.Email}
	col.FindOne(context.TODO(), filter).Decode(&res)
	if res.Email == "" {
		return false, false, "", nil
	}
	if !Compare(res.Password, form.Password) {
		return false, false, "", err
	}
	return true, false, res.UserID.Hex(), nil
}

// SignupNew accepts a SignupForm and detects if the account already exists
// If not, it sign up a new account in the database
func SignupNew(form SignupForm) (id string, err error) {
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	col := client.Database("users").Collection("users")
	form.Password = Encrypt(form.Password)
	ObjID, err := col.InsertOne(context.TODO(), form)
	if ObjID != nil {
		id = ObjID.InsertedID.(primitive.ObjectID).Hex()
	}

	if err != nil {
		return "", err
	}
	return "", nil
}

func AccountExist(email string) (exist bool, err error) {
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
		return false, err
	}

	col := client.Database("users").Collection("users")
	filter := bson.M{"email": email}
	var res User
	col.FindOne(context.TODO(), filter).Decode(&res)
	if res.Email != "" {
		return true, nil
	}
	return false, nil
}

// Update aids to update info in the database
func Update(newForm SignupForm, id string) (err error) {
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
		return err
	}

	col := client.Database("users").Collection("users")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objID}
	_, err = col.UpdateOne(context.TODO(), filter, bson.M{"$set": newForm})
	return err
}

// Delete accepts an userID and delete relative record in the database
func Delete(id string) (err error) {
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
		return err
	}

	col := client.Database("users").Collection("users")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objID}
	_, err = col.DeleteOne(context.TODO(), filter)
	return err
}

func QueryOne(id string) (user User, err error) {
	col := client.Database("users").Collection("users")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	filter := bson.M{"_id": objID}
	err = col.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func QueryAll(limit int, page int) (users []User, err error) {
	col := client.Database("users").Collection("users")
	cur, err := col.Find(context.TODO(), nil)
	if err != nil || cur == nil {
		return users, err
	}

	for cur.Next(context.TODO()) {
		var user User
		err := cur.Decode(&user)
		if err != nil {
			return []User{}, err
		}
		users = append(users, user)
	}
	// -------- Not complete yet ---------
	return users, nil
}
