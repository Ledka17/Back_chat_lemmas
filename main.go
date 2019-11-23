package main

import (
	"fmt"
	userRepo "github.com/Ledka17/Back_chat_lemmas/user/repository"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func main() {
	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
	}
	repo := userRepo.NewDatabaseRepository(db)
	users, err := repo.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
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
