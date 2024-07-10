package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"twib/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Info("Env file not found, using enviornment values")
	}
	port := os.Getenv("PORT")

	router := chi.NewMux()
	router.Handle("/*", public())
	router.Get("/", handlers.HandleHome)

	// Server -------------------------
	server := http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           router,
		ReadHeaderTimeout: time.Second * 10,
		WriteTimeout:      time.Second * 20,
		IdleTimeout:       time.Minute * 2,
	}

	// Start Server
	slog.Info("Server listing on", "port", port)
	slog.Error("Server terminated", "Err", server.ListenAndServe())
}
