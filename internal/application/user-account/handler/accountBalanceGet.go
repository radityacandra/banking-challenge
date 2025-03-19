package handler

import (
	"github.com/labstack/echo/v4"
	useraccount "github.com/radityacandra/banking-challenge/api/user-account"
)

func (h *Handler) AccountBalanceGet(ctx echo.Context, noRekening useraccount.AccountNumberParam) error {
	return nil
}
