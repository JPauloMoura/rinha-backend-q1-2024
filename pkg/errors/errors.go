package errors

import "errors"

var (
	ErrClientNotFound     = errors.New("client not found")
	ErrInvalidTransaction = errors.New("invalid transaction")
)
