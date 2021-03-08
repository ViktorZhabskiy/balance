package model

import "time"

type TransactionType int

const (
	CurrencyEur = "EUR"
	CurrencyUsd = "USD"

	TransactionTypeDeposit    = "deposit"
	TransactionTypeWithdrawal = "withdrawal"

	// convert float (for example: 0.333) to int value multiplied by 1000 (eg. 333)
	BalancePrecision = 1000

	Deposit    TransactionType = 1
	Withdrawal TransactionType = 2
)

type Balance struct {
	Id       int64   `db:"id"`
	UserId   int64   `db:"user_id"`
	Balance  float64 `db:"balance"`
	Currency string  `db:"currency"`
}

type BalanceTransaction struct {
	Id              int64           `db:"id"`
	UserId          int64           `db:"user_id"`
	BalanceId       int64           `db:"balance_id"`
	BalanceBefore   float64         `db:"balance_before"`
	BalanceAfter    float64         `db:"balance_after"`
	Currency        string          `db:"currency, omitempty"`
	TransactionType TransactionType `db:"transaction_type"`
	TimePlaced      time.Time       `db:"time_placed"`
}
