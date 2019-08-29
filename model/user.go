package model

import(
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Login string `json:"login"`
	Password string `json:"password"`
	Email string `json:"email"`
}

func NewUser(login, password, email string) *User {
	return &User{
		Login: login,
		Password: password,
		Email: email,
	}
}


func (u *User) GetLogin() string {
	return u.Login
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetEmail() string {
	return u.Email
}