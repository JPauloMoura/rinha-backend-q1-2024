package repository

import (
	"context"
	"errors"
	"time"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/entities"
)

func ListTransaction(ctx context.Context, id int) ([]ExtractDB, error) {
	// defer timeTrack(time.Now(), "ListTransaction: ")

	items, err := DB.Query(ctx, `
		SELECT 
			c.id, c.nome, c.limite, c.saldo, 
			t.clientId, t.valor, t.tipo, t.descricao, t.realizada_em
		FROM clientes c
		LEFT JOIN (
			SELECT clientId, valor, tipo, descricao, realizada_em FROM transactions
			WHERE clientId = $1
			ORDER BY realizada_em DESC
			LIMIT 10
		) t ON c.id = $1
		WHERE c.id = $1;
	`, id)

	if err != nil {
		// slog.Error(err.Error())
		return nil, errors.New("internal error")
	}

	var e []ExtractDB

	for items.Next() {
		var t Resp

		if err := items.Scan(
			&t.ClientID,
			&t.ClientName,
			&t.ClientLimit,
			&t.ClientSaldo,
			&t.TransactionID,
			&t.Valor,
			&t.Tipo,
			&t.Descricao,
			&t.RealizadaEm,
		); err != nil {
			// slog.Error(err.Error())
			return nil, errors.New("internal error")
		}

		if t.TransactionID == nil {
			e = append(e, ExtractDB{
				entities.Client{
					Id:    *t.ClientID,
					Name:  *t.ClientName,
					Limit: *t.ClientLimit,
					Saldo: *t.ClientSaldo,
				}, nil},
			)

			continue
		}

		e = append(e, ExtractDB{
			entities.Client{
				Id:    *t.ClientID,
				Name:  *t.ClientName,
				Limit: *t.ClientLimit,
				Saldo: *t.ClientSaldo,
			},
			&entities.Transaction{
				ClientId:    *t.ClientID,
				Value:       *t.Valor,
				Type:        *t.Tipo,
				Description: *t.Descricao,
				CreatedAt:   *t.RealizadaEm,
			},
		})
	}

	return e, nil
}

type Resp struct {
	ClientID      *int
	ClientName    *string
	ClientLimit   *int
	ClientSaldo   *int
	TransactionID *int
	Valor         *int
	Tipo          *string
	Descricao     *string
	RealizadaEm   *time.Time
}

type ExtractDB struct {
	entities.Client
	*entities.Transaction
}
