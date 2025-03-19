package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/radityacandra/banking-challenge/api/user"
	useraccount "github.com/radityacandra/banking-challenge/api/user-account"
	userAccountHandler "github.com/radityacandra/banking-challenge/internal/application/user-account/handler"
	"github.com/radityacandra/banking-challenge/internal/application/user/handler"
	"github.com/radityacandra/banking-challenge/internal/core"
	"github.com/radityacandra/banking-challenge/pkg/validator"
	"go.uber.org/zap"
)

func InitServer(ctx context.Context, deps *core.Dependency) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Validator = validator.NewValidator()

	e.Use(middleware.CORS())

	userHandler := handler.NewHandler(deps)
	user.RegisterHandlers(e, userHandler)

	userAccountHandler := userAccountHandler.NewHandler(deps)
	useraccount.RegisterHandlers(e, userAccountHandler)

	deps.Echo = e

	deps.Logger.Info("Web server ready", zap.Int("port", 8080))
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			deps.Logger.Fatal("Failed to start web server", zap.Error(err))
		}
	}()
}
