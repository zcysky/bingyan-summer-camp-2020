package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//收藏信息存储表单
type CollectionForm struct {
	ID		string	`json:"id"`
	Title 	string	`json:"title"`
}

//向数据库中添加用户的收藏信息
func AddCollection(collectionForm CollectionForm, userName string) (err error) {
	//获取ID对应的商品名
	objectID, err := primitive.ObjectIDFromHex(collectionForm.ID)
	if err != nil {
		return err
	}
	commodityFilter := bson.M{
		"_id":	objectID,
	}
	var result Commodity
	err = commoditiesCol.FindOne(context.TODO(), commodityFilter).Decode(&result)
	if err != nil {
		return err
	}
	collectionForm.Title = result.Title

	//修改商品的收藏信息
	commodityUpdate := bson.M{
		"$inc":	bson.M{
			"collect":	1,
		},
	}
	_, err = commoditiesCol.UpdateOne(context.TODO(), commodityFilter, commodityUpdate)
	if err != nil {
		return err
	}

	//向用户的收藏栏中添加收藏信息
	collectionFilter := bson.M{
		"username":	userName,
	}
	collectionUpdate := bson.M{
		"$push": bson.M{
			"collections":	collectionForm.ID,
		},
	}
	_, err = usersCol.UpdateOne(context.TODO(), collectionFilter, collectionUpdate)
	if err != nil {
		return err
	}

	return nil
}

//从数据库中获取用户的收藏信息
func GetUserCollection(userName string) (collections []CollectionForm, err error) {
	collectionsFilter := bson.M{
		"username":	userName,
	}
	err = usersCol.FindOne(context.TODO(), collectionsFilter).Decode(&collections)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

//从数据库中删除用户的收藏
func DeleteCollection(collectionForm CollectionForm, userName string) (err error) {
	//收藏信息
	collectionFilter := bson.M{
		"username":	userName,
	}
	collectionUpdate := bson.M{
		"$pull": bson.M{
			"collections": bson.M{
				"id": collectionForm.ID,
			},
		},
	}

	//被收藏商品信息
	objectID, err := primitive.ObjectIDFromHex(collectionForm.ID)
	commodityFilter := bson.M{
		"_id":	objectID,
	}
	commodityUpdate := bson.M{
		"$inc":	bson.M{
			"collect":	1,
		},
	}

	//修改用户的收藏信息
	_, err = usersCol.UpdateOne(context.TODO(),collectionFilter, collectionUpdate)
	if err != nil {
		return err
	}

	//修改商品的被收藏信息
	_, err = commoditiesCol.UpdateOne(context.TODO(), commodityFilter, commodityUpdate)
	if err != nil {
		return err
	}

	return err
}