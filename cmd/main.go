package main

import (
	"log"
	"os"

	"github.com/s444v/go-weather-api/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Llongfile)
	server := server.NewServer(logger)
	err := server.HTTP.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
