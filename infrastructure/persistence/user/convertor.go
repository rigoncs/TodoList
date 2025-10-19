package user

import "github.com/rigoncs/TodoList/domain/user/entity"

func Entity2PO(user *entity.User) *User {
	return &User{
		UserName:       user.Username,
		PasswordDigest: user.Password,
	}
}

func PO2Entity(user *User) *entity.User {
	return &entity.User{
		ID:        uint64(user.ID),
		Username:  user.UserName,
		Password:  user.PasswordDigest,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
