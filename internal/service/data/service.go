package data

import (
	"context"

	"github.com/yuriy-kazanov/arch-go-example/internal/model"
)

type userService interface {
	ProcessUser(ctx context.Context, user model.User) error
}

type Service struct {
	userService
}
