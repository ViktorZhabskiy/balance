package service

import (
	"balance/internal/model"
	"balance/internal/repository"
)

const BalancePrecision = 1000

type Balance struct {
	balanceRepo repository.UserBalance
}

func NewBalance(balanceRepo repository.UserBalance) UserBalance {
	return &Balance{balanceRepo: balanceRepo}
}

func (b *Balance) Get(userId int64) ([]model.Balance, error) {
	balance, err := b.balanceRepo.Get(userId)
	if err != nil {
		return balance, err
	}

	for i := 0; i < len(balance); i++ {
		balance[i].Balance /= BalancePrecision
	}
	return b.balanceRepo.Get(userId)
}
