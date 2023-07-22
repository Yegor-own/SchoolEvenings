package repository

import (
	"backend/src/domain/entity"
	"backend/src/intermediary/storage"
	"backend/src/periphery/hash"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type UserRepository struct {
	Storage storage.Access
}

func NewUserRepository(access storage.Access) *UserRepository {
	return &UserRepository{Storage: access}
}

func (r UserRepository) Get(id uint) (*entity.User, error) {
	return r.Storage.FetchUser(id)
}

func (r UserRepository) GetByEmail(email string) (*entity.User, error) {
	params := make(map[string]any)
	params["email"] = email
	return r.Storage.FetchUserParams(params)
}

func (r UserRepository) Register(name, surname, patronymic, email string, phone int, password string) (user *entity.User, err error) {
	// Hash password
	password, err = hash.HashPassword(password)
	if err != nil {
		log.Println(err)
	}

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

	user, err := r.Storage.FetchUserParams(params)
	if err != nil {
		return "", err
	}

	// Compare password
	if ok := hash.ComparePassword(password, user.Password); !ok {
		return "", nil
	}

	// Get env salt
	err = godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	salt := []byte(os.Getenv("JWT"))

	// Generate JWT token
	t := jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(5 * time.Hour * 24).Unix()

	token, err = t.SignedString(salt)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r UserRepository) ChangeData(data entity.ChangeUserData) (*entity.User, error) {
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
