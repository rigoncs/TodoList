package user

import (
	"github.com/rigoncs/TodoList/domain/user/entity"
	"github.com/rigoncs/TodoList/interfaces/types"
)

func LoginResponse(u *entity.User, token string) *types.TokenData {
	return &types.TokenData{
		User: types.UserResp{
			ID:        uint(u.ID),
			UserName:  u.Username,
			CreatedAt: u.CreatedAt.Unix(),
		},
		Token: token,
	}
}

func RegisterResponse(u *entity.User) *types.UserResp {
	return &types.UserResp{
		ID:        uint(u.ID),
		UserName:  u.Username,
		CreatedAt: u.CreatedAt.Unix(),
	}
}
