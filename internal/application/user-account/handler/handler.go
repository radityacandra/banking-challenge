package handler

import (
	"github.com/radityacandra/banking-challenge/internal/application/user-account/repository"
	"github.com/radityacandra/banking-challenge/internal/application/user-account/service"
	"github.com/radityacandra/banking-challenge/internal/core"
	"go.uber.org/zap"
)

type Handler struct {
	Logger  *zap.Logger
	Service service.IService
}

func NewHandler(deps *core.Dependency) *Handler {
	repository := repository.NewRepository(deps.DB)
	service := service.NewService(repository)

	return &Handler{
		Logger:  deps.Logger,
		Service: service,
	}
}
