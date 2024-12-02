package demo

import "log"

type User struct {
	Id   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

func (u *User) Hello() (name string) {
	log.Printf("hello %s", u.Name)
	return u.Name
}

type Data map[string]string
