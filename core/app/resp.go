package app

import (
	"hello-go/core/code"
	"net/http"
	"sync/atomic"
)

type RespBase struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Trace  int64       `json:"trace"`
	Detail string      `json:"detail"`
}

func (a *AppGin) R(data interface{}) error {
	a.SuccResp(data)
	return nil
}

func (a *AppGin) SuccResp(data interface{}) {
	// 确保只调用一次响应逻辑
	if a.C == nil || !atomic.CompareAndSwapInt32(&a.respCount, 0, 1) {
		return
	}
	resp := RespBase{
		Code:  code.SUCCESS.Code,
		Msg:   code.SUCCESS.Msg,
		Data:  data,
		Trace: a.trace,
	}
	a.jsonResp(resp)
}

func (a *AppGin) jsonResp(resp RespBase) {
	a.C.JSON(http.StatusOK, resp)
}

func (a *AppGin) ErrResp(code code.AppCode) {
	if a.C == nil || !atomic.CompareAndSwapInt32(&a.respCount, 0, 1) {
		return
	}
	resp := RespBase{
		Code:   code.Code,
		Msg:    code.Msg,
		Data:   nil,
		Trace:  a.trace,
		Detail: code.Error(),
	}
	a.jsonResp(resp)
}
