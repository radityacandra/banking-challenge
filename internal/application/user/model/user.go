package model

import "time"

type UserAccount interface{}

type User struct {
	Id         string        `db:"id"`
	Name       string        `db:"name"`
	PhoneNo    string        `db:"phone_number"`
	IdentityNo string        `db:"identity_number"`
	CreatedAt  int64         `db:"created_at"`
	CreatedBy  string        `db:"created_by"`
	UpdatedAt  *int64        `db:"updated_at"`
	UpdatedBy  *string       `db:"updated_by"`
	Accounts   []UserAccount `db:"-"`
}

func NewUser(id, name, phoneNo, identityNo string, account UserAccount) *User {
	return &User{
		Id:         id,
		Name:       name,
		PhoneNo:    phoneNo,
		IdentityNo: identityNo,
		CreatedAt:  time.Now().UnixMilli(),
		CreatedBy:  id,
		Accounts:   []UserAccount{account},
	}
}
