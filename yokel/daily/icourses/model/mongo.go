package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"icourses/config"
	"icourses/defination"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var MongoClient *mongo.Client
func ConnectDataBase()error{
	var err error
	MongoClient,err=mongo.NewClient(options.Client().ApplyURI(config.Config.MongoConfig.DbAddress))
	if err!=nil{
		return err
	}
	ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
	err=MongoClient.Connect(ctx)
	if err!=nil{
		return err
	}
	return nil
}


func ConvertLessonFilter(filter defination.LessonFilter)bson.M{
	var TypeFilter []bson.M
	var AtdcFilter []bson.M
	var ExamFilter []bson.M
	var PositionFilter []bson.M
	var TimeFilter []bson.M
	for _,v := range filter.Type{
		if v==""{continue}
		TypeFilter=append(TypeFilter,bson.M{"type":v})
	}
	for _,v := range filter.Atdc{
		if v==""{continue}
		AtdcFilter=append(AtdcFilter,bson.M{"type":"atdc","value":v})
	}
	for _,v := range filter.Exam{
		if v==""{continue}
		ExamFilter=append(ExamFilter,bson.M{"type":"exam","value":v})
	}
	for _,v := range filter.Position{
		if v==""{continue}
		PositionFilter=append(PositionFilter,bson.M{"position":v})
	}
	for _,v := range filter.Time{
		if v==""{continue}
		TimeFilter=append(TimeFilter,bson.M{"time":v})
	}
	//set filter as
	//type=TypeFilter  tag=AtdcFilter&&ExamFilter  arrange=Timefilter||postitionfilter
	ColFilter:=bson.M{
		"type": TypeFilter,
		"tag":  bson.M{"$and": []bson.M{bson.M{"$or": AtdcFilter}, bson.M{"$or": ExamFilter}}},
		"arrange":  bson.M{"or": []bson.M{bson.M{"$or": TimeFilter}, bson.M{"$or": PositionFilter}}},
	}
	return ColFilter
}

func FindAllCourseWithFilter(filter defination.LessonFilter)([]defination.Lesson,error){
	col:=MongoClient.Database(config.Config.MongoConfig.Db).Collection(config.Config.MongoConfig.Courses)
	ColFilter:=ConvertLessonFilter(filter)
	cur,err:=col.Find(context.TODO(),ColFilter)
	if err!=nil{
		return []defination.Lesson{},nil
	}
	var result []defination.Lesson
	for cur.Next(context.TODO()){
		var tmp defination.Lesson
		err=cur.Decode(&tmp)
		if err!=nil{
			return []defination.Lesson{},nil
		}
		result=append(result,tmp)
	}
	return result,nil
}

func FindCourseWithLid(lid int)(defination.Lesson,error){
	col:=MongoClient.Database(config.Config.MongoConfig.Db).Collection(config.Config.MongoConfig.Courses)
	Colfilter:=bson.M{"lid":lid}
	var result defination.Lesson
	err:=col.FindOne(context.TODO(),Colfilter).Decode(&result)
	if err!=nil{
		return defination.Lesson{},err
	}
	return result,nil
}

func FindCommentAll(lid int)([]defination.Comment,error){
	col:=MongoClient.Database(config.Config.MongoConfig.Db).Collection(config.Config.MongoConfig.Comments)
	Colfilter:=bson.M{"lid":lid}
	var result []defination.Comment
	cur,err:=col.Find(context.TODO(),Colfilter)
	if err !=nil{
		return []defination.Comment{},err
	}
	for cur.Next(context.TODO()){
		var tmp defination.Comment
		err:=cur.Decode(&tmp)
		if err !=nil{
			return []defination.Comment{},err
		}
		result=append(result,tmp)
	}
	return result,nil
}

