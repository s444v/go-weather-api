package server

import (
	"log"
	"net/http"
	"time"

	"github.com/s444v/go-weather-api/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	HTTP   *http.Server
}

func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.MainHandler)
	mux.HandleFunc("/weather", handlers.UploadHandler)

	httpServer := &http.Server{
		Addr:         "127.0.0.1:5500",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		HTTP:   httpServer,
	}
}
