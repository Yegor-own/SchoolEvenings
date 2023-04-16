package entity

import "time"

type Registry struct {
	ID           uint
	UserID       *uint `json:"user_id"`
	CourseID     *uint `json:"course_id"`
	Reserve      bool
	Confirmation bool
	CreatedAt    time.Time
}

func NewRegistry(userId, courseId uint) *Registry {
	return &Registry{
		UserID:    &userId,
		CourseID:  &courseId,
		CreatedAt: time.Now(),
	}
}
