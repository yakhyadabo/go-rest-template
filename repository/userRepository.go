package repository

import (
	"github.com/yakhyadabo/go-rest-template/model"
)

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Save(*model.User) error
}

type userRepository struct {
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (r *userRepository) Save(user *model.User) (*model.User, error) {

	GetDB().Create(user)

	return user, nil
}

func (r *userRepository) FindByLogin(login string) (*model.User, error) {

	user := &model.User{}
	err := GetDB().Table("users").Where("login = ?", login).First(user).Error
	if err != nil {
		return nil, nil
	}
	return user, nil

}
