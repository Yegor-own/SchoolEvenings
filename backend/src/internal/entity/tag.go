package entity

type Tag struct {
	ID    uint   `json:"id" :"id"`
	Title string `json:"title" :"title"`
}

func NewTag(title string) *Tag {
	return &Tag{
		Title: title,
	}
}

type ConnectedTags struct {
	ID       uint `json:"id"`
	TagID    uint `json:"tag_id"`
	CourseID uint `json:"course_id"`
}
