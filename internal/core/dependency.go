package core

import (
	"context"
	"errors"
	"syscall"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Dependency struct {
	Logger *zap.Logger
	DB     *sqlx.DB
	Config *Config
	Echo   *echo.Echo
}

func NewDependency(logger *zap.Logger, db *sqlx.DB, config *Config) *Dependency {
	return &Dependency{
		Logger: logger,
		DB:     db,
		Config: config,
	}
}

func (d *Dependency) GracefulShutdown(ctx context.Context) int {
	<-ctx.Done()
	code := 0
	d.Logger.Info("Gracefully shutting down web server...")
	err := d.Echo.Shutdown(ctx)
	if err != nil {
		d.Logger.Error("failed to close server", zap.Error(err))
		code = 1
	} else {
		d.Logger.Info("web server shutted down")
	}

	d.Logger.Info("Gracefully shutting down db connection...")
	err = d.DB.Close()
	if err != nil {
		d.Logger.Error("failed to close database connection", zap.Error(err))
		code = 1
	} else {
		d.Logger.Info("success to close database connection")
	}

	err = d.Logger.Sync()
	if err != nil && !errors.Is(err, syscall.ENOTTY) {
		d.Logger.Error("failed to flush log", zap.Error(err))
		code = 1
	} else {
		d.Logger.Info("success to flush log")
	}

	return code
}
