package app

import (
	"fmt"
	"hello-go/core/code"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

type AppGin struct {
	C         *gin.Context
	respCount int32
	trace     int64
}

func NewAppGin(c *gin.Context) *AppGin {
	now := time.Now()
	return &AppGin{
		C:     c,
		trace: now.UnixNano(),
	}
}

type bindFun func(c *gin.Context, req interface{}) error

func BindJson(path string, handler interface{}) (string, gin.HandlerFunc) {
	return bindGinHandler(path, handler, func(c *gin.Context, req interface{}) error {
		return c.BindJSON(req)
	})
}

func BindQuery(path string, handler interface{}) (string, gin.HandlerFunc) {
	return bindGinHandler(path, handler, func(c *gin.Context, req interface{}) error {
		return c.BindQuery(req)
	})
}

func BindUriAndQuery(path string, handler interface{}) (string, gin.HandlerFunc) {
	return bindGinHandler(path, handler, func(c *gin.Context, req interface{}) error {
		err := c.BindQuery(req)
		if err != nil {
			return err
		}
		return c.BindUri(req)
	})
}

func bindGinHandler(path string, handler interface{}, f bindFun) (string, gin.HandlerFunc) {
	hType := reflect.TypeOf(handler)
	if hType.Kind() != reflect.Func {
		panic("handler must be a function")
	}

	errType := reflect.TypeOf((*error)(nil)).Elem()

	if hType.NumOut() != 1 || hType.Out(0) != errType {
		panic("handler must return only one error")
	}

	if hType.In(0) != reflect.TypeOf(&AppGin{}) {
		panic("handler first param must be *AppGin")
	}
	reqType := hType.In(1)
	if reqType.Kind() != reflect.Ptr && reqType.Kind() != reflect.Struct {
		panic(fmt.Sprintf("handler second param must be a pointer or struct, but got %s", reqType.Kind().String()))
	}

	return path, newGinHandler(handler, reqType, f)
}

func newGinHandler(handler interface{}, reqType reflect.Type, f bindFun) gin.HandlerFunc {
	return func(c *gin.Context) {
		a := NewAppGin(c)
		isPtr := reqType.Kind() == reflect.Ptr
		// 绑定请求参数
		var reqVal reflect.Value
		if isPtr {
			reqVal = reflect.New(reqType.Elem())
		} else {
			reqVal = reflect.New(reqType)
		}
		err := f(c, reqVal.Interface())
		if err != nil {
			a.ErrResp(code.PARAM_INVALID)
			return
		}
		// 调用handler
		var results []reflect.Value
		if isPtr {
			results = reflect.ValueOf(handler).Call([]reflect.Value{reflect.ValueOf(a), reqVal})
		} else {
			results = reflect.ValueOf(handler).Call([]reflect.Value{reflect.ValueOf(a), reqVal.Elem()})
		}

		// 方法中断异常响应
		if len(results) == 0 {
			a.ErrResp(code.SERVER_ERR)
			return
		}
		// 无数据成功响应
		if results[0].IsNil() {
			a.SuccResp(struct{}{})
			return
		}

		// 内部错误响应
		err = results[0].Interface().(error)
		val, ok := err.(code.AppCode)
		if ok {
			a.ErrResp(val)
			return
		}
		// 非内部错误响应
		a.ErrResp(code.AppCode{Code: code.UNKONW_ERR.Code, Msg: fmt.Sprintf(code.UNKONW_ERR.Msg, err.Error())})
	}
}
