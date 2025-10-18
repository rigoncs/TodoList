package user

import (
	"context"
	"errors"
	"github.com/rigoncs/TodoList/domain/user/entity"
	"github.com/rigoncs/TodoList/domain/user/service"
	"sync"
)

type Service interface {
	Login(ctx context.Context, user *entity.User) (any, error)
}

type ServiceImpl struct {
	ud service.UserDomain
}

var (
	ServiceImplInst *ServiceImpl
	ServiceImplOnce sync.Once
)

func GetServiceImpl(svc service.UserDomain) *ServiceImpl {
	ServiceImplOnce.Do(func() {
		ServiceImplInst = &ServiceImpl{
			ud: svc,
		}
	})
	return ServiceImplInst
}

func (s *ServiceImpl) Register(ctx context.Context, userEntity *entity.User) (any, error) {
	userExist, err := s.ud.FindUserByName(ctx, userEntity.Username)
	if err != nil {
		return nil, err
	}
	if userExist.IsActive() {
		return nil, errors.New("user is already active")
	}
	user, err := s.ud.CreateUser(ctx, userEntity)
	if err != nil {
		return nil, err
	}
	return user, nil
}
