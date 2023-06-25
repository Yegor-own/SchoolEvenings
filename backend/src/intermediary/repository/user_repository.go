package repository

import (
	"backend/src/intermediary/storage"
	"backend/src/internal/entity"
	"encoding/json"
)

type UserRepository struct {
	Storage storage.Access
}

func NewUserRepository(access storage.Access) *UserRepository {
	return &UserRepository{Storage: access}
}

func (r UserRepository) Register(name, surname, patronymic, email string, phone int, password string) (user *entity.User, err error) {
	user = entity.NewUser(name, surname, patronymic, email, phone, password, false, false)
	err = r.Storage.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepository) Login(email, password string) (token string, err error) {
	params := make(map[string]any)
	params["email"] = email
	//TODO encrypt password
	params["password"] = password

	user, err := r.Storage.FetchUserParams(params)
	if err != nil {
		return "", err
	}

	//TODO develop a JWT token generator
	token = "aaa_token" + user.Email

	return token, nil
}

func (r UserRepository) ChangeData(data entity.User) (*entity.User, error) {
	var mapData map[string]any
	m, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(m, &mapData)
	if err != nil {
		return nil, err
	}

	err = r.Storage.UpdateUser(data.ID, mapData)
	if err != nil {
		return nil, err
	}

	return r.Storage.FetchUser(data.ID)
}

func (r UserRepository) Delete(id uint) error {
	return r.Storage.DeleteUser(id)
}

func (r UserRepository) Verify(id uint) error {
	return r.Storage.UpdateUser(id, map[string]any{"verified": true})
}

func (r UserRepository) Enroll(userID, courseID uint) error {
	//TODO develop functions for enrollment
	return nil
}

func (r UserRepository) Reserve(userID, courseID uint) error {
	//TODO develop functions for
	return nil
}

func (r UserRepository) CancelReservation(userID, courseID uint) error {
	//TODO develop functions for reservation
	return nil
}

func (r UserRepository) Confirm(userID, courseID uint) error {
	//TODO develop functions for confirmation
	return nil
}
