package repository

import "balance/internal/model"

type (
	UserBalance interface {
		Get(userId int64) ([]model.Balance, error)
		GetByType(userId int64, currency string) (model.Balance, error)
		PostTransaction(transaction model.BalanceTransaction) error
	}
)
