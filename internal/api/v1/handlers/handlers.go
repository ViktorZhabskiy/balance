package handlers

import "balance/internal/service"

type Handlers struct {
	balanceHandler BalanceHandler
}

func NewHandlers(balanceSrv service.IBalance) Handlers {
	return Handlers{
		balanceHandler: NewBalanceHandler(balanceSrv),
	}
}
