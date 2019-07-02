package service

import (
	"github.com/yakhyadabo/go-rest-template/src/repository"
)

type UserService struct {
  repo repository.UserRepository
}

func (s *UserService) Duplicated(email string) error {
  user, err := s.repo.FindByEmail(email)
  if user != nil {
      return fmt.Errorf("%s already exists", email)
  }
  if err != nil {
      return err
  }
  return nil
}