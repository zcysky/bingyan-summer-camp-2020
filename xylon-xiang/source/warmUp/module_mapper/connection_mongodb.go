package module_mapper

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"warmUp/config"
)

var UserCol *mongo.Collection

func UserDB() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Config.DataBase.DatabaseAddress))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	UserCol = client.Database(config.Config.DataBase.DatabaseName).
		Collection(config.Config.DataBase.CollectionUserName)

	return nil
}
