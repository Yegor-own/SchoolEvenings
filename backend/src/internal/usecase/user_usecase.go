package usecase

import (
	"backend/src/internal/entity"
)

type UserService interface {
	Register(name, surname, patronymic, email string, phone int, password string) (user entity.User, err error)
	Login(email, password string) (token string, err error)
	ChangeData(data entity.User) (*entity.User, error)
	Delete(id uint) error
	Verify(id uint) error
	Enroll(userID, courseID uint) error
	Reserve(userID, courseID uint) error
	CancelReservation(userID, courseID uint) error
	Confirm(userID, courseID uint) error
}

type userService struct {
	repo UserService
}

func NewUserService(repo UserService) UserService {
	return &userService{
		repo: repo,
	}
}

func (s userService) Register(name, surname, patronymic, email string, phone int, password string) (user entity.User, err error) {
	return s.repo.Register(name, surname, patronymic, email, phone, password)
}

func (s userService) Login(email, password string) (token string, err error) {
	return s.repo.Login(email, password)
}

func (s userService) ChangeData(data entity.User) (usr *entity.User, err error) {
	return s.repo.ChangeData(data)
}

func (s userService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s userService) Verify(id uint) error {
	return s.repo.Verify(id)
}

func (s userService) Enroll(userID, courseID uint) error {
	return s.repo.Enroll(userID, courseID)
}

func (s userService) Reserve(userID, courseID uint) error {
	return s.repo.Reserve(userID, courseID)
}

func (s userService) CancelReservation(userID, courseID uint) error {
	return s.repo.CancelReservation(userID, courseID)
}

func (s userService) Confirm(userID, courseID uint) error {
	return s.repo.Confirm(userID, courseID)
}
