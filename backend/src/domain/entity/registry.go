package entity

import "time"

type Registry struct {
	ID        uint      `json:"id" :"id"`
	UserID    uint      `json:"user_id" :"user_id"`
	CourseID  uint      `json:"course_id" :"course_id"`
	Reserve   bool      `json:"reserve" :"reserve"`
	Confirmed bool      `json:"confirmed" :"confirmed"`
	CreatedAt time.Time `json:"created_at" :"created_at"`
}

func NewRegistry(userID, courseID uint, reserve, confirmed bool) *Registry {
	return &Registry{
		UserID:    userID,
		CourseID:  courseID,
		Reserve:   reserve,
		Confirmed: confirmed,
		CreatedAt: time.Now(),
	}
}
