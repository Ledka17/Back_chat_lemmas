package main

import (
	chatDelivery "github.com/Ledka17/Back_chat_lemmas/chat/delivery/ws"
	chatRepo "github.com/Ledka17/Back_chat_lemmas/chat/repository"
	chatUsecase "github.com/Ledka17/Back_chat_lemmas/chat/usecase"
	supportDelivery "github.com/Ledka17/Back_chat_lemmas/support/delivery/http"
	supportRepo "github.com/Ledka17/Back_chat_lemmas/support/repository"
	supportUsecase "github.com/Ledka17/Back_chat_lemmas/support/usecase"
	userDelivery "github.com/Ledka17/Back_chat_lemmas/user/delivery/http"
	userRepo "github.com/Ledka17/Back_chat_lemmas/user/repository"
	userUsecase "github.com/Ledka17/Back_chat_lemmas/user/usecase"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	e := echo.New()
	e.Use(corsMiddleware)

	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
	}

	chatDelivery.NewChatHandler(e, chatUsecase.NewChatUsecase(chatRepo.NewDatabaseRepository(db)))
	userDelivery.NewUserHandler(e, userUsecase.NewUserUsecase(userRepo.NewDatabaseRepository(db)))
	supportDelivery.NewSupportHandler(e, supportUsecase.NewSupportUsecase(supportRepo.NewDatabaseRepository(db)))

	log.Println("http server started on :8000")
	port := getPort()
	log.Fatal(e.Start(":" + port))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
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

func corsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		origin := c.Request().Header.Get(echo.HeaderOrigin)
		allowOrigin := origin
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, allowOrigin)

		allowedMethods := []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodOptions, http.MethodDelete}
		c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, strings.Join(allowedMethods, ","))
		c.Response().Header().Set(echo.HeaderAccessControlAllowCredentials, "true")

		allowedHeaders := []string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}
		c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, strings.Join(allowedHeaders, ","))

		if c.Request().Method == http.MethodOptions {
			return c.NoContent(http.StatusNoContent)
		}
		return next(c)
	}
}
