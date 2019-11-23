package main

import (
	user "github.com/Ledka17/Back_chat_lemmas/delivery/ws"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"log"
	"os"
)

func main() {
	e = echo.New()

	init
	http.Handle("/message", user.GetMessagesHandler)

	log.Println("http server started on :8000")
	port := getPort()
	log.Fatal(http.ListenAndServe(port, nil))
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
