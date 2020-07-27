package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Nickname string `json:"nickname" bson:"nickname"`
	Mobile   string `json:"mobile" bson:"mobile"`
	Email    string `json:"email" bson:"email"`
}

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateForm struct {
	Password string `bson:"password"`
	Nickname string `bson:"nickname"`
	Mobile   string `bson:"mobile"`
	Email    string `bson:"email"`
}

type JWTForm struct {
	Secret string `json:"secret"`
	Expire int    `json:"expire"`
}

type Response struct {
	success bool
	err     string
	data    interface{}
}

type Commodity struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	Price     float32            `bson:"price"`
	Category  int                `bson:"category"`
	Picture   string             `bson:"picture"`
	Publisher string             `bson:"publisher"`
}
