package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/banking-challenge/api"
	"github.com/radityacandra/banking-challenge/internal/application/user/types"
	"github.com/radityacandra/banking-challenge/pkg/util"
)

func (h *Handler) UserPost(ctx echo.Context) error {
	var reqBody api.UserPostRequest
	if err := ctx.Bind(&reqBody); err != nil {
		return util.ReturnBadRequest(ctx, err, h.Logger)
	}

	if err := ctx.Validate(reqBody); err != nil {
		return util.ReturnBadRequest(ctx, err, h.Logger)
	}

	reqCtx := ctx.Request().Context()
	output, err := h.Service.RegisterUser(reqCtx, types.RegisterUserInput{
		Name:       reqBody.Nama,
		PhoneNo:    reqBody.NoHp,
		IdentityNo: reqBody.Nik,
	})
	if err != nil {
		return util.ReturnError(ctx, err, h.Logger)
	}

	return ctx.JSON(http.StatusOK, api.UserPostResponse{
		NoRekening: output.BankAccountNo,
	})
}
