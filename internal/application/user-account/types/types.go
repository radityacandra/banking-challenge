package types

type GetBalanceOutput struct {
	Balance   int64
	AccountNo string
}

type TransactionInput struct {
	AccountNo string
	Amount    int64
}
