package service

import (
	"context"

	"github.com/yuriy-kazanov/arch-go-example/internal/model"
)

type UserProcessor interface {
	ProcessUser(ctx context.Context, user model.User) error
}
