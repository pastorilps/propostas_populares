package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"

	_migrations "github.com/pastorilps/propostas_populares/propostas/repository/migrations"
)

type Env struct {
	db *sql.DB
}

func init() {
	gotenv.Load()
}

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

	_migrations.Exec(dbConn)

	log.Fatal(e.Start(":8087"))
}
