package repository

import "myapp/internal/model"

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetByID(id uint) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id uint) error
}
