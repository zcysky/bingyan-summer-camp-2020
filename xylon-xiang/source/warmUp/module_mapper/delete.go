package module_mapper

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

// delete by id
func Delete(id string) error {
	_, err := UserCol.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		return err
	}

	return nil
}
