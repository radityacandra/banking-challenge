package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/banking-challenge/api"
	"github.com/radityacandra/banking-challenge/internal/application/user-account/types"
	"github.com/radityacandra/banking-challenge/pkg/util"
)

func (h *Handler) TransactionDebitPost(ctx echo.Context) error {
	var reqBody api.TransactionRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return util.ReturnBadRequest(ctx, err, h.Logger)
	}

	if err := ctx.Validate(reqBody); err != nil {
		return util.ReturnBadRequest(ctx, err, h.Logger)
	}

	reqCtx := ctx.Request().Context()
	result, err := h.Service.WithdrawCash(reqCtx, types.TransactionInput{
		AccountNo: reqBody.NoRekening,
		Amount:    reqBody.Nominal,
	})
	if err != nil {
		return util.ReturnError(ctx, err, h.Logger)
	}

	return ctx.JSON(http.StatusOK, api.TransactionResponse{
		Saldo: result.Balance,
	})
}
