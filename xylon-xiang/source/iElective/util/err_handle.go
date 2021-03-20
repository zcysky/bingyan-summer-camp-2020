package util

import (
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func HttpErrorHandle(ctx echo.Context, err error) error {
	switch err {
	case mongo.ErrNoDocuments:
		return ctx.String(http.StatusNotFound, "No such a document in the database")

	case mongo.ErrNilDocument:
		return ctx.String(http.StatusBadRequest, "The document is nil")

	case mongo.ErrNilCursor:
		return ctx.String(http.StatusBadRequest, "The cursor in the select function is nil")

	default:
		return ctx.String(http.StatusInternalServerError, "Unknown error")
	}
}
