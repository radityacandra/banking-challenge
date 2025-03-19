package model

import "time"

type UserAccount struct {
	Id           string  `db:"id"`
	UserId       string  `db:"user_id"`
	AccountNo    string  `db:"account_number"`
	TotalBalance int64   `db:"total_balance"`
	CreatedAt    int64   `db:"created_at"`
	CreatedBy    string  `db:"created_by"`
	UpdatedAt    *int64  `db:"updated_at"`
	UpdatedBy    *string `db:"updated_by"`
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
