package repository

import (
	"backend/src/intermediary/storage"
	"backend/src/internal/entity"
)

type UserRepository struct {
	Storage storage.Access
}

func NewUserRepository(access storage.Access) *UserRepository {
	return &UserRepository{Storage: access}
}

func (r UserRepository) Register(name, surname, patronymic, email string, phone int, password string) (user *entity.User, err error) {
	user = entity.NewUser(name, surname, patronymic, email, phone, password, false, false)
	_, err = r.Storage.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepository) Login(email, password string) (token string, err error) {
	user := entity.User{
		Email:    email,
		Password: password,
	}

	params := make(map[string]any)
	params["email"] = email
	params["password"] = password

	_, err = r.Storage.FetchParams(user, params)
	if err != nil {
		return "", err
	}

	//TODO develop a JWT token generator
	token = ""

	return token, nil
}

func (r UserRepository) ChangeData(data entity.User) (*entity.User, error) {

}

func (r UserRepository) Delete(id uint) error {
}

func (r UserRepository) Verify(id uint) error {
}

func (r UserRepository) Enroll(userID, courseID uint) error {
}

func (r UserRepository) Reserve(userID, courseID uint) error {
}

func (r UserRepository) CancelReservation(userID, courseID uint) error {
}

func (r UserRepository) Confirm(userID, courseID uint) error {
}
