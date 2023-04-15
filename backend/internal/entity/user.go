package entity

type User struct {
	ID         uint       `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Debts      []Debt     `json:"debts"`
	Attendance Attendance `json:"attendance"`
	Sections   []Section  `json:"sections"`
	SchoolID   uint       `json:"school_id"`
}

func NewUser(id uint, name, email string) *User {
	return &User{}
}
