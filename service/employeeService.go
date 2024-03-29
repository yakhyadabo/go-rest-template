package service

import (
	"fmt"

	"github.com/yakhyadabo/go-rest-template/model"
	"github.com/yakhyadabo/go-rest-template/repository"
)

type EmployeeService interface {
	ListEmployee() ([]*model.Employee, error)
	RegisterEmployee(email string) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

// Dependency Injection ...
func NewEmployeeService(repo repository.EmployeeRepository) *employeeService {
	return &employeeService{
		repo: repo,
	}
}

func (s *employeeService) Duplicated(email string) error {
	user, err := s.repo.FindByEmail(email)
	if user != nil {
		return fmt.Errorf("%s already exists", email)
	}
	if err != nil {
		return err
	}
	return nil
}

func (s *employeeService) ListEmployee() ([]*model.Employee, error) {
	employees, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return employees, nil
}
