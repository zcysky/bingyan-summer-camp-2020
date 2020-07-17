package module_mapper

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

// return the result
func FindMapper(name string, i interface{}, all bool) (interface{}, error) {

	if !all {
		var result User
		err := UserCol.FindOne(context.TODO(), bson.M{
			name: i,
		}).Decode(&result)
		if err != nil {
			return nil, err
		}

		return result, nil
	} else {
		var result []User

		cur, err := UserCol.Find(context.TODO(), bson.M{
			name: i,
		})
		if err != nil {
			return nil, err
		}

		if err = cur.All(context.TODO(), &result); err != nil {
			return nil, err
		}

		return result, nil
	}
}
