package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
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