package repository

import (
	"librarymvc/internal/users/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) models.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) CreateUser(user *models.User) error {
	return u.db.Create(user).Error
}

func (u *UserRepository) GetUser(id int64) (*models.User, error) {
	var user models.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepository) UpdateUser(id int64, user *models.User) error {
	return u.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
}

func (u *UserRepository) DeleteUser(id int64) error {
	return u.db.Delete(&models.User{}, id).Error
}
