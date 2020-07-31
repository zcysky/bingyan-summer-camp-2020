package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)


func Mstart() *mongo.Client{
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
	//开始操作
}

func InsertUser(client *mongo.Client, u *UserInfoAll) error{
	collection := client.Database(Info.DataBase).Collection(Info.CollectionU)
	_, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		fmt.Println("插入数据库失败")
		return err
	}
	return nil
}

func CheckUser(client *mongo.Client, key string, value string) (bool, *UserInfoAll){
	collection := client.Database(Info.DataBase).Collection(Info.CollectionU)
	filter := bson.D{{key, value}}
	var result UserInfoAll
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	return true, &result
}

func CheckUserPass (client *mongo.Client, u *UserInfoAll) (bool, *UserInfoAll){
	collection := client.Database(Info.DataBase).Collection(Info.CollectionU)
	filter := bson.D{{"username", u.Username},{"password", u.Password}}
	var result UserInfoAll
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return false, nil
	}
	return true, &result
}

func UserView(client *mongo.Client, ID string, key string) bool {
	collection := client.Database(Info.DataBase).Collection(Info.CollectionU)
	filter := bson.D{{"id", ID}}
	update := bson.D{
		{"$inc", bson.D{
			{key, 1},
		}},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false
	}
	return true
}

func UserUpdate(client *mongo.Client, ID string, key string, value string) bool {
	collection := client.Database(Info.DataBase).Collection(Info.CollectionU)
	filter := bson.D{{"id", ID}}
	update := bson.D{
		{"$set", bson.D{
			{key, value},
		}},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false
	}
	return true
}

func UserDelete(client *mongo.Client, ID string) bool{
	collection := client.Database(Info.DataBase).Collection(Info.CollectionU)
	_, err := collection.DeleteOne(context.TODO(), bson.D{{"id",ID}})
	if err != nil {
		return false
	}
	return true
}

func InsertGood(client *mongo.Client, g *GoodInfoALL) error{
	collection := client.Database(Info.DataBase).Collection(Info.CollectionG)
	_, err := collection.InsertOne(context.TODO(), g)
	if err != nil {
		fmt.Println("插入数据库失败")
		return err
	}
	return nil
}

func CheckGood(client *mongo.Client, key string, value string) (bool, *GoodInfoALL){
	collection := client.Database(Info.DataBase).Collection(Info.CollectionG)
	filter := bson.D{{key, value}}
	var result GoodInfoALL
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	return true, &result
}

