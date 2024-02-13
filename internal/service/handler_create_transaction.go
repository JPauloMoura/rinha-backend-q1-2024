package service

import (
	"errors"
	"log/slog"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/entities"
	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/repository"
)

type CreateTransactionResponse struct {
	Limit int `json:"limite"`
	Saldo int `json:"saldo"`
}

func CreateTransaction(clientId int, t entities.Transaction) (*CreateTransactionResponse, error) {
	// 'c' -> adicionat o valor ao saldo
	// 'd' -> descontar o valor do saldo
	// o limite é apenas um valor de consulta para as  validações

	// buscar as informaçõe do client
	client, err := repository.FindClient(clientId)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	switch t.Type {
	case "c":
		client.Saldo += t.Value
	case "d":
		client.Saldo -= t.Value
	}

	if !client.SaldoIsValid() {
		slog.Error("new saldo is invalid")
		return nil, errors.New("transaction invalid")
	}

	t.ClientId = client.Id
	err = repository.CreateTransaction(client, t)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return &CreateTransactionResponse{
		Limit: client.Limit,
		Saldo: client.Saldo,
	}, nil
}
