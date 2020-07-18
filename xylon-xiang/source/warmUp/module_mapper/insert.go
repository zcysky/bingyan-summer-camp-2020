package module_mapper

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertMapper(user User) error {
	_, err := UserCol.InsertOne(context.TODO(), bson.D{
		{"id", user.ID},
		{"name", user.Name},
		{"password", user.Password},
		{"email", user.Email},
		{"phone", user.Phone},
	})

	if err != nil{
		return err
	}

	return nil
}
