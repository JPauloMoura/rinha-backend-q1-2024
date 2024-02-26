package repository

import (
	"context"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/entities"
)

func (h Repo) ListTransaction(ctx context.Context, id int) (*TransactionsWithAccount, error) {
	items, err := h.DB.Connection.Query(ctx, `
		SELECT value, type, description, created_at 
		FROM transaction
		WHERE account_id = $1
		ORDER BY created_at DESC
		LIMIT 10
	`, id)

	if err != nil {
		return nil, err
	}

	defer items.Close()

	var extract TransactionsWithAccount

	for items.Next() {
		var t entities.Transaction

		if err := items.Scan(
			&t.Value,
			&t.Type,
			&t.Description,
			&t.CreatedAt,
		); err != nil {
			return nil, err
		}

		extract.Transactions = append(extract.Transactions, t)
	}

	err = h.DB.Connection.QueryRow(ctx, `
		SELECT balance, acc_limit
		FROM client_account
		WHERE id = $1`, id,
	).Scan(&extract.Account.Balance, &extract.Account.Limit)

	if err != nil {
		return nil, err
	}

	return &extract, nil
}

type TransactionsWithAccount struct {
	Account      entities.ClientAccount
	Transactions []entities.Transaction
}
