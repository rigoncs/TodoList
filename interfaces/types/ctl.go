package types

import _const "github.com/rigoncs/TodoList/const"

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	Error   string      `json:"error"`
	TraceID string      `json:"trace_id"`
}

func RespSuccess(code ...int) *Response {
	status := _const.SUCCESS
	if len(code) > 0 {
		status = code[0]
	}
	return &Response{
		Status: status,
		Data:   _const.MsgFlags[status],
		Msg:    _const.MsgFlags[status],
	}
}

func RespSuccessWithData(data interface{}, code ...int) *Response {
	status := _const.SUCCESS
	if len(code) > 0 {
		status = code[0]
	}
	return &Response{
		Status: status,
		Data:   data,
		Msg:    _const.MsgFlags[status],
	}
}

func RespError(err error, data string, code ...int) *Response {
	status := _const.ERROR
	if len(code) > 0 {
		status = code[0]
	}
	return &Response{
		Status: status,
		Data:   data,
		Msg:    _const.MsgFlags[status],
		Error:  err.Error(),
	}
}
