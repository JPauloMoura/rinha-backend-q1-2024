package service

import (
	"errors"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/entities"
	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/repository"
)

type CreateTransactionResponse struct {
	Saldo int `json:"saldo"`
	Limit int `json:"limite"`
}

func CreateTransaction(clientId int, t entities.Transaction) (*CreateTransactionResponse, error) {
	client, err := repository.FindClient(clientId)
	if err != nil {
		return nil, err
	}

	switch t.Type {
	case "c":
		client.Saldo += t.Value
	case "d":
		client.Saldo -= t.Value
	}

	if !client.SaldoIsValid() {
		return nil, errors.New("transaction invalid")
	}

	t.ClientId = client.Id
	err = repository.CreateTransaction(client, t)
	if err != nil {
		return nil, err
	}

	return &CreateTransactionResponse{
		Saldo: client.Saldo,
		Limit: client.Limit,
	}, nil
}
