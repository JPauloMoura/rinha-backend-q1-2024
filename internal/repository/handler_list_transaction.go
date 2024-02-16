package repository

import (
	"context"
	"errors"
	"log/slog"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/entities"
)

func ListTransaction(id int) ([]entities.Transaction, error) {
	items, err := DB.Query(context.TODO(), "SELECT * FROM transactions WHERE clientId=$1", id)
	if err != nil {
		slog.Error(err.Error())
		return nil, errors.New("internal error")
	}

	var transactions []entities.Transaction

	for items.Next() {
		var t entities.Transaction

		if err := items.Scan(&t.ClientId, &t.Value, &t.Type, &t.Description, &t.CreatedAt); err != nil {
			slog.Error(err.Error())
			return nil, errors.New("internal error")
		}

		transactions = append(transactions, t)
	}

	return transactions, nil
}
