package model

import(
// "github.com/jinzhu/gorm"
)

type User struct {
	// gorm.Model
	Id    string `json:"id"`
	Email string `json:"email"`
}

func NewUser(id, email string) *User {
	return &User{
		Id:    id,
		Email: email,
	}
}

func (u *User) GetID() string {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}