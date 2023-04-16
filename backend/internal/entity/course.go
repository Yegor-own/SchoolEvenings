package entity

import "time"

type Course struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	MaxListeners int    `json:"max_listeners"`
	Archive      bool   `json:"archive"`
	// Confirmation bool   `json:"confirmation"`
	ExpiresOn time.Time
	From      time.Time
	To        time.Time
	Timetable []time.Time
	CreatedAt time.Time
}

func NewCourse(title, description string, maxListeners int, expiration, from, to time.Time, timetable []time.Time) *Course {
	return &Course{
		Title:        title,
		Description:  description,
		MaxListeners: maxListeners,
		ExpiresOn:    expiration,
		From:         from,
		To:           to,
		Timetable:    timetable,
		CreatedAt:    time.Now(),
	}
}
