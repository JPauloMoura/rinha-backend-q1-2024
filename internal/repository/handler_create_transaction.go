package repository

import (
	"context"
	"log/slog"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/entities"
	"github.com/jackc/pgx/v5"
)

func CreateTransaction(client *entities.Client, transaction entities.Transaction) error {
	tx, err := DB.Begin(context.TODO())
	if err != nil {
		slog.Error("Error starting transaction: ", err)
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback(context.TODO())
			slog.Error("Transaction rolled back: ", err)
		}
	}()

	err = saveTransaction(tx, transaction)
	if err != nil {
		return err
	}

	err = updateClientSaldo(tx, client)
	if err != nil {
		return err
	}

	err = tx.Commit(context.TODO())
	if err != nil {
		slog.Error("Error committing transaction: ", err)
		return err
	}

	return nil
}

func saveTransaction(tx pgx.Tx, transaction entities.Transaction) error {
	_, err := tx.Exec(context.TODO(), "INSERT INTO transactions (clientId, valor, tipo, descricao) VALUES ($1, $2, $3, $4)", transaction.ClientId, transaction.Value, transaction.Type, transaction.Description)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}

func updateClientSaldo(tx pgx.Tx, client *entities.Client) error {
	_, err := tx.Exec(context.TODO(), "UPDATE clientes SET saldo = $1 WHERE id = $2", client.Saldo, client.Id)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}
