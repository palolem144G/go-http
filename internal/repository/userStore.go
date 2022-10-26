package repository

import (
	"errors"
	"http-demo/internal/domain"
)

type InMemoryUserStore struct {
	userStore map[int]domain.User
}

func NewInMemoryUserStore() InMemoryUserStore {
	return InMemoryUserStore{
		userStore: make(map[int]domain.User),
	}
}

func (us InMemoryUserStore) CreateUser(user domain.User) (domain.User, error) {
	newId := len(us.userStore) + 1
	user.Id = newId
	us.userStore[newId] = user
	return user, nil
}

func (us InMemoryUserStore) GetUser(userId int) (domain.User, error) {
	user, ok := us.userStore[userId]
	if !ok {
		return domain.User{}, errors.New("user not found")
	}

	return user, nil
}

func (us InMemoryUserStore) DeleteUser(userId int) (domain.User, error) {
	user, ok := us.userStore[userId]
	delete(us.userStore, userId)
	if !ok {
		return domain.User{}, errors.New("user doest delete")
	}
	return user, nil
}

// добавить GetAllUsers
