package rule

import (
	"backend/internal/entity"
	"backend/internal/repo"
)

type UserService interface {
	CreateUser(id uint, name, email string) (*entity.User, error)
}

type userService struct {
	repo repo.UserRepo
}

func NewUserService(r repo.UserRepo) UserService {
	return &userService{
		repo: r,
	}
}

func (s *userService) CreateUser(id uint, name, email string) (*entity.User, error) {
	return &entity.User{}, nil
}
