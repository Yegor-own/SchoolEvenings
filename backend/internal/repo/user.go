package repo

import "backend/internal/entity"

type UserRepo interface {
	CreateUser(id uint, name, email string) (*entity.User, error)
}
