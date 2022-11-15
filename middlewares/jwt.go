package middlewares

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Claims struct {
	jwt.StandardClaims
	UserID int16
}

var jwtKey = []byte("unicorns")

func GenerateJwt(mp map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["expires"] = time.Now().Add(time.Minute * 30).Unix()
	for key, val := range mp {
		claims[key] = val
	}

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		logrus.Error("Something Went Wrong: %s", err.Error())
		return " ", err
	}

	return tokenString, nil
}

func GetUserIDJWT(tokenstr string) int16 {
	const secret_key = ""
	Claims := &Claims{}
	jwt.ParseWithClaims(tokenstr, Claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret_key), nil
	})

	return Claims.UserID
}

func GetToken(c echo.Context) int16 {
	header := c.Request().Header.Get("Authorization")
	token := header[len("Bearer "):]
	UserID := GetUserIDJWT(token)
	return UserID
}
