package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/entities"
	"github.com/JPauloMoura/rinha-backend-q1-2024/pkg/errors"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Println("==> ", name, elapsed.Milliseconds())
}

func CreateTransaction(ctx context.Context, clientId int, transaction entities.Transaction) (*CreateTransactionResponse, error) {
	tx, err := DB.Begin(ctx)
	if err != nil {
		// slog.Error("Error starting transaction: ", err)
		return nil, err
	}

	var (
		clientSaldo  int
		clientLimite int
	)

	err = tx.QueryRow(ctx, "SELECT saldo, limite FROM clientes WHERE id = $1 FOR UPDATE", clientId).
		Scan(&clientSaldo, &clientLimite)

	if err != nil {
		// slog.Error("client not found", err)
		return nil, err
	}

	switch transaction.Type {
	case "c":
		clientSaldo += transaction.Value
	case "d":
		clientSaldo -= transaction.Value
	}

	if clientSaldo < (clientLimite * -1) {
		return nil, errors.ErrInvalidTransaction
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
			// slog.Error("Transaction rolled back: ", err)
		}
	}()

	_, err = tx.Exec(ctx, "INSERT INTO transactions (clientId, valor, tipo, descricao) VALUES ($1, $2, $3, $4)", clientId, transaction.Value, transaction.Type, transaction.Description)
	if err != nil {
		// slog.Error(err.Error())
		return nil, err
	}

	_, err = tx.Exec(ctx, "UPDATE clientes SET saldo = $1 WHERE id = $2", clientSaldo, clientId)
	if err != nil {
		// slog.Error(err.Error())
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		// slog.Error("Error committing transaction: ", err)
		return nil, err
	}

	return &CreateTransactionResponse{
		Saldo: clientSaldo,
		Limit: clientLimite,
	}, nil
}

type CreateTransactionResponse struct {
	Saldo int `json:"saldo"`
	Limit int `json:"limite"`
}
