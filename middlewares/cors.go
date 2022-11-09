package middlewares

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type GoMiddleware struct {
}

func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Accept", "application/json")
		c.Response().Header().Set("Content-Type", "application/json")
		c.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		return next(c)
	}
}

func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}

var (
	// ErrInternalServerError will thow if the internal server error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound wil thow if the requested item is not exists
	ErrorNotFound = errors.New("Your requested Item is not found")
	// ErrConflict willthow if the current action already exists
	ErrConflict = errors.New("Your Item already exist")
	// ErrBadParamInput will throw if the given request-body or param is not valid
	ErrBadParamInput = errors.New("Given Param is not valid")
)
