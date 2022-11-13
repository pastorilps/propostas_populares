package delivery

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	_ "github.com/pastorilps/propostas_populares/app/docs"
	domain "github.com/pastorilps/propostas_populares/authenticate/domain"
	"github.com/pastorilps/propostas_populares/authenticate/entity"
	"github.com/pastorilps/propostas_populares/middlewares"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type ResponseError struct {
	Message string `json:"message"`
}

type AuthenticateHandler struct {
	AUsecase domain.AuthenticateUsecase
}

func NewAuthenticateHandler(e *echo.Echo, du domain.AuthenticateUsecase) {
	handler := &AuthenticateHandler{
		AUsecase: du,
	}
	e.POST("/v1/auth/users/signing", handler.GetByUserLogin)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

// Login godoc
// @Summary User Login.
// @Description User Login.
// @Tags Login
// @Accept json
// @Param Body body entity.Receive_Login_Data true "The body to login"
// @Produce json
// @Success 200 {object} entity.Auth_Token
// @Router /v1/auth/users/signing [post]
func (uh *AuthenticateHandler) GetByUserLogin(c echo.Context) error {
	var receiveLoginData entity.Receive_Login_Data
	err := c.Bind(&receiveLoginData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if ok, err := isRequestValid(&receiveLoginData); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	UserData, err := uh.AUsecase.GetByUserLogin(ctx, &receiveLoginData)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, UserData)
}

func isRequestValid(er *entity.Receive_Login_Data) (bool, error) {
	validate := validator.New()
	err := validate.Struct(er)
	if err != nil {
		return false, err
	}
	return true, nil
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)

	switch err {
	case middlewares.ErrInternalServerError:
		return http.StatusInternalServerError
	case middlewares.ErrorNotFound:
		return http.StatusNotFound
	case middlewares.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
