package obj

import "hello-go/global/model"

type CsUser struct {
	model.ObjBase
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (c CsUser) TableName() string { return "cs_user" }
