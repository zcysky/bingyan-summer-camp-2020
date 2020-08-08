package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

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
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Price       float32            `bson:"price"`
	Category    int                `bson:"category"`
	Picture     string             `bson:"picture"`
	Publisher   string             `bson:"publisher"`
	Deleted     bool               `bson:"deleted"`
	View        int                `bson:"view"`
	Collect     int                `bson:"collect"`
}

type SingleData struct {
	ID    string `bson:"id"`
	Title string `bson:"title"`
}

type CommodityRequest struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Category int    `json:"category"`
	Keyword  string `json:"keyword"`
}

type PublishRequest struct {
	Title    string  `json:"title"`
	Desc     string  `json:"desc"`
	Category int     `json:"category"`
	Price    float32 `json:"price"`
	Picture  string  `json:"picture"`
}

type Counter struct {
	ViewCnt    int
	CollectCnt int
}

type Comment struct {
	ID             primitive.ObjectID `bson:"_id"`
	Username       string             `bson:"username"`
	Content        string             `bson:"content"`
	Time           time.Time          `bson:"time"`
	UnderCommodity primitive.ObjectID `bson:"under_commodity"`
	ReplyTo        primitive.ObjectID `bson:"reply_to"`
}

type CommentRequest struct {
	Username       string `bson:"username"`
	Content        string `bson:"content"`
	UnderCommodity string `bson:"under_commodity"`
	ReplyTo        string `bson:"reply_to"`
}

func makeComment(form CommentRequest) (comment Comment, err error) {
	ObjID := primitive.NewObjectID()
	t := time.Now()
	replyTo, err := primitive.ObjectIDFromHex(form.ReplyTo)
	if err != nil {
		return comment, err
	}
	commodity, err := primitive.ObjectIDFromHex(form.UnderCommodity)
	if err != nil {
		return comment, err
	}
	comment = Comment{
		ID:             ObjID,
		Username:       form.Username,
		Content:        form.Content,
		Time:           t,
		UnderCommodity: commodity,
		ReplyTo:        replyTo,
	}
	return comment, nil
}
