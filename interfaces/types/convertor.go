package types

import "github.com/rigoncs/TodoList/domain/user/entity"

func UserReq2Entity(user *UserReq) *entity.User {
	return &entity.User{
		Username: user.UserName,
		Password: user.Password,
	}
}
