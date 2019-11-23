package main

import (
	userDelivery "github.com/Ledka17/Back_chat_lemmas/user/delivery/http"
	"github.com/Ledka17/Back_chat_lemmas/user/repository"
	"github.com/Ledka17/Back_chat_lemmas/user/usecase"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"log"
	"os"
)

func main() {
	e := echo.New()

	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewDatabaseRepository(db)
	userDelivery.NewUserHandler(usecase.NewUserUsecase(repo), e)

	log.Println("http server started on :8000")
	port := getPort()
	log.Fatal(e.Start(":" + port))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
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
