package module_mapper

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateMapper(user User) {
	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{"id", user.ID}}

	var result User
	_ = UserCol.FindOne(context.TODO(), bson.D{
		{"id", user.ID},
	}).Decode(&result)

	//update := bson.D{{"$set", bson.D{}}
	update := bson.D{{"$set", bson.D{
		{"id", user.ID},
		{"name", user.Name},
		{"password", user.Password},
		{"phone", user.Phone},
		{"email", user.Email},
		{"admin", result.Admin},
	}}}

	UserCol.FindOneAndUpdate(context.TODO(), filter, update, opts)

}
