package controller

import (
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"qq-bot-backend/model"
	"qq-bot-backend/util"
	"time"
)

type responseEventAdd struct {
	ID string `json:"_id"`
}

func EventAdd(context echo.Context) error {
	var param model.Event
	if err := context.Bind(&param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}
	if err := context.Validate(param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}

	idHex, err := model.AddEvent(param)
	if err != nil {
		return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
	}
	return util.SuccessResponse(context, http.StatusOK, responseEventAdd{ID: idHex})
}

type paramEventGet struct {
	ID     string `query:"_id"`
	User   string `query:"user"`
	Remind int    `query:"remind"`
}

type responseEventGet struct {
	Result []model.Event `json:"result"`
}

func EventGet(context echo.Context) error {
	var param paramEventGet
	if err := context.Bind(&param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}
	if err := context.Validate(param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}

	condition := make(bson.M)
	if param.ID != "" {
		id, err := primitive.ObjectIDFromHex(param.ID)
		if err != nil {
			return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
		}
		condition["_id"] = id
	}
	if param.User != "" {
		condition["user"] = param.User
	}
	if param.Remind != 0 {
		condition["remind"] = true
		condition["remind_time"] = bson.M{"$lte": time.Now().Unix()}
	}

	events, err := model.GetEventsByCondition(condition)
	if err != nil {
		return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
	}
	return util.SuccessResponse(context, http.StatusOK, responseEventGet{Result: events})
}

type paramEventDelete struct {
	ID string `query:"_id" validate:"required"`
}

func EventDelete(context echo.Context) error {
	var param paramEventDelete
	if err := context.Bind(&param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}
	if err := context.Validate(param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}

	err := model.DeleteEvent(param.ID)
	if err != nil {
		return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
	}
	return util.SuccessResponse(context, http.StatusOK, nil)
}

type paramEventUpdate struct {
	ID             string `json:"_id" validate:"required"`
	User           string `json:"user"`
	Desc           string `json:"desc"`
	Time           int64  `json:"time"`
	Remind         bool   `json:"remind"`
	RemindTime     int64  `json:"remind_time"`
	RemindInterval int    `json:"remind_interval"`
}

func EventUpdate(context echo.Context) error {
	var param paramEventUpdate
	if err := context.Bind(&param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}
	if err := context.Validate(param); err != nil {
		return util.ErrorResponse(context, http.StatusBadRequest, err.Error())
	}

	event, found, err := model.GetEventByID(param.ID)
	if !found {
		return util.ErrorResponse(context, http.StatusBadRequest, "event with _id "+param.ID+" not found")
	}
	if err != nil {
		return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
	}

	update := make(bson.M)
	if param.User != event.User {
		update["user"] = param.User
	}
	if param.Desc != event.Desc {
		update["desc"] = param.Desc
	}
	if param.Time != event.Time {
		update["time"] = param.Time
	}
	if param.Remind != event.Remind {
		update["remind"] = param.Remind
		if param.Remind && (param.RemindTime == 0 || param.RemindInterval == 0) {
			return util.ErrorResponse(context, http.StatusBadRequest, "remind_time or remind_interval must be provided")
		}
	}
	if param.RemindTime != event.RemindTime {
		update["remind_time"] = param.RemindTime
	}
	if param.RemindInterval != event.RemindInterval {
		update["remind_interval"] = param.RemindInterval
	}

	err = model.UpdateEvent(param.ID, update)
	if err != nil {
		return util.ErrorResponse(context, http.StatusInternalServerError, err.Error())
	}
	return util.SuccessResponse(context, http.StatusOK, nil)
}
