package types

type UserReq struct {
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type TokenData struct {
	User  any    `json:"user"`
	Token string `json:"token"`
}

type UserResp struct {
	ID        uint   `json:"id" form:"id"`
	UserName  string `json:"user_name" form:"user_name"`
	CreatedAt int64  `json:"created_at" form:"created_at"`
}
