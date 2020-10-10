package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

//从请求信息中获取的注册表单
type RawRegisterForm struct {
	IsAdmin		bool	`json:"is_admin"`
	Invitation	string	`json:"invitation"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	Tel			string	`json:"tel"`
	Email		string	`json:"email"`
}

//精简后的用于存放在数据库中的注册表单
type DBRegisterForm struct {
	Username	string	`json:"username" bson:"username"`
	Password	string	`json:"password" bson:"password"`
	Tel			string	`json:"tel" bson:"tel"`
	Email		string	`json:"email" bson:"email"`
}

//将原始信息形式的注册表单转换为数据库形式的注册表单
func GetDBRegisterForm(rawForm RawRegisterForm) DBRegisterForm {
	dbForm := DBRegisterForm{
		Username: 	rawForm.Username,
		Password: 	rawForm.Password,
		Tel:		rawForm.Tel,
		Email:  	rawForm.Email,
	}
	return dbForm
}

func SaveInDB(form DBRegisterForm, colName string) (string, error) {
	//测试连接
	err := TestConnection()
	if err != nil {
		log.Println(err)
		return "", err
	}

	//加密用户密码
	form.Password = Encrypt(form.Password)

	//选择表
	collection := client.Database("users").Collection(colName)

	//添加信息
	objID, err := collection.InsertOne(context.TODO(), form)
	if err != nil {
		return "", err
	}

	//整理信息准备返回
	if objID != nil {
		id := objID.InsertedID.(primitive.ObjectID).Hex()
		return id, nil
	} else {
		return "", nil
	}
}