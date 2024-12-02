package model

import "encoding/json"

type CsUser struct {
	ObjBase
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (c CsUser) TableName() string { return "cs_user" }

func (c *CsUser) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

func (c *CsUser) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}
