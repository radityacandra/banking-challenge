package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/banking-challenge/api"
	"github.com/radityacandra/banking-challenge/internal/application/user-account/types"
)

func (h *Handler) TransactionCreditPost(ctx echo.Context) error {
	var reqBody api.TransactionRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.DefaultErrorResponse{
			Remarks: err.Error(),
		})
	}

	if err := ctx.Validate(reqBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.DefaultErrorResponse{
			Remarks: err.Error(),
		})
	}

	reqCtx := ctx.Request().Context()
	result, err := h.Service.StoreCash(reqCtx, types.TransactionInput{
		AccountNo: reqBody.NoRekening,
		Amount:    reqBody.Nominal,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, api.DefaultErrorResponse{
			Remarks: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, api.TransactionResponse{
		Saldo: result.Balance,
	})
}
