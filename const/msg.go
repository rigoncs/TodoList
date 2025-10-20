package _const

var MsgFlags = map[int]string{
	SUCCESS:                    "操作成功",
	ERROR:                      "操作失败",
	InvalidParams:              "参数错误",
	ErrorExistUser:             "用户已存在",
	ErrorNotExistUser:          "用户不存在",
	ErrorNotCompare:            "不匹配",
	ErrorFailEncryption:        "密码加密失败",
	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
	ErrorDatabase:              "数据库操作错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
