package controller

import (
	"encoding/json"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"icourses/config"
	"icourses/defination"
	"icourses/model"
	"icourses/util"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func GetWxToekn(code string)(defination.WxToken,error){

	Url,err:=url.Parse(config.Config.WxConfig.TokenApiAddress)
	if err!=nil{
		return defination.WxToken{},err
	}

	params:=url.Values{}
	params.Set("appid",config.Config.WxConfig.AppId)
	params.Set("secret",config.Config.WxConfig.Secret)
	params.Set("js_code",code)
	params.Set("grant_type",config.Config.WxConfig.GrantType)
	Url.RawQuery=params.Encode()
	urlPath:=Url.String()

	res,err:=http.Get(urlPath)
	defer res.Body.Close()
	if err!=nil{
		return defination.WxToken{},err
	}

	var wxToken defination.WxToken
	body,err:=ioutil.ReadAll(res.Body)
	if err!=nil{
		return defination.WxToken{},err
	}
	err=json.Unmarshal(body,&wxToken)
	if err!=nil{
		return defination.WxToken{},err
	}
	return wxToken,err
}

func HandleGetToken(ctx echo.Context)error{
	code:=ctx.Param("code")
	wxToken,err:=GetWxToekn(code)
	if err!=nil{
		return ctx.String(http.StatusInternalServerError,"无法进行认证")
	}
	if wxToken.Errcode==-1{
		return ctx.String(http.StatusInternalServerError,"微信系统繁忙")
	}
	if wxToken.Errcode==40029{
		return ctx.String(http.StatusInternalServerError,"无效code")
	}
	if wxToken.Errcode==45011{
		return ctx.String(http.StatusInternalServerError,"频率限制")
	}
	if wxToken.Errcode!=0{
		return ctx.String(http.StatusInternalServerError,"微信api错误")
	}
	userInfo,err:=model.FindUserWithOpenid(wxToken.Openid)
	if err==mongo.ErrNoDocuments{
		var newUser defination.User
		newUser.Key=wxToken.Key
		newUser.OpenId=wxToken.Openid
		newUser.Type="general"
		config.Config.MongoConfig.Uid++
		newUser.Uid=config.Config.MongoConfig.Uid
		err:=model.InsertUser(newUser)
		if err!=nil{
			return ctx.String(http.StatusInternalServerError,"无法创建新用户")
		}
		userInfo=newUser
	}else if err!=nil{
		return ctx.String(http.StatusInternalServerError,"无法创建新用户")
	}else {
		err:=model.UpdateUserKey(wxToken.Openid,wxToken.Key)
		if err!=nil{
			return ctx.String(http.StatusInternalServerError,"无法更新用户session_key")
		}
		userInfo.Key=wxToken.Key
	}
	jwtToken,err:=util.SignJwtToken(userInfo)
	if err!=nil{
		return ctx.String(http.StatusInternalServerError,"无法生成jwt令牌")
	}
	return ctx.String(http.StatusOK,jwtToken)
}

func HandlePostUser(ctx echo.Context)error{
	uidStr:=ctx.Param("uid")
	uid,err:=strconv.Atoi(uidStr)
	if err!=nil{
		return ctx.String(http.StatusBadRequest,"参数格式错误")
	}
	var userInfo defination.User
	ctx.Bind(&userInfo)
	userInfo.Uid=uid
	err=model.UpdateUser(userInfo)
	if err!=nil{
		return ctx.String(http.StatusInternalServerError,"数据库更新错误")
	}
	return ctx.String(http.StatusOK,"succeed")
}

func HandleGetUserAllComments(ctx echo.Context)error{
	uidStr:=ctx.Param("uid")
	uid,err:=strconv.Atoi(uidStr)
	if err!=nil{
		return ctx.String(http.StatusBadRequest,"参数格式错误")
	}
	comments,err:=model.FindCommentWithUid(uid)
	if err!=nil{
		return ctx.String(http.StatusInternalServerError,"数据库查询错误")
	}
	commentsJSON,err:=json.Marshal(comments)
	if err!=nil{
		return ctx.String(http.StatusInternalServerError,"生成JSON数据错误")
	}
	return ctx.JSON(http.StatusOK,commentsJSON)
}