package main

import (
	"log/slog"
	"net/http"
	"os"

	// _ "net/http/pprof"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/handler"
	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/repository"
	"github.com/go-chi/chi/v5"
)

func main() {
	// defineLogger()
	repository.ConnectDB()

	router := chi.NewMux()
	router.Post("/clientes/{id}/transacoes", handler.CreateTransaction)
	router.Get("/clientes/{id}/extrato", handler.GenerateExtract)

	if os.Getenv("PPROF") != "" {
		go func() {
			http.ListenAndServe(":3333", nil)
		}()
	}

	slog.Info("server is running...")
	http.ListenAndServe(":3001", router)
}

func defineLogger() {
	l := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})

	slog.SetDefault(slog.New(l))
}
