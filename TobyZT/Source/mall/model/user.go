package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func VerifyLogin(form LoginForm) (valid bool) {
	filter := bson.M{"username": form.Username}
	var res User
	userColl.FindOne(context.TODO(), filter).Decode(&res)
	if res.Username != "" && Compare(res.Password, form.Password) {
		return true
	}
	return false
}

func Signup(form User) (err error) {
	form.Password = Encrypt(form.Password)
	_, err = userColl.InsertOne(context.TODO(), form)
	return err
}

func QueryOneUser(username string) (form User, cnt Counter, err error) {
	filter := bson.M{"username": username}
	var res User
	err = userColl.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		return res, cnt, err
	}
	filter = bson.M{"publisher": username}
	cur, err := commodityColl.Find(context.TODO(), filter)
	if cur != nil {
		var commodity Commodity
		for cur.Next(context.TODO()) {
			cur.Decode(&commodity)
			cnt.ViewCnt += commodity.View
			cnt.CollectCnt += commodity.Collect
		}
	}
	return res, cnt, err
}

func UpdateUser(username string, updateForm UpdateForm) (err error) {
	filter := bson.M{"username": username}
	update := bson.M{"$set": updateForm}
	_, err = userColl.UpdateOne(context.TODO(), filter, update)
	return err
}

func CreateHistory(commodityID string, username string) {
	filter := bson.M{"username": username}
	update := bson.M{"$push": bson.M{"history": commodityID}}
	userColl.UpdateOne(context.TODO(), filter, update)
}

