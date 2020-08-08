package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetComments(commodityID string) (comments []Comment, err error) {
	ObjID, err := primitive.ObjectIDFromHex(commodityID)
	if err != nil {
		return comments, err
	}
	filter := bson.M{"under_commodity": ObjID, "reply_to": ""}
	cur, _ := commentColl.Find(context.TODO(), filter)
	if cur != nil {
		cur.All(context.TODO(), &comments)
	}
	return comments, nil
}

func AddComment(form CommentRequest) (err error) {
	comment, err := makeComment(form)
	if err != nil {
		return err
	}
	_, err = commentColl.InsertOne(context.TODO(), comment)
	return err
}

func QueryReplies(commentID string) (comments []Comment, err error) {
	ObjID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return comments, err
	}
	filter := bson.M{"reply_to": ObjID}
	cur, _ := commentColl.Find(context.TODO(), filter)
	if cur != nil {
		cur.All(context.TODO(), &comments)
	}
	return comments, nil
}
