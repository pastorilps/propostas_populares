package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"

	_ "github.com/pastorilps/propostas_populares/app/docs"
	middleware_cors "github.com/pastorilps/propostas_populares/middlewares"
	_migrations "github.com/pastorilps/propostas_populares/migrations"
	_userHttpDelivery "github.com/pastorilps/propostas_populares/users/delivery/http"
	_userRepo "github.com/pastorilps/propostas_populares/users/repository"
	_userUseCase "github.com/pastorilps/propostas_populares/users/usecase"

	_authHttpDelivery "github.com/pastorilps/propostas_populares/authenticate/delivery"
	_authRepo "github.com/pastorilps/propostas_populares/authenticate/repository"
	_authUseCase "github.com/pastorilps/propostas_populares/authenticate/usecase"

	_proposalHttpDelivery "github.com/pastorilps/propostas_populares/propostas/delivery/http"
	_proposalRepo "github.com/pastorilps/propostas_populares/propostas/repository"
	_proposalUseCase "github.com/pastorilps/propostas_populares/propostas/usecase"
)

type Env struct {
	db *sql.DB
}

func init() {
	gotenv.Load()
}

// @title Echo Propostas Populares API
// @version 1.0
// @description This is a sample crud api.
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

	// Middleware
	mid := middleware_cors.InitMiddleware()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(mid.CORS)

	userRepo := _userRepo.NewUserRepo(dbConn)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	_userHttpDelivery.NewUserHandler(e, userUseCase)

	authRepo := _authRepo.NewAuthenticateRepo(dbConn)
	authUseCase := _authUseCase.NewAuthenticateUsecase(authRepo)
	_authHttpDelivery.NewAuthenticateHandler(e, authUseCase)

	proposalRepo := _proposalRepo.NewProposalRepo(dbConn)
	proposalUseCase := _proposalUseCase.NewProposalUseCase(proposalRepo)
	_proposalHttpDelivery.NewProposalHandler(e, proposalUseCase)

	_migrations.Exec(dbConn)

	log.Fatal(e.Start(":7500"))
}
