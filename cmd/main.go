package main

import (
	"log/slog"
	"net/http"
	"os"

	// _ "net/http/pprof"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/handler"
	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/repository"
	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	pool := repository.ConnectDB()

	h := handler.Handler{
		Svc: service.Service{
			Repo: repository.Repo{
				DB: repository.Database{
					Connection: pool,
				},
			},
		},
	}

	router := chi.NewMux()
	router.Post("/clientes/{id}/transacoes", h.CreateTransaction)
	router.Get("/clientes/{id}/extrato", h.GenerateExtract)

	// if os.Getenv("PPROF") != "" {
	// 	println("======== SEM PPROF =======")
	// 	go func() {
	// 		http.ListenAndServe(":3333", nil)
	// 	}()
	// }

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
