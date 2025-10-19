package user

import (
	"context"
	"errors"
	"github.com/rigoncs/TodoList/domain/user/entity"
	"github.com/rigoncs/TodoList/domain/user/repository"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepositoryImpl(db *gorm.DB) repository.User {
	return &RepositoryImpl{
		db: db,
	}
}

func (r *RepositoryImpl) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	u := Entity2PO(user)
	err := r.db.WithContext(ctx).Model(&User{}).Create(u).Error
	if err != nil {
		return nil, err
	}
	return PO2Entity(u), nil
}

func (r *RepositoryImpl) GetUserByName(ctx context.Context, username string) (*entity.User, error) {
	var u *User
	err := r.db.WithContext(ctx).Model(&User{}).Where("user_name = ?", username).Find(&u).Error
	if err != nil {
		return nil, err
	}
	return PO2Entity(u), nil
}

func (r *RepositoryImpl) GetUserById(ctx context.Context, id uint64) (*entity.User, error) {
	var u *User
	err := r.db.WithContext(ctx).Model(&User{}).Where("id = ?", id).Find(&u).Error
	if err != nil {
		return nil, err
	}
	if u.ID == 0 {
		return nil, errors.New("user not found")
	}
	return PO2Entity(u), nil
}
