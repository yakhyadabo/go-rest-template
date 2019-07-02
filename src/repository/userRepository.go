package repository

import (
	"github.com/yakhyadabo/go-rest-template/src/model"
)

type UserRepository interface {
	FindAll() ([]*model.User, error)
  FindByEmail(email string) (*model.User, error)
  Save(*model.User) error
}