package user

import (
	"errors"
)

type Service interface {
	Register(input InputUser) (User, error)
	Login(input InputUser) (string, error)
	GetById(id int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) Register(input InputUser) (User, error) {
	var user User
	user.Email = input.Email
	user.Password = hashPassword(input.Password)

	user, err := s.repository.Register(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) Login(input InputUser) (string, error) {

	token, err := s.repository.Login(input.Email, input.Password)
	if err != nil {
		return "", errors.New("Credentials invalid")
	}

	return token, nil
}

func (s *service) GetById(id int) (User, error) {
	user, err := s.repository.GetById(id)
	if err != nil {
		return user, err
	}
	user.Password = ""
	return user, nil
}
