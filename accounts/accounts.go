package accounts

import "errors"

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

// Deposit x amount on your account
func (a *bankAccount) Deposit(amount int) {
	a.balance += amount
}

// Balance of an account
func (a bankAccount) Balance() int {
	return a.balance
}

// Withdraw from an account
func (a *bankAccount) Withdraw(amount int) {
	a.balance -= amount
}

// Withdraw2 from an account
func (a *bankAccount) Withdraw2(amount int) error {
	if a.balance < amount {
		return errors.New("Can't withdraw you are poor :(")
	}
	a.balance -= amount
	return nil
}
