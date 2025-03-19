package service

import (
	"context"

	"github.com/radityacandra/banking-challenge/internal/application/user/repository"
	"github.com/radityacandra/banking-challenge/internal/application/user/types"
)

type IService interface {
	RegisterUser(ctx context.Context, input types.RegisterUserInput) (types.RegisterUserOutput, error)
}

type Service struct {
	Repository repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		Repository: repository,
	}
}
