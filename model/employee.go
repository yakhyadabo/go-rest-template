package model

import "github.com/jinzhu/gorm"

type Employee struct {
	gorm.Model
	//ID        string   `json:"id"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Login     string  `json:"login"`
	Email     string  `json:"email"`
	Company   Company `json:"company"`
	CompanyID uint    //`gorm:"foreignkey:CompanyId"`
}

func NewEmployee(firstname, lastname, login, email string) *Employee {
	return &Employee{
		Firstname: firstname,
		Lastname:  lastname,
		Login:     login,
		Email:     email,
	}
}

func (e *Employee) GetFirtname() string {
	return e.Firstname
}

func (e *Employee) GetLastname() string {
	return e.Lastname
}

func (e *Employee) GetLogin() string {
	return e.Login
}

func (e *Employee) GetEmail() string {
	return e.Email
}
