package repository

import (
	"balance/internal/model"
	"database/sql"
)

type BalanceMock struct {
}

func (bm *BalanceMock) Get(userId int64) ([]model.Balance, error) {
	return []model.Balance{}, nil
}

func (bm *BalanceMock) GetByType(userId int64, currency string) (model.Balance, error) {
	if userId == 1 {
		return model.Balance{
			Id:       int64(1),
			UserId:   int64(1),
			Balance:  10.367,
			Currency: "USD",
		}, nil
	}

	return model.Balance{}, sql.ErrNoRows
}

func (bm *BalanceMock) PostTransaction(transaction model.BalanceTransaction) error {
	return nil
}
