package repository

import (
	"context"
	"time"
)

func (h Repo) ListTransaction(ctx context.Context, id int) ([]TransactionWithAccount, error) {
	// defer timeTrack(time.Now(), "ListTransaction: ")

	items, err := h.DB.Connection.Query(ctx, `
		SELECT 
			c.id, c.name, c.acc_limit, c.balance, 
			t.id, t.value, t.type, t.description, t.created_at
		FROM client_account c 
		LEFT JOIN (
			SELECT id, account_id, value, type, description, created_at FROM transaction
			WHERE account_id = $1
			ORDER BY created_at DESC
			LIMIT 10
		) t ON c.id = $1
		WHERE c.id = $1;
	`, id)

	if err != nil {
		return nil, err
	}

	defer items.Close()

	var list []TransactionWithAccount

	for items.Next() {
		var t TransactionWithAccount

		if err := items.Scan(
			&t.AccountId,
			&t.AccountName,
			&t.AccountLimit,
			&t.AccountBalance,
			&t.TransactionId,
			&t.Value,
			&t.Type,
			&t.Description,
			&t.CreatedAt,
		); err != nil {
			return nil, err
		}

		list = append(list, t)

	}

	return list, nil
}

type TransactionWithAccount struct {
	AccountId      int
	AccountName    string
	AccountLimit   int
	AccountBalance int

	TransactionId *int
	Value         *int
	Type          *string
	Description   *string
	CreatedAt     *time.Time
}
