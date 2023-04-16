package gateway

import "time"

type GetCourses struct {
	Limit int `json:"limit"`
}

type CreateCourse struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	MaxListeners int    `json:"max_listeners"`
	Expiration   time.Time
	From         time.Time
	To           time.Time
	Timetable    []time.Time
}

type GetCourseById struct {
	Id uint `json:"id"`
}
