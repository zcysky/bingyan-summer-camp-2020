package controller

import (
	"encoding/base64"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"warmup-ref/config"
	"warmup-ref/model"
	"warmup-ref/util"
)

type paramUserGetToken struct {
	Username string `query:"username" validate:"required"`
	Password string `query:"password" validate:"required"`
}

type responseUserGetToken struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire_time"`
}

func UserGetToken(context echo.Context) error {
	var param paramUserGetToken
	if err := context.Bind(&param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}
	if err := context.Validate(param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}

	user, found, err := model.GetUserWithUsername(param.Username)
	if !found {
		return util.ErrorResponse(context, http.StatusBadRequest, "user not found")
	}
	if err != nil {
		return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
	}
	if user.Password != base64.StdEncoding.EncodeToString([]byte(param.Password)) {
		return util.ErrorResponse(context, http.StatusForbidden, "username and password don't match")
	}
	if !user.Verified {
		return util.ErrorResponse(context, http.StatusForbidden, "please verify your email")
	}

	token, expireTime, err := util.NewJWTToken(user.ID.Hex())
	if err != nil {
		return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
	}
	return util.SuccessResponse(context, http.StatusOK, responseUserGetToken{
		Token:  token,
		Expire: expireTime.Unix(),
	})
}

type paramUserRegister struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone" validate:"required,numeric"`
	Email    string `json:"email" validate:"required,email"`
}

type responseUserRegister struct {
	ID string `json:"_id"`
}

func UserRegister(context echo.Context) error {
	var param paramUserRegister
	if err := context.Bind(&param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}
	if err := context.Validate(param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}

	_, found, err := model.GetUserWithUsername(param.Username)
	if err != nil {
		return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
	}
	if found {
		return util.ErrorResponse(context, http.StatusBadRequest, "username already exists")
	}

	user := model.User{
		Username: param.Username,
		Password: base64.StdEncoding.EncodeToString([]byte(param.Password)),
		Phone:    param.Phone,
		Email:    param.Email,
		IsAdmin:  false,
		Verified: false,
	}
	idHex, err := model.AddUser(user)

	verifyCode := util.RandomString(config.Config.App.VerifyCodeLength)
	err = util.SendEmail(param.Email, "注册邮箱验证码", "您的邮箱验证码为：<code>"+verifyCode+"</code>")
	if err != nil {
		return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
	}
	err = model.AddVerifyCode(verifyCode, idHex)
	if err != nil {
		return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
	}

	return util.SuccessResponse(context, http.StatusCreated, responseUserRegister{
		ID: idHex,
	})
}

type paramUserVerify struct {
	ID   string `json:"_id" validate:"required"`
	Code string `json:"code" validate:"required"`
}

func UserVerify(context echo.Context) error {
	var param paramUserVerify
	if err := context.Bind(&param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}
	if err := context.Validate(param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}

	id, found, err := model.GetVerifyCode(param.Code)
	if !found {
		return util.ErrorResponse(context, http.StatusBadRequest, "verify code not found or expired")
	}
	if err != nil {
		return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
	}
	if id != param.ID {
		return util.ErrorResponse(context, http.StatusBadRequest, "verify code doesn't match")
	}
	err = model.DeleteVerifyCode(param.Code)
	if err != nil {
		return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
	}

	err = model.UpdateUser(param.ID, bson.M{"verified": true})
	if err != nil {
		return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
	}
	return util.SuccessResponse(context, http.StatusOK, nil)
}

type paramUserUpdateInfo struct {
	ID       string `json:"_id" validate:"required"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone" validate:"omitempty,numeric"`
	Email    string `json:"email" validate:"omitempty,email"`
}

func UserUpdateInfo(context echo.Context) error {
	var param paramUserUpdateInfo
	if err := context.Bind(&param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}
	if err := context.Validate(param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}

	id := util.MustGetIDFromContext(context)
	user, _, err := model.GetUserWithID(id)
	if err != nil {
		return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
	}

	info := make(bson.M)
	if param.Username != "" {
		_, found, err := model.GetUserWithUsername(param.Username)
		if err != nil {
			return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
		}
		if param.Username != user.Username && found {
			return util.ErrorResponse(context, http.StatusBadRequest, "username already exists")
		}
		info["username"] = param.Username
	}
	if param.Password != "" {
		info["password"] = base64.StdEncoding.EncodeToString([]byte(param.Password))
	}
	if param.Email != "" {
		verifyCode := util.RandomString(config.Config.App.VerifyCodeLength)
		err := util.SendEmail(param.Email, "注册邮箱验证码", "您的邮箱验证码为：<code>"+verifyCode+"</code>")
		if err != nil {
			return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
		}
		err = model.AddVerifyCode(verifyCode, param.ID)
		if err != nil {
			return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
		}
		info["email"] = param.Email
		info["verified"] = false
	}
	if param.Phone != "" {
		info["phone"] = param.Phone
	}

	err = model.UpdateUser(param.ID, info)
	if err != nil {
		return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
	}
	return util.SuccessResponse(context, http.StatusOK, nil)
}

type paramUserDelete struct {
	ID string `query:"_id" validate:"required"`
}

func UserDelete(context echo.Context) error {
	var param paramUserDelete
	if err := context.Bind(&param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}
	if err := context.Validate(param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}

	id := util.MustGetIDFromContext(context)
	isAdmin, err := model.IsUserAdmin(id)
	if err != nil {
		return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
	}
	if !isAdmin {
		return util.ErrorResponse(context, http.StatusForbidden, "you are not admin")
	}

	err = model.DeleteUser(param.ID)
	if err != nil {
		return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
	}
	return util.SuccessResponse(context, http.StatusOK, nil)
}

type paramUserGetInfo struct {
	ID string `query:"_id"`
}

func UserGetInfo(context echo.Context) error {
	var param paramUserGetInfo
	if err := context.Bind(&param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}
	if err := context.Validate(param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}

	var result []model.User
	id := util.MustGetIDFromContext(context)
	isAdmin, err := model.IsUserAdmin(id)
	if err != nil {
		return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
	}

	if param.ID != "" {
		if param.ID != id && !isAdmin {
			return util.ErrorResponse(context, http.StatusForbidden, "you are not admin")
		}

		user, found, err := model.GetUserWithID(param.ID)
		if !found {
			return util.ErrorResponse(context, http.StatusBadRequest, "user not found")
		}
		if err != nil {
			return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
		}
		result = []model.User{user}
	} else { // id not set, query for all users
		if !isAdmin {
			return util.ErrorResponse(context, http.StatusForbidden, "you are not admin")
		}

		users, err := model.GetAllUsers()
		if err != nil {
			return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
		}
		result = users
	}

	response := make([]echo.Map, 0)
	for _, user := range result {
		response = append(response, echo.Map{
			"_id":      user.ID.Hex(),
			"username": user.Username,
			"phone":    user.Phone,
			"email":    user.Email,
		})
	}
	return util.SuccessResponse(context, http.StatusOK, echo.Map{"result": response})
}
