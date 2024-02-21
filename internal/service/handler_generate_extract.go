package service

import (
	"context"
	"time"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/entities"
	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/repository"
)

type Service struct {
	Repo repository.Repo
}

func (s Service) GenerateExtract(ctx context.Context, id int) (*Extract, error) {
	transactions, err := s.Repo.ListTransaction(ctx, id)
	if err != nil {
		return nil, err
	}

	return constructExtract(transactions), nil
}

func constructExtract(transactions []repository.TransactionWithAccount) *Extract {
	if len(transactions) == 0 {
		return &Extract{}
	}

	ext := Extract{
		Balancer: BalanceUserInfo{
			Total:     transactions[0].AccountBalance,
			CreatedAt: time.Now().UTC().String(),
			Limit:     transactions[0].AccountLimit,
		},
		LatestTransactions: make([]entities.Transaction, 0),
	}

	for _, t := range transactions {
		if t.TransactionId != nil {
			ext.LatestTransactions = append(ext.LatestTransactions, entities.Transaction{
				Value:       *t.Value,
				Type:        *t.Type,
				Description: *t.Description,
				CreatedAt:   *t.CreatedAt,
			})
		}
	}

	return &ext
}

type Extract struct {
	Balancer           BalanceUserInfo        `json:"saldo"`
	LatestTransactions []entities.Transaction `json:"ultimas_transacoes"`
}

type BalanceUserInfo struct {
	Total     int    `json:"total"`
	CreatedAt string `json:"data_extrato"`
	Limit     int    `json:"limite"`
}
