package repository

import "github.com/yuriy-kazanov/arch-go-example/internal/model"

type UserRepo struct {
}

func (r *UserRepo) SaveUser(user model.User) error {
	return nil
}
