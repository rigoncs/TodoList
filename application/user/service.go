package user

import (
	"context"
	"errors"
	"github.com/rigoncs/TodoList/domain/user/entity"
	"github.com/rigoncs/TodoList/domain/user/service"
	"github.com/rigoncs/TodoList/infrastructure/auth"
	"github.com/rigoncs/TodoList/interfaces/types"
	"sync"
)

type Service interface {
	Register(ctx context.Context, user *types.UserReq) (any, error)
	Login(ctx context.Context, user *entity.User) (any, error)
	GetUserInfo(ctx context.Context) (any, error)
}

type ServiceImpl struct {
	ud           service.UserDomain
	tokenService auth.TokenService
}

var (
	ServiceImplInst *ServiceImpl
	ServiceImplOnce sync.Once
)

func GetServiceImpl(svc service.UserDomain, jwt auth.TokenService) *ServiceImpl {
	ServiceImplOnce.Do(func() {
		ServiceImplInst = &ServiceImpl{
			ud:           svc,
			tokenService: jwt,
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
	return RegisterResponse(user), nil
}

func (s *ServiceImpl) Login(ctx context.Context, entity *entity.User) (any, error) {
	user, err := s.ud.FindUserByName(ctx, entity.Username)
	if err != nil {
		return nil, err
	}
	err = s.ud.CheckUserPwd(ctx, user, entity.Password)
	if err != nil {
		return nil, errors.New("invalid password")
	}
	token, err := s.tokenService.GenerateToken(ctx, uint(user.ID), user.Username)
	if err != nil {
		return nil, err
	}
	return LoginResponse(user, token), nil
}

func (s *ServiceImpl) GetUserInfo(ctx context.Context) (any, error) {
	return nil, nil
}
