package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

//检查数据库中是否已经存了这个email
func checkEmail(email string, collectionName string) bool {
	//检查链接是否正常
	err := TestConnection()
	if err != nil {
		return false
	}

	collection := client.Database("users").Collection(collectionName)
	filter := bson.M{"email": email}

	var result User
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return false
	} else {
		return true
	}
}

//从所有表中查找
func CheckEmail(email string) bool {
	collections := []string{"users", "admins"}
	for _,name := range(collections) {
		exist := checkEmail(email, name)
		if exist == true {
			return true
		}
	}
	return false
}