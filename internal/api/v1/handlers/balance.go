package handlers

import (
	"balance/internal/api/v1/dto"
	"balance/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type BalanceHandler struct {
	balanceSrv service.UserBalance
}

func NewBalanceHandler(balanceSrv service.UserBalance) BalanceHandler {
	return BalanceHandler{
		balanceSrv: balanceSrv,
	}
}
func (h *BalanceHandler) GetUserBalance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//data := &dto.UserBalanceRequest{}
		//if err := render.Bind(r, data); err != nil {
		//	render.Render(w, r, &dto.ErrResponse{
		//		Err:            err,
		//		HTTPStatusCode: 400,
		//		StatusText:     "Invalid request.",
		//		ErrorText:      err.Error(),
		//	})
		//	return
		//}

		userId := chi.URLParam(r, "userId")
		if userId == "" {
			logrus.WithField("user_id", userId).Errorf("Unknown user id in request")
			render.Render(w, r, &dto.ErrResponse{
				HTTPStatusCode: 400,
				ErrorText:      "User balance not found",
			})
			return
		}

		userIdInt, err := strconv.Atoi(userId)
		if err != nil {
			logrus.Errorf("User id not valid. Err: %s", err)
			render.Render(w, r, &dto.ErrResponse{
				Err:            err,
				HTTPStatusCode: 400,
				ErrorText:      err.Error(),
			})
			return
		}

		balance, err := h.balanceSrv.Get(int64(userIdInt))
		if err != nil {
			logrus.Errorf("Error get user balance. Reason: %s", err)
			render.Render(w, r, &dto.ErrResponse{
				Err:            err,
				HTTPStatusCode: 400,
				ErrorText:      err.Error(),
			})
			return
		}
		userBalance := &dto.UserBalanceResponse{}
		userBalance.ToDTO(balance)

		render.Status(r, http.StatusOK)
		render.Render(w, r, userBalance)
	}
}
