package module_mapper

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

// return the result
func FindMapper(name string, i string, all bool) (interface{}, error) {

	if !all {
		var result User
		err := UserCol.FindOne(context.TODO(), bson.D{
			{name, i},
		}).Decode(&result)
		if err != nil {
			return nil, err
		}

		return result, nil
	} else {
		var result []User
		cur, err := UserCol.Find(context.TODO(), bson.D{})
		if err != nil {
			return nil, err
		}
		defer cur.Close(context.TODO())

		for cur.Next(context.TODO()) {
			var element User

			err := cur.Decode(&element)
			if err != nil {
				return nil, err
			}

			result = append(result, element)
		}

		return result, nil
	}
}
