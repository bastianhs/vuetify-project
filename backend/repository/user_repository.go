package repository

import (
	"errors"

	"github.com/bastianhs/vuetify-project/backend/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Create a new user
func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	result := r.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetAll retrieves all users
func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	result := r.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// Update updates a user
func (r *UserRepository) Update(user *models.User) error {
	if r.DB.Model(user).Updates(user).RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

// Delete deletes a user
func (r *UserRepository) Delete(id uint) error {
	result := r.DB.Delete(&models.User{}, id)
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return result.Error
}
