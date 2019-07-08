package model

import(
// "github.com/jinzhu/gorm"
)

type User struct {
	// gorm.Model
	Id    string `json:"id"`
	Login string `json:"login"`
	Password string `json:"password"`
	Email string `json:"email"`
}

func NewUser(id string, login, password, email string) *User {
	return &User{
		Id:    id,
		Login: login,
		Password: password,
		Email: email,
	}
}

func (u *User) GetID() string {
	return u.Id
}

func (u *User) SetID(id string) {
	u.Id = id
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