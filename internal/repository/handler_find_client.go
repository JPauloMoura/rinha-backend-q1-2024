package repository

import (
	"context"
	"errors"
	"log/slog"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/entities"
)

func FindClient(id int) (*entities.Client, error) {
	item, err := DB.Query(context.TODO(), "SELECT * FROM clientes WHERE id = $1", id)
	if err != nil {
		slog.Error(err.Error())
		return nil, errors.New("internal error")
	}

	defer item.Close()

	var client entities.Client

	if !item.Next() {
		return nil, errors.New("client not found")
	}

	err = item.Scan(&client.Id, &client.Name, &client.Limit, &client.Saldo)
	if err != nil {
		slog.Error(err.Error())
		return nil, errors.New("internal error")
	}

	return &client, nil
}
