package service

import (
	"context"
	"errors"
	"github.com/rigoncs/TodoList/domain/user/entity"
	"github.com/rigoncs/TodoList/domain/user/repository"
)

type UserDomain interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	FindUserByName(ctx context.Context, name string) (*entity.User, error)
	CheckUserPwd(ctx context.Context, user *entity.User, src string) error
}

type UserDomainImpl struct {
	repo    repository.User
	encrypt repository.PwdEncrypt
}

func NewUserDomainImpl(repo repository.User, encrypt repository.PwdEncrypt) UserDomain {
	return &UserDomainImpl{
		repo:    repo,
		encrypt: encrypt,
	}
}

func (u *UserDomainImpl) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	encryptPwd, err := u.encrypt.Encrypt([]byte(user.Password))
	if err != nil {
		return nil, err
	}
	err = user.SetPwd(encryptPwd)
	if err != nil {
		return nil, err
	}
	return u.repo.CreateUser(ctx, user)
}

func (u *UserDomainImpl) FindUserByName(ctx context.Context, name string) (*entity.User, error) {
	return u.repo.GetUserByName(ctx, name)
}

func (u *UserDomainImpl) CheckUserPwd(ctx context.Context, user *entity.User, src string) error {
	if u.encrypt.Check([]byte(user.Password), []byte(src)) {
		return nil
	}
	return errors.New("wrong password")
}
