package domain

import (
	"hello-go/core/app"
	"time"
)

type ListUserReq struct {
	Start *app.DateTime `json:"start"`
	End   *app.DateTime `json:"end"`
}

type ListUserPageReq struct {
	Page
	Start time.Time `form:"start" time_format:"2006-01-02 15:04:05"`
	End   time.Time `form:"end" time_format:"2006-01-02 15:04:05"`
	Sort  string    `form:"sort"`
}

type UserReq struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
