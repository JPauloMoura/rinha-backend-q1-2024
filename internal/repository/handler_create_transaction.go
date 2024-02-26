package repository

import (
	"context"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/entities"
)

type CreateTransactionResponse struct {
	Balance int `json:"saldo"`
	Limit   int `json:"limite"`
}

func (h Repo) CreateTransaction(ctx context.Context, accId int, transaction entities.Transaction) (*CreateTransactionResponse, error) {
	var (
		balance       int
		limit         int
		clientAccount entities.ClientAccount
	)

	tx, err := h.DB.Connection.Begin(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.QueryRow(ctx, `
			SELECT balance, acc_limit FROM client_account
			WHERE id = $1 FOR UPDATE`, accId,
	).Scan(&balance, &limit)

	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	clientAccount = entities.ClientAccount{
		Balance: balance,
		Limit:   limit,
	}

	err = clientAccount.SetNewBalanceByTransaction(transaction)
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	_, err = tx.Exec(ctx, `
			UPDATE client_account SET balance = $1
			WHERE id = $2`,
		clientAccount.Balance, accId,
	)

	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO transaction (account_id, value, type, description)
		VALUES ($1, $2, $3, $4)`,
		accId, transaction.Value, transaction.Type, transaction.Description,
	)

	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	return &CreateTransactionResponse{
		Balance: clientAccount.Balance,
		Limit:   clientAccount.Limit,
	}, nil
}
