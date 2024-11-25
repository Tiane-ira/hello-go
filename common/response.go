package common

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func Success() *Response {
	return NewResponse(200, "success", nil)
}

func SuccessWithData(data interface{}) *Response {
	return NewResponse(200, "success", data)
}

func Failed() *Response {
	return NewResponse(-1, "failed", nil)
}

func FailedWithMsg(msg string) *Response {
	return NewResponse(-1, msg, nil)
}
