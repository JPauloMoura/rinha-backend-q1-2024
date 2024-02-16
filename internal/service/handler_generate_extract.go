package service

import (
	"time"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/entities"
	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/repository"
	"github.com/JPauloMoura/rinha-backend-q1-2024/pkg/date"
)

func GenerateExtract(clientId int) (*Extract, error) {
	client, err := repository.FindClient(clientId)
	if err != nil {
		return nil, err
	}

	transactions, err := repository.ListTransaction(client.Id)
	if err != nil {
		return nil, err
	}

	extract := NewExtract(*client, transactions)

	return extract, nil
}

func NewExtract(c entities.Client, t []entities.Transaction) *Extract {
	ext := Extract{
		Saldo: ClientSaldo{
			Total:       c.Saldo,
			DataExtrato: time.Now().In(date.LocationBR()).String(),
			Limite:      c.Limit,
		},
		UltimasTransacoes: t,
	}

	return &ext
}

type ClientSaldo struct {
	Total       int    `json:"total"`
	DataExtrato string `json:"data_extrato"`
	Limite      int    `json:"limite"`
}

type Extract struct {
	Saldo             ClientSaldo            `json:"saldo"`
	UltimasTransacoes []entities.Transaction `json:"ultimas_transacoes"`
}
