package entity

type Debt struct {
	ID          uint
	UserID      *uint
	Title       string
	Description string
}
