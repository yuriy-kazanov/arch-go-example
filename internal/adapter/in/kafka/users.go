package kafka

import (
	"context"

	"github.com/yuriy-kazanov/arch-go-example/internal/model"
)

type producer interface {
	SendMessage(ctx context.Context, user model.User) error
}

type Consumer struct {
	producer
}

func (c *Consumer) HandleMessage(ctx context.Context, userID string) error {
	return c.producer.SendMessage(ctx, model.User{
		ID: userID,
	})
}
