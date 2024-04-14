package repository

import (
	"fmt"
	"weekly-newsletter/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(model.User) (model.User, error)
	Delete(model.User) error
	GetUsers() ([]model.User, error)
}

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(user model.User) (model.User, error) {
	fmt.Println("create user", user)
	if err := r.db.Create(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *Repository) Delete(user model.User) error {
	tx := r.db.Begin()
	tx.Unscoped().Delete(&user, "email = ?", user.Email)
	if err := tx.Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (r *Repository) GetUsers() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
