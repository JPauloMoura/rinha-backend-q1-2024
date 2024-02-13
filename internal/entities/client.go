package entities

type Client struct {
	Id    int    `json:"id"`
	Name  string `json:"nome"`
	Limit int    `json:"limite"`
	Saldo int    `json:"saldo"`
}

func (c Client) SaldoIsValid() bool {
	return c.Saldo >= (c.Limit * -1)
}
