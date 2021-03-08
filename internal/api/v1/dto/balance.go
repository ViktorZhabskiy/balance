package dto

import (
	"balance/internal/model"
	"net/http"
)

type CreateBalanceTransactionRequest struct {
}

type UserBalanceRequest struct {
	UserId int64 `json:"user_id"`
}

func (ubr *UserBalanceRequest) Bind(r *http.Request) error {
	return nil
}

type UserBalanceResponse struct {
	Balance []UserBalance
}

type UserBalance struct {
	Id       int64  `json:"id"`
	UserId   int64  `json:"user_id"`
	Balance  int    `json:"balance"`
	Currency string `json:"currency"`
}

func (ubs *UserBalanceResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (ubs *UserBalanceResponse) ToDTO(balance []model.Balance) {
	ubs.Balance = make([]UserBalance, 0)
	for _, v := range balance {
		ubs.Balance = append(ubs.Balance, UserBalance{
			Id:       v.Id,
			UserId:   v.UserId,
			Balance:  v.Balance,
			Currency: v.Currency,
		})
	}
}
