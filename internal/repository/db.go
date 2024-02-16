package repository

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	ctx := context.TODO()
	connectionString := os.Getenv("DB_CONNECTION_STRING")

	cfg, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		slog.Debug("failed parse config. check if the database is running or if the connection string is correct",
			slog.Any("error", err),
			slog.String("connectionString", connectionString),
		)
		log.Fatal("failed parse config", err)
	}

	cfg.MaxConns = 20
	cfg.MinConns = 10

	conn, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		slog.Debug("failed to create connection. check if the database is running or if the connection string is correct",
			slog.Any("error", err),
			slog.String("connectionString", connectionString),
		)
		log.Fatal("failed to create connection", err)
	}

	if err != nil || conn == nil {
		slog.Debug("failed to open conection. check if the database is running or if the connection string is correct",
			slog.Any("error", err),
			slog.String("connectionString", connectionString),
		)
		log.Fatal("failed to open conection", err)
	}

	if err := conn.Ping(ctx); err != nil {
		slog.Debug("failed to ping on database. check if the database is running or if the connection string is correct",
			slog.Any("error", err),
			slog.String("connectionString", connectionString),
		)
		log.Fatal("failed to ping on database", err)
	}

	DB = conn
}
