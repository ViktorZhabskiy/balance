package utils

import "balance/internal/model"

var allowedCurrencies = map[string]struct{}{model.CurrencyEur: {}, model.CurrencyUsd: {}}
var allowedTransactionType = map[string]struct{}{model.TransactionTypeDeposit: {}, model.TransactionTypeWithdrawal: {}}

// IsAllowedCurrency validate is system work with current currency
func IsAllowedCurrency(currency string) bool {
	if _, ok := allowedCurrencies[currency]; ok {
		return true
	}
	return false
}

// allowedTransactionType validate is system work with current transaction type
func IsAllowedTransactionType(transactionType string) bool {
	if _, ok := allowedTransactionType[transactionType]; ok {
		return true
	}
	return false
}
