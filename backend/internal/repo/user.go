package repo

import (
	"backend/internal/entity"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(name, email string) (*entity.User, error)
	UpdateUser(id uint, name string) (*entity.User, error)
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		store: db,
	}
}

func (r *userRepo) CreateUser(name, email string) (*entity.User, error) {
	user := entity.NewUser(name, email)

	res := r.store.Create(user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (r *userRepo) UpdateUser(id uint, name string) (*entity.User, error) {
	user := &entity.User{ID: id}
	res := r.store.First(user)

	if res.Error != nil {
		return nil, res.Error
	}

	user.Name = name
	res = r.store.Save(user)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}
