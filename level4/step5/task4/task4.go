package main

import "errors"

var (
	ErrBalanceIsNegative  = errors.New("balance cannot be negative")
	ErrDepositIsNegative  = errors.New("deposit cannot be negative")
	ErrWithdrawIsNegative = errors.New("withdraw cannot be negative")
	ErrInsufficientFunds  = errors.New("insufficient funds")
)

type Account struct {
	balance float64
	owner   string
}

func NewAccount(owner string) *Account {
	return &Account{owner: owner, balance: 0}
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) SetBalance(value float64) error {
	if value < 0 {
		return ErrBalanceIsNegative
	}
	a.balance = value
	return nil
}

func (a *Account) Deposit(value float64) error {
	if value < 0 {
		return ErrDepositIsNegative
	}
	a.balance += value
	return nil
}

func (a *Account) insufficientFunds(value float64) bool {
	return a.balance-value < 0
}

func (a *Account) Withdraw(value float64) error {
	if value < 0 {
		return ErrWithdrawIsNegative
	}
	if a.insufficientFunds(value) {
		return ErrInsufficientFunds
	}
	a.balance -= value
	return nil
}
