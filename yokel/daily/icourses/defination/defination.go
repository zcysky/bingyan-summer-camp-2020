package defination

import "github.com/dgrijalva/jwt-go"

type User struct {
	Uid    int    `json:"uid"`
	OpenId string `json:"openid"`
	User   string `json:"user"`
	Type   string `json:"type"`
	Avatar string `json:"avatar"`
	Key    string `json:"key"`
}

//Icnt stands for incognito

type SubComment struct {
	Cid     int    `json:"cid"form:"cid"`
	Uid     int    `json:"uid" form:"uid"`
	User    string `json:"user" form:"user"`
	Icnt    bool   `json:"icnt" form:"icnt"`
	Avatar  string `json:"avatar" form:"avatar"`
	Time    int    `json:"time" form:"time"`
	Content string `json:"content" form:"content"`
}

type Comment struct {
	Cid       int          `json:"cid"`
	Uid       int          `json:"uid"`
	Lid       int          `json:"lid"`
	User      string       `json:"user"`
	Icnt      bool         `json:"icnt"`
	TermEnd   bool         `json:"term-end"`
	Avatar    string       `json:"avatar"`
	Time      int          `json:"time"`
	Content   string       `json:"content"`
	Like      int          `json:"like" form:"like"`
	SubCmtNum int          `json:"subcmt-num" `
	SubCmt    []SubComment `json:"subcmt"`
	Img       string       `json:"img"`
}

type CommentForm struct {
	Uid        int    `json:"uid" form:"uid"`
	Avatar     string `json:"avatar"form:"avatar"`
	User       string `json:"user" form:"user"`
	Atdc       string `json:"atdc" form:"atdc"`
	Time       int    `json:"time" form:"time"`
	Exam       string `json:"exam" form:"exam"`
	Icnt       bool   `json:"icnt" form:"icnt"`
	TermEnd    bool   `json:"term-end" form:"term-end"`
	Evaluation string `json:"evaluation" form:"evaluation"`
	Content    string `json:"content" form:"content"`
	Img        string `json:"img" form:"img"`
}

type Arrange struct {
	Time     int    `json:"time"`
	Location string `json:"location"`
}

type Evaluation struct {
	Good int `json:"good"`
	Soso int `json:"soso"`
	Bad  int `json:"bad"`
}

type Tag struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Lesson struct {
	Lid        int          `json:"lid"`
	Course     string       `json:"course"`
	Type       string       `json:"type"`
	Credit     int          `json:"credit"`
	Target     string       `json:"target"`
	Arrange    []Arrange    `json:"arrange"`
	Evaluation []Evaluation `json:"evaluation"`
	Tag        []Tag        `json:"tag"`
}

type LessonFilter struct {
	Type     []string `json:"type" query:"type"`
	Time     []string `json:"time" query:"type"`
	Position []string `json:"position" query:"position"`
	Atdc     []string `json:"atdc" query:"atdc"`
	Exam     []string `json:"exam" query:"exam"`
}

type WxToken struct {
	Openid  string `json:"openid"`
	Key     string `json:"session_key"`
	Errcode int    `json:"errcode"`
}

type JwtToken struct {
	Uid  int    `json:"uid"`
	User string `json:"user"`
	Type string `json:"type"`
	jwt.StandardClaims
}

type MongoConfig struct {
	DbAddress string `json:"db-address"`
	Db        string `json:"db"`
	Courses   string `json:"courses"`
	Comments  string `json:"comments"`
	Users     string `json:"users"`
	Uid       int    `json:"uid"`
	Lid       int    `json:"lid"`
	Cid       int    `json:"cid"`
}

type EchoConfig struct {
	EchoPort       string `json:"echo-port"`
	GetAllCourses  string `json:"get-all-courses"`
	GetCourse      string `json:"get-course"`
	GetAllComments string `json:"get-all-comments"`
	GetComment     string `json:"get-comment"`
	PostComment    string `json:"post-comment"`
	PostSubcmt     string `json:"post-subcmt"`
	GetToken       string `json:"get-token"`
	PostUser       string `json:"post-user"`
	GetUserAllComments string`json:"get-user-all-comments"`
}
type JwtConfig struct {
	Secret        string `json:"secret"`
	TokenDuration int    `json:"token-duration"`
}

type WxConfig struct {
	AppId           string `json:"app-id"`
	Secret          string `json:"secret"`
	GrantType       string `json:"grant-type"`
	TokenApiAddress string `json:"token-api-address"`
}

type ConfigObject struct {
	MongoConfig MongoConfig `json:"mongo-config"`
	EchoConfig  EchoConfig  `json:"echo-config"`
	JwtConfig   JwtConfig   `json:"jwt-config"`
	WxConfig    WxConfig    `json:"wx-config"`
}
