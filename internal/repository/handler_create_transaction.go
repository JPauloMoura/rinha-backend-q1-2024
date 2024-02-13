package repository

import (
	"database/sql"
	"log/slog"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/entities"
)

func CreateTransaction(client *entities.Client, transaction entities.Transaction) error {
	tx, err := DB.Begin()
	if err != nil {
		slog.Error("Error starting transaction: ", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			slog.Error("Transaction rolled back: ", err)
		}
	}()

	err = saveTransaction(tx, transaction)
	if err != nil {
		slog.Error("Error saving transaction: ", err)
		return err
	}

	err = updateClientSaldo(tx, client)
	if err != nil {
		slog.Error("Error updating client saldo: ", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		slog.Error("Error committing transaction: ", err)
		return err
	}

	slog.Info("Transaction completed successfully!")
	return nil
}

func saveTransaction(tx *sql.Tx, transaction entities.Transaction) error {
	_, err := tx.Exec("INSERT INTO transactions (clientId, valor, tipo, descricao) VALUES ($1, $2, $3, $4)",
		transaction.ClientId, transaction.Value, transaction.Type, transaction.Description)

	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}

func updateClientSaldo(tx *sql.Tx, client *entities.Client) error {
	_, err := tx.Exec("UPDATE clientes SET saldo = $1 WHERE id = $2", client.Saldo, client.Id)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	return nil
}
