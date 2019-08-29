package service

import (
	"github.com/yakhyadabo/go-rest-template/src/repository"
  "github.com/yakhyadabo/go-rest-template/src/model"
  "fmt"
)

type UserService interface{
  ListUser() ([]*model.User, error)
  RegisterUser(email string) error
}

type userService struct {
  repo repository.UserRepository
}

// Dependency Injection ...
func NewUserService(repo repository.UserRepository) *userService {
  return &userService {
      repo:    repo,
  }
}

func (s *userService) Duplicated(email string) error {
  user, err := s.repo.FindByEmail(email)
  if user != nil {
      return fmt.Errorf("%s already exists", email)
  }
  if err != nil {
      return err
  }
  return nil
}

func (s *userService) ListUser() ([]*model.User, error) {
  users, err := s.repo.FindAll()
  if err != nil {
      return nil, err
  }
  return users, nil
}

/*
func (s *userService) RegisterUser(email string) error {
  uid, err := uuid.NewRandom()
  if err != nil {
      return err
  }

  if err := s.Duplicated(email); err != nil {
      return err
  }

  user := model.NewUser(uid.String(), email)
  if err := s.repo.Save(user); err != nil {
      return err
  }

  return nil
} */