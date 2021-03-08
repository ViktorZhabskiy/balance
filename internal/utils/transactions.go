package utils

import "balance/internal/model"

func GetTransactionType(tType string) model.TransactionType {
	switch tType {
	case "deposit":
		return model.Deposit
	case "withdrawal":
		return model.Withdrawal
	}
	return 0
}

func CalculateBalance(balance, amount float64, transactionType model.TransactionType) float64 {
	if transactionType == model.Deposit {
		return deposit(balance, amount)
	}

	return withdrawal(balance, amount)
}

func deposit(balance, amount float64) float64 {
	b, a := balance*model.BalancePrecision, amount*model.BalancePrecision
	return (b + a) / model.BalancePrecision
}

func withdrawal(balance, amount float64) float64 {
	b, a := balance*model.BalancePrecision, amount*model.BalancePrecision
	if b < a {
		return balance
	}
	return (b - a) / model.BalancePrecision
}
