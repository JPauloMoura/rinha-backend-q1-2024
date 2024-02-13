package entities

import (
	"errors"
	"time"
)

type Transaction struct {
	Id          int       `json:"id"`
	ClientId    int       `json:"clientId"`
	Value       int       `json:"valor"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	CreatedAt   time.Time `json:"realizada_em"`
}

func (t Transaction) Validate() error {
	if t.Value <= 0 {
		return errors.New("value should be > 0")
	}

	if t.Type != "c" && t.Type != "d" {
		return errors.New("type should be 'c' or 'd'")
	}

	if l := len(t.Description); l < 1 || l > 10 {
		return errors.New("description deve ter entre 1 e 10 caracteres")
	}

	return nil
}
