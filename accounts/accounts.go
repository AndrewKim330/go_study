package accounts

// 220828

// BankAccount struct
type bankAccount struct {
	owner   string
	balance int
}

// NewAccount creates account
func NewAccount(owner string) *bankAccount {
	account := bankAccount{owner: owner, balance: 0}
	return &account
}
