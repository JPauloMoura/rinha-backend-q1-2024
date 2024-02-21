package entities

import (
	"github.com/JPauloMoura/rinha-backend-q1-2024/pkg/errors"
)

type ClientAccount struct {
	Id         int    `json:"id"`
	Name       string `json:"nome"`
	Limit      int    `json:"limite"`
	Balance    int    `json:"saldo"`
	Generation int    `json:"-"`
}

func (c ClientAccount) BalanceIsValid() bool {
	return c.Balance >= (c.Limit * -1)
}

func (b *ClientAccount) SetNewBalanceByTransaction(t Transaction) error {
	newBalance := b.Balance

	switch t.Type {
	case "c": // crédito
		newBalance += t.Value
	case "d": // debito
		newBalance -= t.Value
	default:
		return errors.ErrInvalidTransaction
	}

	if newBalance < (b.Limit * -1) {
		return errors.ErrInvalidTransaction
	}

	b.Balance = newBalance

	return nil
}
