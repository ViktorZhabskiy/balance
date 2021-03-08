package repository

import (
	"balance/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Balance struct {
	db *sqlx.DB
}

func NewBalance(db *sqlx.DB) UserBalance {
	return &Balance{db: db}
}

func (b *Balance) Get(userId int64) ([]model.Balance, error) {
	balance := make([]model.Balance, 0)
	err := b.db.Select(&balance, "SELECT users_balance.id as id, user_id, balance, currency.name as currency FROM users_balance INNER JOIN currency ON currency.id = users_balance.currency_id WHERE user_id = $1", userId)
	if err != nil {
		logrus.Warn("Failed get user balance from db, err %s", err)
		return balance, err
	}
	return balance, nil
}
