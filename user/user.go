package user

import (
	"../settings"
	"fmt"
)

type User struct {
	login     string
	email     string
	firstname string
	lastname  string
	gender    string
}

var user = User{}

func GetUser() *User {
	return &user
}

func (u *User) SetLogin(login string) {
	u.login = login
}

func (u *User) GetLogin() string {
	return u.login
}
