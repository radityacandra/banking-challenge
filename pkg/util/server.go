package util

import (
	"errors"
	"net/http"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/banking-challenge/api"
	"github.com/radityacandra/banking-challenge/internal/application/user-account/types"
	"go.uber.org/zap"
)

var errorMap = map[error]int{
	types.ErrFailedToAcquireTx:   http.StatusBadRequest,
	types.ErrInsufficientBalance: http.StatusBadRequest,
	types.ErrUserAccountNotFound: http.StatusBadRequest,
}

const unknownError string = "unknown error"

func ReturnError(ctx echo.Context, err error, logger *zap.Logger) error {
	var code int
	var message string
	for registeredErr, registeredErrCode := range errorMap {
		if errors.Is(err, registeredErr) {
			code = registeredErrCode
			message = registeredErr.Error()
		}
	}

	if code == 0 {
		code = http.StatusInternalServerError
	}

	pc, file, line, _ := runtime.Caller(1)

	if code == http.StatusInternalServerError {
		logger.Error("error occured",
			zap.Error(err),
			zap.Any("invoker", runtime.FuncForPC(pc).Name()),
			zap.Any("file", file), zap.Any("line", line))

		message = unknownError
	} else {
		logger.Warn("responding with client error",
			zap.Error(err),
			zap.Any("invoker", runtime.FuncForPC(pc).Name()),
			zap.Any("file", file), zap.Any("line", line))
	}

	return ctx.JSON(code, api.DefaultErrorResponse{
		Remarks: message,
	})
}

func ReturnBadRequest(ctx echo.Context, err error, logger *zap.Logger) error {
	pc, file, line, _ := runtime.Caller(1)
	logger.Warn("responding with client error",
		zap.Error(err),
		zap.Any("invoker", runtime.FuncForPC(pc).Name()),
		zap.Any("file", file), zap.Any("line", line))

	return ctx.JSON(http.StatusBadRequest, api.DefaultErrorResponse{
		Remarks: err.Error(),
	})
}
