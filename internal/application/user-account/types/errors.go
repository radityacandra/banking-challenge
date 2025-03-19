package types

import "errors"

var (
	ErrFailedToAcquireTx   = errors.New("failed to acquire transaction key")
	ErrInsufficientBalance = errors.New("user account balance is insufficient to perform such amount")
)
