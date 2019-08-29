package repository

import (
	"github.com/yakhyadabo/go-rest-template/model"
)

type EmployeeRepository interface {
	FindAll() ([]*model.Employee, error)
	FindByEmail(email string) (*model.Employee, error)
	Save(*model.Employee) error
}

type employeeRepository struct {
}

func NewEmployeeRepository() *employeeRepository {
	return &employeeRepository{}
}

func (r *employeeRepository) Save(employee *model.Employee) (*model.Employee, error) {

	GetDB().Create(employee)
	GetDB().Save(employee)

	return employee, nil
}

func (r *employeeRepository) FindByLogin(login string) (*model.Employee, error) {

	employee := &model.Employee{}
	err := GetDB().Table("employees").Where("login = ?", login).First(employee).Error
	if err != nil {
		return nil, nil
	}
	return employee, nil

}
