package demo

import (
	"reflect"
	"testing"
)

func Test_1(t *testing.T) {
	u := &User{
		Id:   1,
		Name: "zs",
	}

	upType := reflect.TypeOf(u)
	if upType.Kind() == reflect.Ptr {
		t.Logf("指针类型")
	} else if upType.Kind() == reflect.Struct {
		t.Log("结构体类型")
	}
	t.Logf("指针类型:%v", upType)
	t.Logf("指针类型:%v", upType.Elem())

	// 获取结构体指针值
	pval := reflect.ValueOf(u)
	// 获取结构体值
	sval := pval.Elem()
	// 获取字段数
	fieldCount := sval.NumField()
	t.Logf("字段数:%d", fieldCount)
	// 设置字段值
	sval.Field(0).SetInt(2)
	t.Logf("操作结果:%v", u)
	// 获取字段名/类型/tag
	for i := 0; i < fieldCount; i++ {
		field := sval.Type().Field(i)
		t.Logf("字段名:%s,字段类型:%v,字段tag json:%v, form:%v", field.Name, field.Type, field.Tag.Get("json"), field.Tag.Get("form"))
	}
	// 获取方法数
	methodCount := pval.NumMethod()
	t.Logf("方法数:%d", methodCount)
	// 获取方法的类型
	helloMethodType := pval.MethodByName("Hello").Type()
	// 获取方法入参
	inCount := helloMethodType.NumIn()
	t.Logf("方法入参个数:%d", inCount)
	// 获取方法出参
	outCount := helloMethodType.NumOut()
	t.Logf("方法出参个数:%d", outCount)
	paramType := helloMethodType.Out(0)
	t.Logf("方法出参类型:%s", paramType.String())
	// 调用方法
	result := pval.MethodByName("Hello").Call([]reflect.Value{})
	t.Logf("调用结果:%v", result)

	// 创建简单值
	intType := reflect.TypeOf(0)
	pInt := reflect.New(intType)
	sInt := reflect.New(intType).Elem()
	sInt.SetInt(100)
	t.Logf("创建简单值指针:%v", pInt)
	t.Logf("创建简单值:%v", sInt)

	// 创建指针结构体
	pUser := reflect.New(upType)
	t.Logf("创建结构体指针:%v", pUser)
	t.Logf("创建结构体:%v", pUser.Elem())
	// 创建结构体
	pUser2 := reflect.New(upType.Elem())
	pUser2.Elem().FieldByName("Id").SetInt(100)
	pUser2.Elem().FieldByName("Name").SetString("ls")
	t.Logf("创建结构体指针:%+v", pUser2)
	t.Logf("创建结构体:%+v", pUser2.Elem())
}

func Test_2(t *testing.T) {
	var data interface{} = Data{
		"Name": "ls",
		"Age":  "18",
	}
	m := map[string]string(data.(Data))
	t.Logf("map:%+v", m)
	s, ok := data.(string)
	if ok {
		t.Logf("s:%s", s)
	}

}
