package service

import (
	"balance/internal/model"
	"balance/internal/repository"
	"balance/internal/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

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
		balance[i].Balance /= model.BalancePrecision
	}
	return balance, err
}

type Transaction struct {
	UserId     string
	Currency   string
	Amount     float64
	TimePlaced string
	Type       string
}

func (b *Balance) PostTransaction(transaction Transaction) error {
	// ignore error, validate before
	timePlaced, _ := time.Parse(fmt.Sprintf("%s-%s-%s %s", "02", "JAN", "06", "15:04:05"), transaction.TimePlaced)
	userId, _ := strconv.Atoi(transaction.UserId)

	balance, err := b.balanceRepo.GetByType(int64(userId), transaction.Currency)
	if err != nil {
		logrus.Errorf("Failed get balance by type. Err %s", err)
		return err
	}
	balance.Balance /= model.BalancePrecision

	balanceTransaction := model.BalanceTransaction{
		UserId:          int64(userId),
		Currency:        transaction.Currency,
		TimePlaced:      timePlaced,
		BalanceId:       balance.Id,
		BalanceBefore:   balance.Balance,
		BalanceAfter:    utils.CalculateBalance(balance.Balance, transaction.Amount, utils.GetTransactionType(transaction.Type)),
		TransactionType: utils.GetTransactionType(transaction.Type),
	}

	return b.balanceRepo.PostTransaction(balanceTransaction)
}
