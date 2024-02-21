package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientAccount_BalanceIsValid(t *testing.T) {
	tests := []struct {
		Limit   int
		Balance int
		valid   bool
	}{
		{Limit: 0, Balance: -1, valid: false},
		{Limit: 10, Balance: -11, valid: false},
		{Limit: 0, Balance: 10, valid: true},
		{Limit: 10, Balance: -10, valid: true},
	}
	for _, tt := range tests {
		t.Run("is valid", func(t *testing.T) {
			b := ClientAccount{
				Limit:   tt.Limit,
				Balance: tt.Balance,
			}
			assert.Equal(t, tt.valid, b.BalanceIsValid())
		})
	}
}

func TestClientAccount_SetNewBalanceByTransaction(t *testing.T) {
	tests := []struct {
		name        string
		account     ClientAccount
		transaction Transaction
		wantErr     bool
		wantAccount ClientAccount
	}{
		{
			name:        "should accept a debit equal to the limit",
			account:     ClientAccount{Balance: 0, Limit: 10},
			transaction: Transaction{Type: "d", Value: 10},
			wantAccount: ClientAccount{Balance: -10, Limit: 10},
			wantErr:     false,
		},
		{
			name:        "should refuse a debit transaction that exceeds the limit",
			account:     ClientAccount{Balance: 0, Limit: 10},
			transaction: Transaction{Type: "d", Value: 11},
			wantAccount: ClientAccount{Balance: 0, Limit: 10},
			wantErr:     true,
		},
		{
			name:        "should accept credit transactions",
			account:     ClientAccount{Balance: 0, Limit: 10},
			transaction: Transaction{Type: "c", Value: 100},
			wantAccount: ClientAccount{Balance: 100, Limit: 10},
			wantErr:     false,
		},
		{
			name:        "should refuse transactions that are not debit or credit",
			account:     ClientAccount{Balance: 0, Limit: 10},
			transaction: Transaction{Type: "x", Value: 11},
			wantAccount: ClientAccount{Balance: 0, Limit: 10},
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := ClientAccount{
				Balance: tt.account.Balance,
				Limit:   tt.account.Limit,
			}

			err := b.SetNewBalanceByTransaction(tt.transaction)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantAccount.Balance, b.Balance)
			assert.Equal(t, tt.wantAccount.Limit, b.Limit)
		})
	}
}
