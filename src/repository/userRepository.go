package repository

import (
  "github.com/yakhyadabo/go-rest-template/src/model"
  "sync"
)

type UserRepository interface {
	FindAll() ([]*model.User, error)
  FindByEmail(email string) (*model.User, error)
  Save(*model.User) error
}

type userRepository struct {
  mu    *sync.Mutex
  users map[string]*model.User
}

func NewUserRepository() *userRepository {
  return &userRepository{
      mu:    &sync.Mutex{},
      users: map[string]*model.User{},
  }
}

func (r *userRepository) Save(user *model.User)  (*model.User, error) {

  GetDB().Create(user)
  
  return user, nil
}

func (r *userRepository) FindAll() ([]*model.User, error) {
  r.mu.Lock()
  defer r.mu.Unlock()
  users := make([]*model.User, len(r.users))
  i := 0
  for _, user := range r.users {
      users[i] = model.NewUser(user.GetLogin(), user.GetPassword(), user.GetEmail())
      i++
  }
  return users, nil
}

func (r *userRepository) FindByLogin(login string) (*model.User, error) {

  user := &model.User{}
	err := GetDB().Table("users").Where("login = ?", login).First(user).Error
	if err != nil {
		return nil,nil
	}
  return user, nil

}