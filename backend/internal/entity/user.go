package entity

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Admin bool   `json:"email"`
}

func NewUser(name, email string) *User {
	return &User{
		Name:  name,
		Email: email,
		Admin: false,
	}
}
