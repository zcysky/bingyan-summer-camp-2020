package module

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"iElective/config"
	"iElective/controller"
	"log"
	"reflect"
	"time"
)

var (
	UserCol          *mongo.Collection
	CourseCol        *mongo.Collection
	CourseCommentCol *mongo.Collection
	UserCommentCol   *mongo.Collection
)

const (
	USER          = "user"
	COURSE        = "course"
	COURSECOMMENT = "course_comment"
	USERCOMMENT   = "user_comment"
)

func init() {

	client, err := mongo.NewClient(options.Client().ApplyURI(config.Config.Mongo.DBAddress))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	dataBase := client.Database(config.Config.Mongo.DBName)
	UserCol = dataBase.Collection(config.Config.Mongo.UserCollection)
	CourseCol = dataBase.Collection(config.Config.Mongo.CourseCollection)
	CourseCommentCol = dataBase.Collection(config.Config.Mongo.CourseCommentCollection)

}

func Insert(colName string, colValue interface{}) error {

	var err error

	switch colName {
	case USER:
		_, err = UserCol.InsertOne(context.Background(), bson.M{
			USER: colValue.(User),
		})

	case COURSE:
		_, err = CourseCol.InsertOne(context.Background(), bson.M{
			COURSE: colValue.(Course),
		})

	case COURSECOMMENT:
		_, err = CourseCommentCol.InsertOne(context.Background(), bson.M{
			COURSECOMMENT: colValue.(CourseComment),
		})

	case USERCOMMENT:
		_, err = UserCommentCol.InsertOne(context.Background(), bson.M{
			USERCOMMENT: colValue.(UserComment),
		})
	}

	if err != nil {
		return err
	}

	return nil
}

// targetId should order the rule that the first is the userId, and the second is the targetId
func Delete(colName string, id string) error {
	var err error

	switch colName {
	case USER:
		_, err = UserCol.DeleteOne(context.Background(), bson.M{
			"user.user_id": id,
		})

	case COURSE:
		_, err = CourseCol.DeleteOne(context.Background(), bson.M{
			"course.course_id": id,
		})

	case COURSECOMMENT:
		_, err = CourseCommentCol.DeleteOne(context.Background(), bson.M{
			"course_comment.comment_id": id,
		})

	case USERCOMMENT:
		_, err = UserCommentCol.DeleteOne(context.Background(), bson.M{
			"user_comment.comment_id": id,
		})
	}

	if err != nil {
		return err
	}

	return nil
}

func Update(colName string, colValue interface{}) error {

	var err error

	switch colName {
	case USER:
		_, err = UserCol.UpdateOne(context.Background(),
			bson.M{"user.user_id": colValue.(User).UserId},
			bson.M{"user": colValue.(User)})

	case COURSE:
		_, err = CourseCol.UpdateOne(context.Background(),
			bson.M{"course.course_id": colValue.(Course).CourseId},
			bson.M{"$set": bson.M{"course": colValue.(Course)}})

	case COURSECOMMENT:
		_, err = CourseCommentCol.UpdateOne(context.Background(),
			bson.M{
				"course_comment.comment_id": colValue.(CourseComment).CommentId,
			},
			bson.M{
				"$set": bson.M{"course_comment": colValue.(CourseComment)},
			})

	case USERCOMMENT:
		_, err = UserCommentCol.UpdateOne(context.Background(),
			bson.M{
				"user_comment.comment_id": colValue.(UserComment).CommentId,
			},
			bson.M{
				"$set": bson.M{"user_comment": colValue.(UserComment)},
			})
	}

	if err != nil {
		return err
	}

	return nil
}

// it can just select the course
func SelectCourseClass(selectObj controller.SelectObject) ([]Course, error) {

	var (
		err     error
		results []Course
	)

	filters, err := bson.Marshal(bson.M{})
	if err != nil {
		return nil, err
	}

	t := reflect.TypeOf(selectObj)
	v := reflect.ValueOf(selectObj)
	for i := 0; i <= v.NumField(); i++ {
		if v.Field(i).IsZero() {
			filter := bson.M{t.Field(i).Name: v.Field(i)}
			filters, err = bson.MarshalAppend(filters, filter)
			if err != nil {
				return nil, err
			}
		}
	}

	cursor, err := CourseCol.Find(context.Background(), filters)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func SelectSpecificCourse(courseId string) (Course, error) {

	var result Course

	err := CourseCol.FindOne(context.Background(), bson.M{
		"course.course_id": courseId,
	}).Decode(&result)
	if err != nil {
		return Course{}, err
	}

	return result, nil
}

// if the colName is COURSECOMMENT, then the interface will be the []CourseComment
// if the colName is USERCOMMNET, then the interface will be the []UserComment
func SelectComment(colName string, colValue string) (interface{}, error) {

	switch colName {
	case COURSECOMMENT:
		var results []CourseComment

		cursor, err := CourseCommentCol.Find(context.Background(), bson.M{"course_comment.user_id": colValue})
		if err != nil {
			return nil, err
		}

		if err = cursor.All(context.Background(), &results); err != nil {
			return nil, err
		}

		return results, nil

	case USERCOMMENT:
		var results []UserComment

		cursor, err := UserCommentCol.Find(context.Background(), bson.M{"user_comment.target_id": colValue})
		if err != nil {
			return nil, err
		}

		if err = cursor.All(context.Background(), &results); err != nil {
			return nil, err
		}

		return results, nil
	}

	err := errors.New("wrong col name error")
	return nil, err
}
