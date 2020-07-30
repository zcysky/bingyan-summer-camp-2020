package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var userColl *mongo.Collection
var commodityColl *mongo.Collection
var keywordColl *mongo.Collection

func SetupDatabase() (err error) {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err = mongo.NewClient(clientOptions)
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	userColl = client.Database("mall").Collection("users")
	commodityColl = client.Database("mall").Collection("commodities")
	keywordColl = client.Database("mall").Collection("keywords")
	log.Println("Database connected successfully!")
	return err
}

func VerifyLogin(form LoginForm) (valid bool) {
	filter := bson.M{"username": form.Username}
	var res User
	userColl.FindOne(context.TODO(), filter).Decode(&res)
	if res.Username != "" && Compare(res.Password, form.Password) {
		return true
	}
	return false
}

func Signup(form User) (err error) {
	form.Password = Encrypt(form.Password)
	_, err = userColl.InsertOne(context.TODO(), form)
	return err
}

func QueryOneUser(username string) (form User, cnt Counter, err error) {
	filter := bson.M{"username": username}
	var res User
	err = userColl.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		return res, cnt, err
	}
	filter = bson.M{"publisher": username}
	cur, err := commodityColl.Find(context.TODO(), filter)
	if cur != nil {
		var commodity Commodity
		for cur.Next(context.TODO()) {
			cur.Decode(&commodity)
			cnt.ViewCnt += commodity.View
			cnt.CollectCnt += commodity.Collect
		}
	}
	return res, cnt, err
}

func UpdateUser(username string, updateForm UpdateForm) (err error) {
	filter := bson.M{"username": username}
	update := bson.M{"$set": updateForm}
	_, err = userColl.UpdateOne(context.TODO(), filter, update)
	return err
}

func GetCommodities(req CommodityRequest) (commodities []Commodity) {
	var filter bson.M
	var filter2 []bson.M
	keywords := strings.Split(req.Keyword, " ")
	for _, k := range keywords {
		filter2 = append(filter2, bson.M{"title": bson.M{"$regex": k}})
	}
	if req.Category != 0 {
		filter = bson.M{"category": req.Category, "$or": filter2}
	} else {
		filter = bson.M{"title": filter2}
	}
	var opts *options.FindOptions
	opts.SetLimit(int64(req.Limit))
	opts.SetSkip(int64((req.Page - 1) * req.Limit))
	opts.SetSort(bson.M{"view": -1})
	cur, _ := commodityColl.Find(context.TODO(), filter)
	if cur != nil {
		cur.All(context.TODO(), &commodities)
	}

	//Save keywords
	for _, k := range keywords {
		filter = bson.M{"key": k}
		res := keywordColl.FindOne(context.TODO(), filter)
		if res == nil {
			keywordColl.InsertOne(context.TODO(), bson.M{"key": k, "value": 1})
			continue
		}
		update := bson.M{"$inc": bson.M{"value": 1}}
		keywordColl.UpdateOne(context.TODO(), filter, update)
	}
	return commodities
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
	filter := bson.M{"_id": ObjID}
	commodityColl.FindOne(context.TODO(), filter).Decode(&form)
	return form
}

func GetSelfCommodities(username string) (commodities []SingleData) {
	filter := bson.M{"publisher": username}
	var res struct {
		Collections []SingleData `bson:"collections"`
	}
	commodityColl.FindOne(context.TODO(), filter).Decode(&res)
	return res.Collections
}

func AddCommodity(form PublishRequest, username string) (err error) {
	commodity := Commodity{Title: form.Title, Description: form.Desc, Price: form.Price,
		Category: form.Category, Picture: form.Picture, Publisher: username,
		View: 0, Collect: 0}
	_, err = commodityColl.InsertOne(context.TODO(), commodity)
	return err
}

func DeleteCommodity(id string, username string) (err error) {
	ObjID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": ObjID}
	target := GetOneCommodity(id)
	if target.Publisher == username {
		commodityColl.DeleteOne(context.TODO(), filter)
		return nil
	}
	return fmt.Errorf("permission denied")
}

func GetSelfCollections(username string) (commodities []SingleData) {
	filter := bson.M{"username": username}
	var res struct {
		Collections []SingleData `bson:"collections"`
	}
	userColl.FindOne(context.TODO(), filter).Decode(&res)
	return res.Collections
}

func AddCollection(username string, id string) (err error) {
	filter := bson.M{"username": username}
	form := GetOneCommodity(id)
	commodity := SingleData{ID: id, Title: form.Title}
	update := bson.M{"$push": bson.M{"collections": commodity}}
	_, err = userColl.UpdateOne(context.TODO(), filter, update)
	return err
}

func DeleteCollection(username string, id string) (err error) {
	filter := bson.M{"username": username}
	update := bson.M{"$pull": bson.M{"collections": bson.M{"id": id}}}
	_, err = userColl.UpdateOne(context.TODO(), filter, update)
	return err
}

func AddViewCounter(id string) {
	ObjID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": ObjID}
	update := bson.M{"$inc": bson.M{"view": 1}}
	commodityColl.UpdateOne(context.TODO(), filter, update)
}

func AddCollectCounter(id string, amount int) {
	ObjID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": ObjID}
	update := bson.M{"$inc": bson.M{"collect": amount}}
	commodityColl.UpdateOne(context.TODO(), filter, update)
}

func CreateHistory(commodityID string, username string) {
	filter := bson.M{"username": username}
	update := bson.M{"$push": bson.M{"history": commodityID}}
	userColl.UpdateOne(context.TODO(), filter, update)
}
