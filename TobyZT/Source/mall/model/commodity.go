package model

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCommodities(req CommodityRequest) (commodities []Commodity, total int64) {
	filter := make(bson.M)
	var keyFilter []bson.M
	var keywords []string
	filter["deleted"] = false
	if req.Keyword != "" {
		keywords = strings.Split(req.Keyword, " ")
		for _, k := range keywords {
			keyFilter = append(keyFilter, bson.M{"title": bson.M{"$regex": k}})
		}
		filter["$or"] = keyFilter
	}
	if req.Category != 0 {
		filter["category"] = req.Category
	}

	opts := options.Find().SetLimit(int64(req.Limit))
	opts.SetSkip(int64(req.Page * req.Limit))
	opts.SetSort(bson.M{"view": -1})
	cur, _ := commodityColl.Find(context.TODO(), filter)
	total, _ = commodityColl.CountDocuments(context.TODO(), filter)
	if cur != nil {
		cur.All(context.TODO(), &commodities)
	}

	//Save keywords
	for _, k := range keywords {
		filter = bson.M{"key": k}
		res, _ := keywordColl.CountDocuments(context.TODO(),filter)
		if res == 0 {
			keywordColl.InsertOne(context.TODO(), bson.M{"key": k, "value": 1})
			continue
		}
		update := bson.M{"$inc": bson.M{"value": 1}}
		keywordColl.UpdateOne(context.TODO(), filter, update)
	}
	return commodities, total
}

func GetHots(limit int) (keywords []string) {
	opts := options.Find().SetSort(bson.M{"value": -1})
	cur, _ := keywordColl.Find(context.TODO(), bson.M{}, opts)
	if cur != nil {
		var res struct {
			key   string
			value int
		}
		i := 0
		for i < limit && cur.Next(context.TODO()) {
			cur.Decode(&res)
			keywords = append(keywords, res.key)
			i++
		}
	}
	return keywords
}

func GetOneCommodity(id string) (form Commodity) {
	ObjID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": ObjID, "deleted": false}
	commodityColl.FindOne(context.TODO(), filter).Decode(&form)
	return form
}

func GetSelfCommodities(username string) (commodities []SingleData) {
	filter := bson.M{"publisher": username, "deleted": false}
	var res struct {
		Collections []SingleData `bson:"collections"`
	}
	commodityColl.FindOne(context.TODO(), filter).Decode(&res)
	return res.Collections
}

func AddCommodity(form PublishRequest, username string) (err error) {
	objID := primitive.NewObjectID()
	commodity := Commodity{ID: objID, Title: form.Title, Description: form.Desc, Price: form.Price,
		Category: form.Category, Picture: form.Picture, Publisher: username,
		View: 0, Collect: 0}

	_, err = commodityColl.InsertOne(context.TODO(), commodity)
	return err
}

func DeleteCommodity(id string, username string) (err error) {
	ObjID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": ObjID, "deleted": false}
	target := GetOneCommodity(id)
	if target.Publisher == username {
		update := bson.M{"$set": bson.M{"deleted": true}}
		commodityColl.UpdateOne(context.TODO(), filter, update)
		return nil
	}
	return fmt.Errorf("permission denied")
}

func AddViewCounter(id string) {
	ObjID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": ObjID, "deleted": false}
	update := bson.M{"$inc": bson.M{"view": 1}}
	commodityColl.UpdateOne(context.TODO(), filter, update)
}

func AddCollectCounter(id string, amount int) {
	ObjID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": ObjID, "deleted": false}
	update := bson.M{"$inc": bson.M{"collect": amount}}
	commodityColl.UpdateOne(context.TODO(), filter, update)
}
