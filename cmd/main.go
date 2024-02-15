package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/handler"
	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	defineLogger()
	repository.ConnectDB()

	router := chi.NewMux()
	router.Use(middleware.Recoverer)
	router.Post("/clientes/{id}/transacoes", handler.CreateTransaction)
	router.Get("/clientes/{id}/extrato", handler.GenerateExtract)

	slog.Info("server is running...")
	http.ListenAndServe(":3001", router)
}

func defineLogger() {
	l := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelWarn,
	})

	slog.SetDefault(slog.New(l))
}
