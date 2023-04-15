package entity

import "time"

type Registry struct {
	ID        uint
	UserID    *uint
	CourseID  *uint
	Reserve   bool
	CreatedAt time.Time
}

func NewRegistry(userId, courseId uint) *Registry {
	return &Registry{
		UserID:    &userId,
		CourseID:  &courseId,
		CreatedAt: time.Now(),
	}
}
