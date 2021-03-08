package dto

import (
	"balance/internal/model"
	"balance/internal/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type BalanceTransactionRequest struct {
	UserId     string  `json:"user_id"`
	Currency   string  `json:"currency"`
	Amount     float64 `json:"amount"`
	TimePlaced string  `json:"time_placed"`
	Type       string  `json:"type"`
}

func (btr *BalanceTransactionRequest) Bind(r *http.Request) error {
	return nil
}

func (btr *BalanceTransactionRequest) Validate() error {
	if !utils.IsAllowedCurrency(btr.Currency) {
		return fmt.Errorf("Unknown currency %s ", btr.Currency)
	}

	if btr.UserId == "" {
		return fmt.Errorf("User id is empty ")
	}

	_, err := strconv.Atoi(btr.UserId)
	if err != nil {
		return fmt.Errorf("User id invalid ")
	}

	if btr.Amount < 0 {
		return fmt.Errorf("Amount can`t be negative ")
	}

	_, err = time.Parse(fmt.Sprintf("%s-%s-%s %s", "02", "JAN", "06", "15:04:05"), btr.TimePlaced)
	if err != nil {
		return fmt.Errorf("Invalid transaction time placed. Err %s ", err)
	}

	if !utils.IsAllowedTransactionType(btr.Type) {
		return fmt.Errorf("Unknown transaction type %s ", btr.Type)
	}

	return nil
}

type UserBalanceRequest struct {
	UserId int64 `json:"user_id"`
}

func (ubr *UserBalanceRequest) Bind(r *http.Request) error {
	return nil
}

type UserBalanceResponse struct {
	Balance []UserBalance `json:"balance"`
}

type UserBalance struct {
	Id       int64   `json:"id"`
	UserId   int64   `json:"user_id"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
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
