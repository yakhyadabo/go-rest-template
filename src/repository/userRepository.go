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

func (r *userRepository) Save(user *model.User) error {
  r.mu.Lock()
  defer r.mu.Unlock()

  //r.users[user.GetID()] = model.NewUser(user.GetID(), user.GetEmail())

  
  r.users[user.GetLogin()] = &model.User{
      Id: user.GetID(),
      Login: user.GetLogin(),
      Password: user.GetPassword(),
      Email: user.GetEmail(),
  }
  return nil
}

func (r *userRepository) FindAll() ([]*model.User, error) {
  r.mu.Lock()
  defer r.mu.Unlock()
  users := make([]*model.User, len(r.users))
  i := 0
  for _, user := range r.users {
      users[i] = model.NewUser(user.GetID(), user.GetLogin(), user.GetPassword(), user.GetEmail())
      i++
  }
  return users, nil
}

func (r *userRepository) FindByLogin(login string) (*model.User, error) {
  r.mu.Lock()
  defer r.mu.Unlock()
  for _, user := range r.users {
      if user.GetLogin() == login {
          return model.NewUser(user.GetID(), user.GetLogin(),user.GetPassword(), user.GetEmail()), nil
      }
  }
  return nil, nil
}