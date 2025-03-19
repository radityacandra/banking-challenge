package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/banking-challenge/api"
	useraccount "github.com/radityacandra/banking-challenge/api/user-account"
)

func (h *Handler) AccountBalanceGet(ctx echo.Context, noRekening useraccount.AccountNumberParam) error {
	reqCtx := ctx.Request().Context()

	output, err := h.Service.GetBalance(reqCtx, noRekening)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, api.DefaultErrorResponse{
			Remarks: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, api.TransactionResponse{
		Saldo: output.Balance,
	})
}
