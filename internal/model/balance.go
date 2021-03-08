package model

type Balance struct {
	Id       int64  `db:"id"`
	UserId   int64  `db:"user_id"`
	Balance  int    `db:"balance"`
	Currency string `db:"currency"`
}
