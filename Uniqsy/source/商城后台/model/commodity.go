package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mall/config"
	"strings"
)

//数据库中商品的存储信息
type Commodity struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Price       float32            `bson:"price"`
	Category    int                `bson:"category"`
	Picture     string             `bson:"picture"`
	Publisher   string             `bson:"publisher"`
	View        int                `bson:"view"`
	Collect     int                `bson:"collect"`
}

//查询所有商品信息使用的表单
type GetCommoditiesForm struct {
	Page 		int		`json:"page"`
	Limit 		int		`json:"limit"`
	Category	int		`json:"category"`
	KeyWord		string	`json:"key_word"`
}

//根据条件获取商品
func GetCommoditiesInfo(getComForm GetCommoditiesForm) (commodities []Commodity, err error) {
	//设置关键词搜索
	var keyWordsFliter []bson.M
	keyWords := strings.Split(getComForm.KeyWord, " ")
	for _, keyWord := range keyWords {
		keyWordsFliter = append(keyWordsFliter, bson.M{
			"title":	bson.M{
				"$regex":	keyWord,
			},
		})
	}

	//可能商品分类为0，即搜索全部商品
	var commoditiesFilter bson.M
	if getComForm.Category == 0 {
		commoditiesFilter = bson.M{
			"$or":	keyWordsFliter,
		}
	} else {
		commoditiesFilter = bson.M{
			"category":	getComForm.Category,
			"$or":		keyWordsFliter,
		}
	}

	//根据limit和page设置查找限制
	var findOptions *options.FindOptions
	findOptions.SetLimit(int64(getComForm.Limit))
	findOptions.SetSkip(int64((getComForm.Page - 1) * getComForm.Limit))
	findOptions.SetSort(bson.M{"view": -1})

	//查找在限制范围内的内容
	results, err := commoditiesCol.Find(context.TODO(), commoditiesFilter, findOptions)
	if err != nil {
		return nil, err
	}

	//将查找到的符合条件的内容转存
	err = results.All(context.TODO(), &commodities)
	if err != nil {
		return nil, err
	}

	//保存本次查找的关键词
	err = SaveKeyWords(keyWords)
	if err != nil {
		return nil, err
	}

	return commodities, nil
}

//数据库中存储关键词使用的格式
type KeyWordsForm struct {
	KeyWord		string 	`bson:"keyword"`
	Count		int 	`bson:"count"`
}

//生成一个新的关键词
func NewKeyWord(keyWord string) *KeyWordsForm {
	newKeyWord := &KeyWordsForm{
		KeyWord:	keyWord,
		Count: 		1,
	}
	return newKeyWord
}

//保存关键词
func SaveKeyWords(keyWords []string) (err error) {
	//对每一个关键词进行更新
	var keyWordFilter bson.M
	for _, keyWord := range keyWords {
		keyWordFilter = bson.M{
			"keyword":	keyWord,
		}

		//先找是否已经存在
		result := keyWordsCol.FindOne(context.TODO(), keyWordFilter)
		if result == nil { //不存在则新增记录
			_, err = keyWordsCol.InsertOne(context.TODO(), NewKeyWord(keyWord))
			if err != nil {
				return err
			}
		} else { //存在则记录加一
			keyWordUpdate := bson.M{
				"$inc":	bson.M{
					"count":	1,
				},
			}
			_, err = keyWordsCol.UpdateOne(context.TODO(),
				keyWordFilter, keyWordUpdate)
		}
	}
	return nil
}

//发布商品使用的表单
type PostCommodityForm struct {
	Title    string  `json:"title"`
	Desc     string  `json:"desc"`
	Category int     `json:"category"`
	Price    float32 `json:"price"`
	Picture  string  `json:"picture"`
}

//向数据库中添加新增的商品信息
func AddNewCommodity(postInfo PostCommodityForm, user string) (err error) {
	commodity := newCommodity(postInfo, user)
	_, err = commoditiesCol.InsertOne(context.TODO(), commodity)
	return err
}

//根据新增信息，补全生成完整商品信息
func newCommodity(postInfo PostCommodityForm, publisher string) *Commodity {
	newCommodity := &Commodity{
		Title:			postInfo.Title,
		Description:	postInfo.Desc,
		Category: 		postInfo.Category,
		Price:  		postInfo.Price,
		Picture:  		postInfo.Picture,
		Publisher:  	publisher,
		View: 			0,
		Collect:  		0,
	}
	return newCommodity
}

//获取当前查询量最高的10个关键词
func GetKeyWords() (keyWords []string, err error) {
	//从数据库中根据热度排序获取keyword
	findOptions := options.Find()
	findOptions.SetSort(bson.M{"count":	-1})
	result, err := keyWordsCol.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	//将查询结果中的关键词提取到字符串组
	var keyWordInfo KeyWordsForm
	cnt := 0
	for result.Next(context.TODO()) {
		err = result.Decode(&keyWordInfo)
		if err != nil {
			return nil, err
		}

		//限制最大数量
		keyWords = append(keyWords, keyWordInfo.KeyWord)
		cnt++
		if cnt > config.Config.KeyWord.Limit {
			break
		}
	}

	return keyWords, nil
}

//根据商品ID在数据库中查询商品信息
func GetCommodityInfo(commodityID string) (commodity Commodity, err error) {
	//准备商品搜索条件以及更新内容
	objectID, err := primitive.ObjectIDFromHex(commodityID)
	commodityFilter := bson.M{
		"_id":	objectID,
	}
	commodityUpdate := bson.M{
		"$inc":	bson.M{
			"view":	1,
		},
	}

	//根据商品ID添加一次商品的浏览记录
	_, err = commoditiesCol.UpdateOne(context.TODO(), commodityFilter, commodityUpdate)
	if err != nil {
		return Commodity{}, err
	}

	//获取商品信息
	err = commoditiesCol.FindOne(context.TODO(), commodityFilter).Decode(&commodity)
	if err != nil {
		return Commodity{}, err
	}

	return commodity, nil
}

//根据商品ID在数据库中删除商品
func DeleteCommodity(commodityID string, userName string) (err error) {
	//获取商品的ObjectID
	objectID, err := primitive.ObjectIDFromHex(commodityID)
	if err != nil {
		return err
	}
	commodityFilter := bson.M{
		"_id":			objectID,
		"publisher":	userName,
	}

	//删除商品信息
	_, err = commoditiesCol.DeleteOne(context.TODO(), commodityFilter)
	return err
}

//获取本人发布的商品的信息
func GetMyCommodities(userName string) (commodities []Commodity, err error) {
	//从数据库中获取信息
	commodityFilter := bson.M{
		"publisher":	userName,
	}
	result, err := commoditiesCol.Find(context.TODO(), commodityFilter)
	if err != nil {
		return []Commodity{}, err
	}

	//将信息转存
	err = result.All(context.TODO(), &commodities)
	if err != nil {
		return []Commodity{}, err
	}

	return commodities, nil
}