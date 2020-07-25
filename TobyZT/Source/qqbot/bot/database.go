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
	col := client.Database("reminder").Collection("reminder1")

	filter := bson.M{"due": bson.M{"$lt": time.Now().Add(-1 * time.Minute).Unix()}}
	col.DeleteMany(context.TODO(), filter)

	opt := options.Find().SetSort(bson.M{"due": 1})
	cur, err := col.Find(context.TODO(), bson.M{}, opt)
	if cur == nil || err != nil {
		return reminders, err
	}
	for cur.Next(context.TODO()) {
		var r Reminder
		err = cur.Decode(&r)
		// upcoming reminder
		if r.Due <= time.Now().Add(1*time.Minute).Unix() {
			reminders = append(reminders, r)
			continue
		}
		// remind in advance
		if time.Now().Add(time.Duration(r.Advance)*time.Minute).Unix() >= r.Due {
			if int((r.Due-time.Now().Unix())/60)%r.Gap == 0 {
				reminders = append(reminders, r)
			}
		}
	}
	return reminders, err
}

func QueryByID(id int) (reminders []Reminder, err error) {
	col := client.Database("reminder").Collection("reminder1")

	filter := bson.M{"id": id}

	opt := options.Find().SetSort(bson.M{"due": 1})
	cur, err := col.Find(context.TODO(), filter, opt)
	if cur != nil {
		err = cur.All(context.TODO(), &reminders)
	}
	return reminders, err
}
