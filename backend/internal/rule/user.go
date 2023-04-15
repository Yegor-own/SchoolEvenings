package rule

import (
	"backend/internal/entity"
	"backend/internal/repo"
)

type UserService interface {
	CreateUser(name, email string) (*entity.User, error)
	UpdateUser(id uint, name string) (*entity.User, error)
}

type userService struct {
	repo repo.UserRepo
}

func NewUserService(r repo.UserRepo) UserService {
	return &userService{
		repo: r,
	}
}

func (s *userService) CreateUser(name, email string) (*entity.User, error) {
	return s.repo.CreateUser(name, email)
}

func (s *userService) UpdateUser(id uint, name string) (*entity.User, error) {
	return s.repo.UpdateUser(id, name)
}
