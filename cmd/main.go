package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/handler"
	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	defineLogger()
	repository.ConnectDB()

	router := chi.NewMux()
	router.Post("/clientes/{id}/transacoes", handler.CreateTransaction)
	router.Get("/clientes/{id}/extrato", handler.GenerateExtract)

	slog.Info("server is running...")
	http.ListenAndServe(":3001", router)
}

func defineLogger() {
	l := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	})

	slog.SetDefault(slog.New(l))
}
