package gateway

type GetCourses struct {
	Limit int `json:"limit"`
}

type CreateCourse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetCourseById struct {
	Id uint `json:"id"`
}