func FindCommentWithCid(cid int)(defination.Comment,error){
	col:=MongoClient.Database(config.Config.MongoConfig.Db).Collection(config.Config.MongoConfig.Comments)
	ColFilter:=bson.M{"cid":cid}
	var result defination.Comment
	err:=col.FindOne(context.TODO(),ColFilter).Decode(&result)
	if err!=nil{
		return defination.Comment{},err
	}
	return result,nil
}

func FindCommentWithUid(uid int)([]defination.Comment,error){
	col:=MongoClient.Database(config.Config.MongoConfig.Db).Collection(config.Config.MongoConfig.Comments)
	ColFilter:=bson.M{
		"$or":[]bson.M{
			bson.M{"uid":uid},
			bson.M{"subcmt.uid":uid},
		},
	}
	var result []defination.Comment
	cur,err:=col.Find(context.TODO(),ColFilter)
	if err !=nil{
		return []defination.Comment{},err
	}
	for cur.Next(context.TODO()){
		var tmp defination.Comment
		err:=cur.Decode(&tmp)
		if err !=nil{
			return []defination.Comment{},err
		}
		result=append(result,tmp)
	}
	return result,nil
}

func InsertComment(comment defination.Comment)error{
	col:=MongoClient.Database(config.Config.MongoConfig.Db).Collection(config.Config.MongoConfig.Comments)
	_,err:=col.InsertOne(context.TODO(),comment)
	if err !=nil{
		return err
	}
	return nil
}

func InsertSubCmt(cid int,subCmt defination.SubComment)error{
	col:=MongoClient.Database(config.Config.MongoConfig.Db).Collection(config.Config.MongoConfig.Comments)
	result,err:=FindCommentWithCid(cid)
	if err!=nil{
		return err
	}
	result.SubCmt=append(result.SubCmt, subCmt)
	filter:=bson.M{"cid":cid}
	update:=bson.M{"$set":result}

	var updatedDocument bson.M
	err=col.FindOneAndUpdate(context.TODO(),filter,update).Decode(&updatedDocument)
	if err!=nil{
		return err
	}
	return nil
}

func FindUserWithOpenid(openid string)(defination.User,error){
	col:=MongoClient.Database(config.Config.MongoConfig.Db).Collection(config.Config.MongoConfig.Users)
	filter:=bson.M{"openid":openid}
	var result defination.User
	err:=col.FindOne(context.TODO(),filter).Decode(&result)
	if err!=nil{
		return defination.User{},err
	}
	return result,nil
}

func FindUserWithUid(Uid int)(defination.User,error){
	col:=MongoClient.Database(config.Config.MongoConfig.Db).Collection(config.Config.MongoConfig.Users)
	filter:=bson.M{"uid":Uid}
	var result defination.User
	err:=col.FindOne(context.TODO(),filter).Decode(&result)
	if err!=nil{
		return defination.User{},err
	}
	return result,nil
}

func InsertUser(user defination.User)error {
	col := MongoClient.Database(config.Config.MongoConfig.Db).Collection(config.Config.MongoConfig.Users)
	_, err := col.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(user defination.User)error{
	col := MongoClient.Database(config.Config.MongoConfig.Db).Collection(config.Config.MongoConfig.Users)
	filter:=bson.M{"uid":user.Uid}
	update:=bson.M{"$set":[]bson.M{bson.M{"user":user.User},bson.M{"avatar":user.Avatar}}}

	var updatedDocument bson.M
	err:=col.FindOneAndUpdate(context.TODO(),filter,update).Decode(&updatedDocument)
	if err!=nil{
		return err
	}
	return nil
}

func UpdateUserKey(openid string,key string)error{
	col := MongoClient.Database(config.Config.MongoConfig.Db).Collection(config.Config.MongoConfig.Users)
	filter:=bson.M{"openid":openid}
	update:=bson.M{"$set":bson.M{"key":key}}

	var updatedDocument bson.M
	err:=col.FindOneAndUpdate(context.TODO(),filter,update).Decode(&updatedDocument)
	if err!=nil{
		return err
	}
	return nil

}