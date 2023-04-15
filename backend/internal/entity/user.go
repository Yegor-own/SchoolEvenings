package entity

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Debts []Debt `json:"debts"` //Долги по учебе
	// Attendance Attendance `json:"attendance"` // Посещаемость
	// Sections []Section `json:"sections"` // Посещаемые секции
	SchoolID *uint `json:"school_id"`
}

func NewUser(id uint, name, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
