package repository

import (
	"context"

	"github.com/JPauloMoura/rinha-backend-q1-2024/internal/entities"
)

type CreateTransactionResponse struct {
	Balance int `json:"saldo"`
	Limit   int `json:"limite"`
}

func (h Repo) CreateTransaction(ctx context.Context, accId int, transaction entities.Transaction) (*CreateTransactionResponse, error) {
	var (
		balance       int
		limit         int
		generation    int
		clientAccount entities.ClientAccount
	)
	// Lock otimista
	// Buscamos as informações da conta juntamente com sua generation
	// e quando formos atualiza o novo salvo só atualizamos
	// se a generation que iremos atuliza for a mesma que realizamos a leitura
	// caso contrario retentemos novamente em cima da nova generation
	// como temos muita concorrencia, tentamos esse metódo umas 100x hehehe
	var errorLock error
	for retry := 0; retry < 100; retry++ {

		errorLock = h.DB.Connection.QueryRow(ctx, `
			SELECT balance, acc_limit, generation FROM client_account
			WHERE id = $1`, accId,
		).Scan(&balance, &limit, &generation)

		if errorLock != nil {
			// fmt.Println("select ", errorLock)
			break
		}

		clientAccount = entities.ClientAccount{
			Balance:    balance,
			Limit:      limit,
			Generation: generation,
		}

		errorLock = clientAccount.SetNewBalanceByTransaction(transaction)
		if errorLock != nil {
			// fmt.Println("set ", errorLock, clientAccount)
			break
		}

		cmd, errorLock := h.DB.Connection.Exec(ctx, `
			UPDATE client_account SET 
			balance = $1, generation = generation+1
			WHERE id = $2 and generation = $3`,
			clientAccount.Balance, accId, clientAccount.Generation,
		)

		// no update
		if cmd.RowsAffected() <= 0 || errorLock != nil {
			// fmt.Println("update ", errorLock, clientAccount)
			errorLock = nil
			continue
		}

		// fmt.Println("break ", errorLock, clientAccount, cmd.RowsAffected())

		break

		// fmt.Printf("client: %d v%d loopError:%d\n erro:%v", clientId, versao, retry, errorLock)
	}

	if errorLock != nil {
		// fmt.Println("error lock ", errorLock)
		return nil, errorLock
	}

	_, err := h.DB.Connection.Exec(ctx, `
		INSERT INTO transaction (account_id, value, type, description)
		VALUES ($1, $2, $3, $4)`,
		accId, transaction.Value, transaction.Type, transaction.Description,
	)

	if err != nil {
		// fmt.Println("error insert transaction ", err)
		return nil, err
	}

	return &CreateTransactionResponse{
		Balance: clientAccount.Balance,
		Limit:   clientAccount.Limit,
	}, nil
}
