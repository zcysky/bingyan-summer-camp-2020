package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const colNameEvent = "event"

var colEvent *mongo.Collection

func initModelEvent() {
	colEvent = MongoDatabase.Collection(colNameEvent)
}

type Event struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id"`
	User           string             `bson:"user" json:"user"`
	Desc           string             `bson:"desc" json:"desc"`
	Time           int64              `bson:"time" json:"time"`
	Remind         bool               `bson:"remind" json:"remind"`
	RemindTime     int64              `bson:"remind_time" json:"remind_time"`
	RemindInterval int                `bson:"remind_interval" json:"remind_interval"`
}

func GetEventByID(idHex string) (Event, bool, error) {
	var event Event

	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return event, false, err
	}

	err = colEvent.FindOne(context.Background(), bson.M{"_id": id}).Decode(&event)
	if err == mongo.ErrNoDocuments {
		return event, false, nil
	}
	if err != nil {
		return event, false, err
	}
	return event, true, nil
}

func GetEventsByCondition(condition bson.M) ([]Event, error) {
	var events []Event

	result, err := colEvent.Find(context.Background(), condition)
	if err != nil {
		return events, err
	}
	err = result.All(context.Background(), &events)
	return events, err
}

func AddEvent(event Event) (string, error) {
	event.ID = primitive.NewObjectID()
	_, err := colEvent.InsertOne(context.Background(), event)
	if err != nil {
		return "", err
	}
	return event.ID.Hex(), nil
}

func UpdateEvent(idHex string, update bson.M) error {
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return err
	}

	updateQuery := bson.M{
		"$set": update,
	}

	_, err = colEvent.UpdateOne(context.Background(), bson.M{"_id": id}, updateQuery)
	return err
}

func DeleteEvent(idHex string) error {
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		return err
	}

	_, err = colEvent.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
