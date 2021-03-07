package handlers

import "balance/internal/service"

type BalanceHandler struct {
	balanceSrv service.IBalance
}

func NewBalanceHandler(balanceSrv service.IBalance) BalanceHandler {
	return BalanceHandler{
		balanceSrv: balanceSrv,
	}
}
