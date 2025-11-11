package user

import (
	"context"

	"github.com/yuriy-kazanov/arch-go-example/internal/model"
)

type UserRepo interface {
	SaveUser(user model.User) error
}

type Service struct {
	repo UserRepo
}

func (s *Service) ProcessUser(ctx context.Context, user model.User) error {
	return s.repo.SaveUser(user)
}
