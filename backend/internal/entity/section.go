package entity

type Section struct {
	ID          uint
	SchoolID    *uint
	Title       string
	Description string
}

type UserSections struct {
	SectionID *uint `json:"section_id"`
	UserID    *uint `json:"user_id"`
}
