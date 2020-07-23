package bot

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// SetupDatabase connects to local database
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

func Insert(reminder Reminder) (err error) {
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
		return err
	}

	col := client.Database("reminder").Collection("reminder1")
	_, err = col.InsertOne(context.TODO(), reminder)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Query deletes overdue reminders,
// then queries and returns upcoming(5 min) reminders
func Query() (reminders []Reminder, err error) {
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
		return reminders, err
	}

	col := client.Database("reminder").Collection("reminder1")

	filter := bson.M{"due": bson.M{"$lt": time.Now().Unix()}}
	col.DeleteMany(context.TODO(), filter)

	border := time.Now().Add(5 * time.Minute).Unix()
	filter = bson.M{"due": bson.M{"$lte": border}}
	opt := options.Find().SetSort(bson.M{"due": 1})
	cur, err := col.Find(context.TODO(), filter, opt)
	if cur != nil {
		err = cur.All(context.TODO(), &reminders)
	}
	return reminders, err
}
