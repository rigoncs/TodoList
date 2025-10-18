package repository

import (
	"context"
	"github.com/rigoncs/TodoList/domain/user/entity"
)

type User interface {
	UserBase
}

type UserBase interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUserByName(ctx context.Context, username string) (*entity.User, error)
	GetUserById(ctx context.Context, id uint64) (*entity.User, error)
}
