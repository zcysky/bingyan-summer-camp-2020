package model

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

//用户登录部分

//用户登录表单
type LoginForm struct {
	UserName	string	`json:"username"`
	Password 	string	`json:"password"`
}

//根据用户名搜索用户信息并比对密码
func VerifyLogin(loginForm LoginForm) (err error) {
	//搜索条件
	loginFilter := bson.M{"username":	loginForm.UserName}
	var result RegisterForm

	//搜索不到则用户名不正确
	err = usersCol.FindOne(context.TODO(), loginFilter).Decode(&result)
	if err != nil {
		return errors.New("wrong username")
	}

	//搜索到但密码不匹配则密码不正确
	err = Compare(result.Password, loginForm.Password)
	if err != nil {
		return errors.New("wrong password")
	}

	return nil
}

//用户注册部分

//用户注册信息表单
type RegisterForm struct {
	UserName string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	NickName string `json:"nickname" bson:"nickname"`
	Mobile   string `json:"mobile" bson:"mobile"`
	Email    string `json:"email" bson:"email"`
}

//将用户信息加密后存放到数据库中
func Register(registerForm RegisterForm) (err error) {
	//密码加密
	registerForm.Password, err = Encrypt(registerForm.Password)
	if err != nil {
		return err
	}

	//用户信息填入
	_, err = usersCol.InsertOne(context.TODO(), registerForm)
	return err
}

//检查用户名是否已经注册
func CheckExist(registerForm RegisterForm) (err error) {
	log.Println("Checking whether the account is existed")

	//检测连接是否正常
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	var result RegisterForm

	//在表中寻找有无相同用户名
	userNameFilter := bson.M{"username":	registerForm.UserName}
	err = usersCol.FindOne(context.TODO(), userNameFilter).Decode(&result)
	if err == nil {
		err = errors.New("username has been registered")
		return err
	}

	return nil
}

//用户更新信息部分

//用户更新信息表单
type UpdateForm struct {
	Password string `bson:"password"`
	Nickname string `bson:"nickname"`
	Mobile   string `bson:"mobile"`
	Email    string `bson:"email"`
}

//合并用户更新的信息与原注册信息
func MergeUpdateInfo(update UpdateForm, register RegisterForm) (RegisterForm, error) {
	var err error

	if update.Password != "" {
		register.Password, err = Encrypt(update.Password)
		if err != nil {
			return RegisterForm{}, err
		}
	}

	register.NickName = update.Nickname
	register.Mobile = update.Mobile
	register.Email = update.Email

	return register, nil
}

//根据用户名，将用户新信息更新至数据库
func UpdateUserInfo(user RegisterForm) (err error) {
	updateFilter := bson.M{"username":	user.UserName}
	update := bson.M{"$set":	user}
	_, err = usersCol.UpdateOne(context.TODO(), updateFilter, update)
	return err
}

//用户获取信息部分

//根据用户名，获取用户注册信息和商品的浏览、收藏信息
func QueryUser(userName string) (user RegisterForm, cnt Cnt, err error) {
	log.Println("Querying user's information by user name")
	//从用户表中查询用户基本信息
	userNameFilter := bson.M{"username":	userName}
	err = usersCol.FindOne(context.TODO(), userNameFilter).Decode(&user)
	if err != nil {
		return RegisterForm{}, Cnt{}, nil
	}

	//从商品表中查询其发布商品的浏览量、收藏量
	cntFilter := bson.M{"publisher":	userName}
	resultCommodities, err := commoditiesCol.Find(context.TODO(), cntFilter)
	if err != nil {
		return RegisterForm{}, Cnt{}, nil
	}

	//整合所发布商品的浏览量、收藏量
	var commodity Commodity
	for resultCommodities.Next(context.TODO()) {
		err = resultCommodities.Decode(&commodity)
		if err != nil {
			return RegisterForm{}, Cnt{}, nil
		}
		cnt.ViewCnt += commodity.View
		cnt.CollectCnt += commodity.Collect
	}

	return user, cnt, nil
}