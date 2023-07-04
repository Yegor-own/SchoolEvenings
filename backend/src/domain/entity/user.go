package entity

type User struct {
	ID         uint   `json:"id" :"id"`
	Name       string `json:"name" :"name"`
	Surname    string `json:"surname" :"surname"`
	Patronymic string `json:"patronymic" :"patronymic"`
	Email      string `json:"email" :"email"`
	Phone      int    `json:"phone" :"phone"`
	Password   string `json:"password" :"password"`
	Verified   bool   `json:"verified" :"verified"`
	Admin      bool   `json:"admin" :"admin"`
}

// NewUser returns a linker from new user model with given params
func NewUser(name, surname, patronymic, email string, phone int, password string, verified, admin bool) *User {
	return &User{
		Name:       name,
		Surname:    surname,
		Patronymic: patronymic,
		Email:      email,
		Phone:      phone,
		Password:   password,
		Verified:   verified,
		Admin:      admin,
	}
}
