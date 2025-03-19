package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/radityacandra/banking-challenge/internal/application/user-account/types"
)

type UserAccount struct {
	Id                   string               `db:"id"`
	UserId               string               `db:"user_id"`
	AccountNo            string               `db:"account_number"`
	TotalBalance         int64                `db:"total_balance"`
	CreatedAt            int64                `db:"created_at"`
	CreatedBy            string               `db:"created_by"`
	UpdatedAt            *int64               `db:"updated_at"`
	UpdatedBy            *string              `db:"updated_by"`
	TransactionHistories []TransactionHistory `db:"-"`
}

type TransactionType string

const (
	TransactionTypeCredit TransactionType = "CREDIT"
	TransactionTypeDebit  TransactionType = "DEBIT"
)

type TransactionHistory struct {
	Id              string          `db:"id"`
	UserAccountId   string          `db:"user_account_id"`
	TransactionType TransactionType `db:"transaction_type"`
	Amount          int64           `db:"amount"`
	CreatedAt       int64           `db:"created_at"`
	CreatedBy       string          `db:"created_by"`
	UpdatedAt       *int64          `db:"updated_at"`
	UpdatedBy       *string         `db:"updated_by"`
}

func NewUserAccount(id, userId, accountNo string) *UserAccount {
	return &UserAccount{
		Id:        id,
		UserId:    userId,
		AccountNo: accountNo,
		CreatedAt: time.Now().UnixMilli(),
		CreatedBy: userId,
	}
}

func (ua *UserAccount) TransactionCredit(amount int64) *TransactionHistory {
	ua.TotalBalance += amount
	now := time.Now().UnixMilli()
	ua.UpdatedAt = &now
	ua.UpdatedBy = &ua.UserId

	transactionHistory := TransactionHistory{
		Id:              uuid.NewString(),
		UserAccountId:   ua.Id,
		TransactionType: TransactionTypeCredit,
		Amount:          amount,
		CreatedAt:       time.Now().UnixMilli(),
		CreatedBy:       ua.UserId,
	}

	ua.TransactionHistories = append(ua.TransactionHistories, transactionHistory)

	return &transactionHistory
}

func (ua *UserAccount) TransactionDebit(amount int64) (*TransactionHistory, error) {
	if amount > ua.TotalBalance {
		return nil, types.ErrInsufficientBalance
	}

	ua.TotalBalance -= amount
	now := time.Now().UnixMilli()
	ua.UpdatedAt = &now
	ua.UpdatedBy = &ua.UserId

	transactionHistory := TransactionHistory{
		Id:              uuid.NewString(),
		UserAccountId:   ua.Id,
		TransactionType: TransactionTypeDebit,
		Amount:          amount,
		CreatedAt:       time.Now().UnixMilli(),
		CreatedBy:       ua.UserId,
	}

	ua.TransactionHistories = append(ua.TransactionHistories, transactionHistory)

	return &transactionHistory, nil
}
