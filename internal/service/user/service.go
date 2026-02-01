package user

import (
	"context"
	"errors"
	"fmt"

	"go.opentelemetry.io/otel"

	"github.com/yuriy-kazanov/arch-go-example/internal/model"
)

var ErrNotFound = errors.New("user not found")

type UserRepo interface {
	GetUser(id string) (model.User, error)
	SaveUser(user model.User) error
}

type Service struct {
	repo UserRepo
}

func (s *Service) GetUser(id string) (model.User, error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		//if errors.Is(err, pgx.ErrNoRows) {
		//	return model.User{}, fmt.Errorf("no users found, %w", ErrNotFound)
		//}
		return model.User{}, fmt.Errorf("error getting user, %w", err)
	}

	return user, nil
}

func (s *Service) ProcessUser(ctx context.Context, user model.User) error {
	ctx, span := otel.Tracer("user").Start(ctx, "ProcessUser")
	defer span.End()

	return s.repo.SaveUser(user)
}
