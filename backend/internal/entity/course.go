package entity

import "time"

type Course struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Archive     bool
	CreatedAt   time.Time
}

func NewCourse(title, description string) *Course {
	return &Course{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
	}
}
