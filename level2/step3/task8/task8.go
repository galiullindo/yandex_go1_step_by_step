package main

import "errors"

var (
	Balance float64 = 0

	ErrAmountIsIncorrect  = errors.New("amount is incorrect")
	ErrBalanceIsIncorrect = errors.New("balance is incorrect")
)

func topUpBalance(amount float64) error {
	if amount <= 0 {
		return ErrAmountIsIncorrect
	}
	Balance += amount
	return nil
}

func chargeFromBalance(amount float64) error {
	if amount <= 0 {
		return ErrAmountIsIncorrect
	}
	Balance -= amount
	return nil
}

func getBalance() (float64, error) {
	if Balance < 0 {
		return 0, ErrBalanceIsIncorrect
	}
	return Balance, nil
}

func TopUpAndGetBalance(amount float64) (float64, error) {
	if err := topUpBalance(amount); err != nil {
		return 0, err
	}

	balance, err := getBalance()
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func ChargeFromAndGetBalance(amount float64) (float64, error) {
	if err := chargeFromBalance(amount); err != nil {
		return 0, err
	}

	balance, err := getBalance()
	if err != nil {
		return 0, err
	}

	return balance, nil
}
