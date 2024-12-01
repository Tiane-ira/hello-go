package model

type CsUser struct {
	ObjBase
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (c CsUser) TableName() string { return "cs_user" }
