package entity

type User struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Surname    string
	Patrynomic string
	Email      string
	Phone      int
	Password   string
	Admin      bool `json:"email"`
}

func NewUser(name, surname, email, password string) *User {
	return &User{
		Name:  name,
		Email: email,
		Admin: false,
	}
}
