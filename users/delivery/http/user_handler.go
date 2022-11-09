package http

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	_ "github.com/pastorilps/propostas_populares/app/docs"
	middleware "github.com/pastorilps/propostas_populares/middlewares"
	"github.com/pastorilps/propostas_populares/users/domain"
	"github.com/pastorilps/propostas_populares/users/entity"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Response struct {
	Message string `json:"message"`
}

type UserHandler struct {
	AUsecase domain.UserUseCase
}

func NewUserHandler(e *echo.Echo, uc domain.UserUseCase) {
	handler := &UserHandler{
		AUsecase: uc,
	}

	// Routes
	e.GET("/", HealthCheck)
	e.GET("/v1/users", handler.GetAllUsers)
	e.GET("/v1/users/:id", handler.GetUserById)
	e.POST("/v1/users/create", handler.CreateUser)
	e.PUT("/v1/users/update/:id", handler.UpdateUser)
	e.DELETE("/v1/users/delete/:id", handler.DeleteUser)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

}

// DeleteUser godoc
// @Summary Delete User.
// @Description Delete User Data.
// @Tags Users
// @Accept json
// @Param id path integer true "User ID"
// @Produce json
// @Success 200 "Usuário deletado com sucesso!"
// @Router /v1/users/delete/{id} [delete]
func (u *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	err = u.AUsecase.DeleteUser(int16(id))
	if err != nil {
		return c.JSON(getStatusCode(err), Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, "Usuário deletado com sucesso!")
}

// UpdateUser godoc
// @Summary Update User.
// @Description Update User Data.
// @Tags Users
// @Accept json
// @Param id path integer true "User ID"
// @Param Body body entity.Send_User true "The body to update a user"
// @Produce json
// @Success 200 {object} entity.Send_User
// @Router /v1/users/update/{id} [put]
func (u *UserHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, Response{Message: err.Error()})
	}

	var receive entity.Receive_User
	err = c.Bind(&receive)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	receive.ID = int16(id)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	list, err := u.AUsecase.UpdateUser(ctx, &receive)
	if err != nil {
		return c.JSON(getStatusCode(err), Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, list)
}

// CreateUser godoc
// @Summary Create User.
// @Description Create User.
// @Tags Users
// @Accept json
// @Param Body body entity.Send_User true "The body to create a user"
// @Produce json
// @Success 200 {object} entity.Users
// @Router /v1/users/create [post]
func (u *UserHandler) CreateUser(c echo.Context) error {
	l := new(entity.Users)
	if err := c.Bind(l); err != nil {
		return err
	}

	_, err := u.AUsecase.CreateUsers(l)
	if err != nil {
		return c.JSON(getStatusCode(err), Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, l)
}

// GetUser godoc
// @Summary Show an user.
// @Description Get user.
// @Tags Users
// @Param id path integer true "User ID"
// @Accept */*
// @Produce json
// @Success 200 {object} entity.Users
// @Router /v1/users/{id} [get]
func (u *UserHandler) GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	uId, err := u.AUsecase.GetUserById(int16(id))
	if err != nil {
		return c.JSON(getStatusCode(err), Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, uId)
}

// GetAllUsers godoc
// @Summary Show all users.
// @Description Get all users list.
// @Tags Users
// @Accept */*
// @Produce json
// @Success 200 {object} entity.Users
// @Router /v1/users [get]
func (u *UserHandler) GetAllUsers(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	data, err := u.AUsecase.GetAllUsers()
	if err != nil {
		return c.JSON(getStatusCode(err), Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Fatalln(err)

	switch err {
	case middleware.ErrInternalServerError:
		return http.StatusInternalServerError
	case middleware.ErrorNotFound:
		return http.StatusNotFound
	case middleware.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
