package repository

import (
	"balance/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
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

func (b *Balance) GetByType(userId int64, currency string) (model.Balance, error) {
	balance := model.Balance{}
	err := b.db.Get(&balance, "SELECT id, user_id, balance FROM users_balance WHERE user_id = $1 AND currency_id = (SELECT id FROM currency WHERE name = $2)", userId, currency)
	if err != nil {
		logrus.Warn("Failed get user balance from db, err %s", err)
		return balance, err
	}
	return balance, nil
}

func (b *Balance) PostTransaction(transaction model.BalanceTransaction) error {
	tx, err := b.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO users_balance_transactions (user_id, balance_id, balance_before, balance_after, transaction_type, time_placed) VALUES ($1, $2, $3, $4, $5, $6)",
		transaction.UserId,
		transaction.BalanceId,
		transaction.BalanceBefore*model.BalancePrecision,
		transaction.BalanceAfter*model.BalancePrecision,
		transaction.TransactionType,
		transaction.TimePlaced)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE users_balance SET balance = $1, updated_at = $2 WHERE id = $3", transaction.BalanceAfter*model.BalancePrecision, time.Now(), transaction.BalanceId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
