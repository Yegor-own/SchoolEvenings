package rule

import (
	"backend/internal/entity"
	"backend/internal/repo"
)

type UserService interface {
	CreateUser(name, surname, email, password string) (*entity.User, error)
	UpdateUser(id uint, name, surname, patrynomic string, phone int) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
}

type userService struct {
	repo repo.UserRepo
}

func NewUserService(r repo.UserRepo) UserService {
	return &userService{
		repo: r,
	}
}

func (s *userService) CreateUser(name, surname, email, password string) (*entity.User, error) {
	return s.repo.CreateUser(name, surname, email, password)
}

func (s *userService) UpdateUser(id uint, name, surname, patrynomic string, phone int) (*entity.User, error) {
	return s.repo.UpdateUser(id, name, surname, patrynomic, phone)
}

func (s *userService) GetUserByEmail(email string) (*entity.User, error) {
	return s.repo.GetUserByEmail(email)
}