func GoodList(client *mongo.Client,r *ReqGoodsList) []*GoodInfoList{
	collection := client.Database(Info.DataBase).Collection(Info.CollectionG)
	findOptions := options.Find()
	var results []*GoodInfoList
	query:=bson.M{"title": bson.M{"$regex": r.Keyword, "$options": "$i"},"category":r.Category}
	cur, err := collection.Find(context.TODO(), query, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem GoodInfoList
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	return results
}

func MyGoodsList(client *mongo.Client, ID string) []*GoodInfoList{
	collection := client.Database(Info.DataBase).Collection(Info.CollectionG)
	findOptions := options.Find()
	var results []*GoodInfoList
	query:=bson.D{{"pubuser", ID}}
	cur, err := collection.Find(context.TODO(), query, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem GoodInfoList
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	return results
}

//func MyCollections(client *mongo.Client, ID string) []*GoodInfoSim {
//	collection := client.Database(Info.DataBase).Collection(Info.CollectionU)
//
//	filter := bson.D{{"id", ID}}
//	var results []*GoodInfoList
//	err := collection.FindOne(context.TODO(), filter).Decode(result)
//	if err != nil {
//		fmt.Println(err)
//		return false, nil
//	}
//}

func UpdateGood(client *mongo.Client, ID string, key string) bool {
	collection := client.Database(Info.DataBase).Collection(Info.CollectionG)
	filter := bson.D{{"id", ID}}
	update := bson.D{
		{"$inc", bson.D{
			{key, 1},
		}},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false
	}
	return true
}

func DeleteGood(client *mongo.Client, ID string) bool {
	collection := client.Database(Info.DataBase).Collection(Info.CollectionG)
	_, err := collection.DeleteOne(context.TODO(), bson.D{{"id",ID}})
	if err != nil {
		return false
	}
	return true
}

func CheckHot(client *mongo.Client, keyword string) (bool, *HotWord){
	collection := client.Database(Info.DataBase).Collection(Info.CollectionH)
	filter := bson.D{{"keyword", keyword}}
	var result HotWord
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	return true, &result
}

func InsertHot(client *mongo.Client, h *HotWord) error{
	collection := client.Database(Info.DataBase).Collection(Info.CollectionH)
	h.Order = 4
	h.Count = 0
	_, err := collection.InsertOne(context.TODO(), h)
	if err != nil {
		fmt.Println("插入数据库失败")
		return err
	}
	UpdateHot(client, h)
	return nil
}

func UpdateHot(client *mongo.Client, h *HotWord) int64{
	collection := client.Database(Info.DataBase).Collection(Info.CollectionH)
	filter := bson.D{{"keyword", h.Keyword}}
	update := bson.D{
		{"$inc", bson.D{
			{"count", 1},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return 0
	}
	r := new(HotWord)
	r1 := new(HotWord)
	r2 := new(HotWord)
	r3 := new(HotWord)
	err = collection.FindOne(context.TODO(), filter).Decode(r)
	if err != nil {
		fmt.Println(err)
	}
	err = collection.FindOne(context.TODO(), bson.D{{"order", 1}}).Decode(r1)
	if err != nil {
		fmt.Println(err)
	}
	err = collection.FindOne(context.TODO(), bson.D{{"order", 2}}).Decode(r2)
	if err != nil {
		fmt.Println(err)
	}
	err = collection.FindOne(context.TODO(), bson.D{{"order", 3}}).Decode(r3)
	if err != nil {
		fmt.Println(err)
	}

	if r.Count > r1.Count || r1.Count == 0{
		if r.Order != 1{
			if r.Order >= 4{
				update = bson.D{
					{"$set", bson.D{
						{"order", 4},
					}},
				}
				_, err = collection.UpdateOne(context.TODO(), bson.D{{"order", 3}}, update)
				if err != nil {
				}
			}
			if r.Order >= 3{
				update = bson.D{
					{"$set", bson.D{
						{"order", 3},
					}},
				}
				_, err = collection.UpdateOne(context.TODO(), bson.D{{"order", 2}}, update)
				if err != nil {
				}
			}
			if r.Order >= 2{
				update = bson.D{
					{"$set", bson.D{
						{"order", 2},
					}},
				}
				_, err = collection.UpdateOne(context.TODO(), bson.D{{"order", 1}}, update)
				if err != nil {
				}
			}
		}
		update = bson.D{
			{"$set", bson.D{
				{"order", 1},
			}},
		}
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
		}
	} else if r.Count > r2.Count || r2.Count == 0{
		if r.Order != 2 {
			if r.Order >= 4{
				update = bson.D{
					{"$set", bson.D{
						{"order", 4},
					}},
				}
				_, err = collection.UpdateOne(context.TODO(), bson.D{{"order", 3}}, update)
				if err != nil {
				}
			}
			if r.Order >= 3{
				update = bson.D{
					{"$set", bson.D{
						{"order", 3},
					}},
				}
				_, err = collection.UpdateOne(context.TODO(), bson.D{{"order", 2}}, update)
				if err != nil {
				}
			}
		}
		update = bson.D{
			{"$set", bson.D{
				{"order", 2},
			}},
		}
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
		}
	}else if r.Count >= r3.Count || r3.Count == 0{
		if r.Order >= 4 {
			update = bson.D{
				{"$set", bson.D{
					{"order", 4},
				}},
			}
			_, err = collection.UpdateOne(context.TODO(), bson.D{{"order", 3}}, update)
			if err != nil {
			}
		}
		update = bson.D{
			{"$set", bson.D{
				{"order", 3},
			}},
		}
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
		}
	}else{
		update = bson.D{
			{"$set", bson.D{
				{"order", 4},
			}},
		}
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
		}
	}

	return result.MatchedCount
}

func Hot(client *mongo.Client) []*HotWord{
	collection := client.Database(Info.DataBase).Collection(Info.CollectionH)
	findOptions := options.Find()
	var results []*HotWord

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem HotWord
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		if elem.Order < 4 {
			results = append(results, &elem)
		}
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	return results
}