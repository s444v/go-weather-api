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
	mux.HandleFunc("/weather", handlers.WeatherHandler)
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	httpServer := &http.Server{
		Addr:         ":8080",
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
