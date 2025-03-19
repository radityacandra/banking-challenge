package handler

import (
	"github.com/radityacandra/banking-challenge/internal/core"
	"go.uber.org/zap"
)

type Handler struct {
	Logger *zap.Logger
}

func NewHandler(deps *core.Dependency) *Handler {
	return &Handler{
		Logger: deps.Logger,
	}
}
