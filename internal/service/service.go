package service

import "balance/internal/model"

type (
	UserBalance interface {
		Get(userId int64) ([]model.Balance, error)
	}
)
