package main

import (
	"log"
	"net/http"
	"os"
	user "github.com/Ledka17/Back_chat_lemmas/delivery/ws"
	"github.com/labstack/echo"
)

func main() {
	e = echo.New()

	init
	http.Handle("/message", user.GetMessagesHandler)

	log.Println("http server started on :8000")
	port := getPort()
	log.Fatal(http.ListenAndServe(port, nil))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	return port
}