package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/pastorilps/propostas_populares/app/docs"
	middlewares "github.com/pastorilps/propostas_populares/middlewares"
	_migrations "github.com/pastorilps/propostas_populares/propostas/repository/migrations"
	"github.com/pastorilps/propostas_populares/users/model"
)

type Env struct {
	db *sql.DB
}

func init() {
	gotenv.Load()
}

// @title Echo Propostas Populares API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:7500
// @BasePath /
// @schemes http
func main() {

	var (
		dbHost = os.Getenv("HOST_DATABASE")
		dbPort = os.Getenv("PORT_DATABASE")
		dbUser = os.Getenv("USER_DATABASE")
		dbPass = os.Getenv("PASS_DATABASE")
		dbName = os.Getenv("DB_DATABASE")
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable  ", dbHost, dbPort, dbUser, dbPass, dbName)
	dbConn, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalln(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatalln(err)

	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	defer dbConn.Close()
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/user/create", func(c echo.Context) error {
		u := new(model.UserData)
		if err := c.Bind(u); err != nil {
			return err
		}

		u.Password = middlewares.SHA256Encoder(u.Password)

		sqlStatement := "INSERT INTO public.user (name, email, password, picture, newsletter)VALUES ($1, $2, $3, $4, $5)"
		res, err := dbConn.Query(sqlStatement, u.Name, u.Email, u.Password, u.Picture, u.Newsletter)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
			return c.JSON(http.StatusCreated, u)
		}
		return c.String(http.StatusOK, "User created")
	})
	_migrations.Exec(dbConn)

	log.Fatal(e.Start(":7500"))
}

type UserData struct {
	Name       string `json:"name"  validate:"required" example:"Name"`
	Email      string `json:"email"  validate:"required" example:"test@test.com"`
	Password   string `json:"password"  validate:"required" example:"aB@1245"`
	Picture    int    `json:"picture"  validate:"required" example:"1"`
	Newsletter bool   `json:"newsletter"  validate:"required" example:"true"`
}

// Create User godoc
// @Summary Creates a user.
// @Description creates users.
// @Tags user
// @Param Body body UserData true "The body to create a thing"
// @Accept json
// @Success 200
// @Router /user/create [post]
func createUser(dbConn *sql.DB, c echo.Context) error {
	u := new(model.UserData)
	if err := c.Bind(u); err != nil {
		return err
	}

	u.Password = middlewares.SHA256Encoder(u.Password)

	sqlStatement := "INSERT INTO public.user (name, email, password, picture, newsletter)VALUES ($1, $2, $3, $4, $5)"
	res, err := dbConn.Query(sqlStatement, u.Name, u.Email, u.Password, u.Picture, u.Newsletter)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
		return c.JSON(http.StatusCreated, u)
	}
	return c.String(http.StatusOK, "User created")
}
