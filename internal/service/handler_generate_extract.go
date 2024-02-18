package service

import (
	"context"
	"time"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/entities"
	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/repository"
	"github.com/JPauloMoura/rinha-backend-q1-2024/pkg/date"
)

func GenerateExtract(ctx context.Context, clientId int) (*Extract, error) {
	exDb, err := repository.ListTransaction(ctx, clientId)
	if err != nil {
		return nil, err
	}

	extract := NewExtract(exDb)

	return extract, nil
}

func NewExtract(ex []repository.ExtractDB) *Extract {
	ext := Extract{
		Saldo: ClientSaldo{
			Total:       ex[0].Saldo,
			DataExtrato: time.Now().In(date.LocationBR()).String(),
			Limite:      ex[0].Limit,
		},
		UltimasTransacoes: make([]entities.Transaction, 0),
	}

	for _, t := range ex {
		if t.Transaction != nil {
			ext.UltimasTransacoes = append(ext.UltimasTransacoes, *t.Transaction)
		}
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
