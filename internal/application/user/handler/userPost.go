package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/banking-challenge/api"
	"github.com/radityacandra/banking-challenge/internal/application/user/types"
)

func (h *Handler) UserPost(ctx echo.Context) error {
	var reqBody api.UserPostRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, api.DefaultErrorResponse{
			Remarks: err.Error(),
		})
	}

	// TODO: perform validator

	reqCtx := ctx.Request().Context()
	output, err := h.Service.RegisterUser(reqCtx, types.RegisterUserInput{
		Name:       reqBody.Nama,
		PhoneNo:    reqBody.NoHp,
		IdentityNo: reqBody.Nik,
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, api.DefaultErrorResponse{
			Remarks: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, api.UserPostResponse{
		NoRekening: output.BankAccountNo,
	})
}
