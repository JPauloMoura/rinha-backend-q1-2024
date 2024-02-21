package repository

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB() *pgxpool.Pool {

	mx, err := strconv.Atoi(os.Getenv("MAX_CONN"))
	if err != nil {
		log.Fatal(err)
	}
	min, err := strconv.Atoi(os.Getenv("MIN_CONN"))
	if err != nil {
		log.Fatal(err)
	}

	var (
		defaultMaxConns          = int32(mx)
		defaultMinConns          = int32(min)
		defaultMaxConnLifetime   = time.Hour
		defaultMaxConnIdleTime   = time.Minute * 30
		defaultHealthCheckPeriod = time.Minute
		defaultConnectTimeout    = time.Second * 30
	)

	connectionString := os.Getenv("DB_CONNECTION_STRING")

	cfg, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		// slog.Debug("failed parse config. check if the database is running or if the connection string is correct",
		// slog.Any("error", err),
		// slog.String("connectionString", connectionString),
		// )
		log.Fatal("failed parse config", err)
	}

	cfg.MaxConns = defaultMaxConns
	cfg.MinConns = defaultMinConns
	cfg.MaxConnLifetime = defaultMaxConnLifetime
	cfg.MaxConnIdleTime = defaultMaxConnIdleTime
	cfg.HealthCheckPeriod = defaultHealthCheckPeriod
	cfg.ConnConfig.ConnectTimeout = defaultConnectTimeout

	ctx := context.TODO()
	conn, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		// slog.Debug("failed to create connection. check if the database is running or if the connection string is correct",
		// slog.Any("error", err),
		// slog.String("connectionString", connectionString),
		// )
		log.Fatal("failed to create connection", err)
	}

	if err != nil || conn == nil {
		// slog.Debug("failed to open conection. check if the database is running or if the connection string is correct",
		// slog.Any("error", err),
		// slog.String("connectionString", connectionString),
		// )
		log.Fatal("failed to open conection", err)
	}

	if err := conn.Ping(ctx); err != nil {
		// slog.Debug("failed to ping on database. check if the database is running or if the connection string is correct",
		// slog.Any("error", err),
		// slog.String("connectionString", connectionString),
		// )
		log.Fatal("failed to ping on database", err)
	}

	return conn
}

type Database struct {
	Connection *pgxpool.Pool
}

type Repo struct {
	DB Database
}
