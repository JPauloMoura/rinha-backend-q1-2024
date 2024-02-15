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
			DataExtrato: time.Now().In(date.LocationBR()).Format(date.DATE_BR_WITH_HOURS),
			Limite:      c.Limit,
		},
		UltimasTransacoes: make([]ExtractTransaction, 0),
	}

	if len(t) == 0 {
		return &ext
	}

	for _, v := range t {
		e := ExtractTransaction{
			Valor:       v.Value,
			Tipo:        v.Type,
			Descricao:   v.Description,
			RealizadaEm: v.CreatedAt.Format(date.DATE_BR_WITH_HOURS),
		}

		ext.UltimasTransacoes = append(ext.UltimasTransacoes, e)
	}

	return &ext
}

type ExtractTransaction struct {
	Valor       int    `json:"valor"`
	Tipo        string `json:"tipo"`
	Descricao   string `json:"descricao"`
	RealizadaEm string `json:"realizada_em"`
}

type ClientSaldo struct {
	Total       int    `json:"total"`
	DataExtrato string `json:"data_extrato"`
	Limite      int    `json:"limite"`
}

type Extract struct {
	Saldo             ClientSaldo          `json:"saldo"`
	UltimasTransacoes []ExtractTransaction `json:"ultimas_transacoes"`
}
