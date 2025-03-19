package types

import "errors"

var (
	ErrFailedToAcquireTx = errors.New("failed to acquire transaction key")
)
