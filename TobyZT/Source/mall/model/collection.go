package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSelfCollections(username string) (commodities []SingleData) {
	filter := bson.M{"username": username}
	var res struct {
		Collections []SingleData `bson:"collections"`
	}
	userColl.FindOne(context.TODO(), filter).Decode(&res)
	return res.Collections
}

func AddCollection(username string, id string) (err error) {
	filter := bson.M{"username": username}
	form := GetOneCommodity(id)
	commodity := SingleData{ID: id, Title: form.Title}
	update := bson.M{"$push": bson.M{"collections": commodity}}
	_, err = userColl.UpdateOne(context.TODO(), filter, update)
	return err
}

func DeleteCollection(username string, id string) (err error) {
	filter := bson.M{"username": username}
	update := bson.M{"$pull": bson.M{"collections": bson.M{"id": id}}}
	_, err = userColl.UpdateOne(context.TODO(), filter, update)
	return err
}

