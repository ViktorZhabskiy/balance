package service

import "balance/internal/repository"

type Balance struct {
	balanceRepo repository.IBalance
}

func NewBalance(balanceRepo repository.IBalance) IBalance {
	return &Balance{balanceRepo: balanceRepo}
}
