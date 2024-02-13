package repository

import (
	"database/sql"
	"log"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connectionString := os.Getenv("DB_CONNECTION_STRING")

	connect, err := sql.Open("postgres", connectionString)
	if err != nil || connect == nil {
		slog.Debug("failed to open conection. check if the database is running or if the connection string is correct",
			slog.Any("error", err),
			slog.String("connectionString", connectionString),
		)
		log.Fatal("failed to open conection", err)
	}

	if err := connect.Ping(); err != nil {
		slog.Debug("failed to ping on database. check if the database is running or if the connection string is correct",
			slog.Any("error", err),
			slog.String("connectionString", connectionString),
		)
		log.Fatal("failed to ping on database", err)
	}

	DB = connect

	if resetDatabase() != nil {
		log.Panic("failed to reset db")
	}
}

func resetDatabase() error {
	tx, err := DB.Begin()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			slog.Error("Transaction rolled back: ", err)
		}
	}()

	// drop data
	_, err = tx.Exec("DROP TABLE IF EXISTS clientes;")
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	//  create table clients
	_, err = tx.Exec(`
	CREATE TABLE IF NOT EXISTS clientes (
		id serial PRIMARY KEY,
		nome VARCHAR(50),
		limite INT,
		saldo INT
	);
`)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	// inserts data
	_, err = tx.Exec(`
		INSERT INTO clientes (nome, limite, saldo)
		VALUES
		('o barato sai caro', 1000 * 100, 0),
		('zan corp ltda', 800 * 100, 0),
		('les cruders', 10000 * 100, 0),
		('padaria joia de cocaia', 100000 * 100, 0),
		('kid mais', 5000 * 100, 0);
	`)

	if err != nil {
		slog.Error(err.Error())
		return err
	}

	// drop data
	_, err = tx.Exec("DROP TABLE IF EXISTS transactions;")
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	//  create table transactions
	_, err = tx.Exec(`
		CREATE TABLE IF NOT EXISTS transactions (
			id serial PRIMARY KEY,
			clientId INT,
			valor INT,
			tipo varchar(1),
			descricao varchar(10),
			realizada_em TIMESTAMP WITH TIME ZONE DEFAULT timezone('America/Sao_Paulo'::text, now())
		);
	`)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}
