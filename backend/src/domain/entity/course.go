package entity

import "time"

type Course struct {
	ID           uint           `json:"id"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	AgeRange     [2]uint        `json:"age_range"`    // AgeRange parameter is about ages [minimum age, maximum age] to involve users
	PreviewUUID  string         `json:"preview_uuid"` // PreviewUUID parameter contains UUID of the preview image
	MaxListeners uint           `json:"max_listeners"`
	Timetable    []time.Weekday `json:"timetable"`
	From         time.Time      `json:"from"`
	To           time.Time      `json:"to"`
	ExpiresAt    time.Time      `json:"expires_at"`
	CreatorID    uint           `json:"creator_id"`
}

func NewCourse(title, description string, ageRange [2]uint, previewUUID string, maxListeners uint, timetable []time.Weekday, from, to, expiresAt time.Time, creatorID uint) *Course {
	return &Course{
		Title:        title,
		Description:  description,
		AgeRange:     ageRange,
		PreviewUUID:  previewUUID,
		MaxListeners: maxListeners,
		Timetable:    timetable,
		From:         from,
		To:           to,
		ExpiresAt:    expiresAt,
		CreatorID:    creatorID,
	}
}

type Doc struct {
	ID       uint   `json:"id" :"id"`
	CourseID uint   `json:"course_id" :"course_id"`
	UUID     string `json:"uuid" :"uuid"`
}

func NewDoc(courseID uint, uuid string) *Doc {
	return &Doc{
		CourseID: courseID,
		UUID:     uuid,
	}
}
