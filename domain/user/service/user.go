package service

import (
	"context"
	"github.com/rigoncs/TodoList/domain/user/entity"
)

type UserDomain interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	FindUserByName(ctx context.Context, name string) (*entity.User, error)
}

type UserDomainImpl struct{}

func NewUserDomainImpl() UserDomain {
	return &UserDomainImpl{}
}

func (u *UserDomainImpl) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return user, nil
}

func (u *UserDomainImpl) FindUserByName(ctx context.Context, name string) (*entity.User, error) {
	return nil, nil
}
