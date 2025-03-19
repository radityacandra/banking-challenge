package types

type RegisterUserInput struct {
	Name       string
	PhoneNo    string
	IdentityNo string
}

type RegisterUserOutput struct {
	UserId        string
	BankAccountNo string
}
