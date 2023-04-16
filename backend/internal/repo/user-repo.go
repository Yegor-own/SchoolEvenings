package repo

import (
	"backend/internal/entity"
	"backend/internal/hutils"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(name, surname, email, password string) (*entity.User, error)
	UpdateUser(id uint, name, surname, patrynomic string, phone int) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		store: db,
	}
}

func (r *userRepo) CreateUser(name, surname, email, password string) (*entity.User, error) {
	password, err := hutils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := entity.NewUser(name, surname, email, password)

	res := r.store.Create(user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (r *userRepo) UpdateUser(id uint, name, surname, patrynomic string, phone int) (*entity.User, error) {
	user := &entity.User{ID: id}
	res := r.store.First(user)

	if res.Error != nil {
		return nil, res.Error
	}

	user.Name = name
	user.Surname = surname
	user.Patrynomic = patrynomic
	user.Phone = phone

	res = r.store.Save(user)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (r *userRepo) GetUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{Email: email}
	res := r.store.Find(user, "email = ?", email)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}
