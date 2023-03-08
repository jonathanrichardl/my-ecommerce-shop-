package main

import (
	"net/http"
	"os"
)

func main() {
	server := createServer()
	err := http.ListenAndServe(os.Getenv("APP_HOST"), server)

	if err != nil {
		panic(err)
	}
}
