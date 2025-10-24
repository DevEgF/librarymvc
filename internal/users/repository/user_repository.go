package repository

import (
	"errors"
	"librarymvc/internal/users/models"
	"sync"
)

type UserRepository struct {
	user   map[int64]*models.User
	mu     sync.RWMutex
	nextId int64
}

func NewUserRepository() models.UserRepository {
	return &UserRepository{
		user:   make(map[int64]*models.User),
		nextId: 1,
	}
}

func (u *UserRepository) CreateUser(user *models.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	user.ID = u.nextId
	u.nextId++
	u.user[user.ID] = user
	return nil
}

func (u *UserRepository) GetUser(id int64) (*models.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	user, exists := u.user[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (u *UserRepository) GetAllUsers() ([]*models.User, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	user := make([]*models.User, 0, len(u.user))
	for _, v := range u.user {
		user = append(user, v)
	}

	return user, nil
}

func (u *UserRepository) UpdateUser(id int64, user *models.User) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	user, exists := u.user[id]
	if !exists {
		return errors.New("user not found")
	}

	u.user[user.ID] = user
	return nil
}

func (u *UserRepository) DeleteUser(id int64) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	_, exists := u.user[id]
	if !exists {
		return errors.New("user not found")
	}

	delete(u.user, id)
	return nil
}
