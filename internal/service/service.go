package service

import (
	"http-demo/internal/domain"
)

type Service struct {
	userStore UserStore
}

func NewService(userStore UserStore) Service {
	return Service{
		userStore: userStore,
	}
}

type UserStore interface {
	CreateUser(domain.User) (domain.User, error)
	GetUser(userId int) (domain.User, error)
	DeleteUser(userId int) (domain.User, error)
}

func (s Service) ChangePassword(user domain.User) error {
	// changing passowrd
	return nil
}

func (s Service) CreateUser(user domain.User) (domain.User, error) {
	newUser, err := s.userStore.CreateUser(user)
	return newUser, err
}

func (s Service) GetUser(userId int) (domain.User, error) {
	user, err := s.userStore.GetUser(userId)
	return user, err
}

func (s Service) DeleteUser(userId int) (domain.User, error) {
	user, err := s.userStore.DeleteUser(userId)
	return user, err
}
