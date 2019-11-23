package main

import (
	"github.com/Ledka17/Back_chat_lemmas/delivery/http"
	user "github.com/Ledka17/Back_chat_lemmas/user/delivery/http"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"log"
	"os"
)

func main() {
	e := echo.New()

	initHandlers(e)

	log.Println("http server started on :8000")
	port := getPort()
	log.Fatal(e.Start(":" + port))
}

func initHandlers(e *echo.Echo) {
	e.GET(http.ApiV1UserGetMessages, user.GetUserMessages)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func NewDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", os.Getenv("POSTGRES_DSN"))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
