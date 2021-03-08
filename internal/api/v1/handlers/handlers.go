package handlers

import (
	"balance/internal/service"
)

type Handlers struct {
	BalanceHandler BalanceHandler
}

func NewHandlers(balanceSrv service.UserBalance) Handlers {
	return Handlers{
		BalanceHandler: NewBalanceHandler(balanceSrv),
	}
}
