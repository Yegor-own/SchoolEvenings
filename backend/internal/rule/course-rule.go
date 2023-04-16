package rule

import (
	"backend/internal/entity"
	"backend/internal/repo"
	"time"
)

type CourseService interface {
	CreateCourse(title, description string, maxListeners int, expiration, from, to time.Time, timetable []time.Time) (*entity.Course, error)
	UpdateCourse(id uint, title, description string, archive bool) (*entity.Course, error)
	FetchCourses(limit int) (*[]entity.Course, error)
	GetCourseById(id uint) (*entity.Course, error)
}

type courseService struct {
	repo repo.CourseRepo
}

func NewCourseService(r repo.CourseRepo) CourseService {
	return &courseService{
		repo: r,
	}
}

func (s *courseService) CreateCourse(title, description string, maxListeners int, expiration, from, to time.Time, timetable []time.Time) (*entity.Course, error) {
	return s.repo.CreateCourse(title, description, maxListeners, expiration, from, to, timetable)
}

func (s *courseService) UpdateCourse(id uint, title, description string, archive bool) (*entity.Course, error) {
	return s.repo.UpdateCourse(id, title, description, archive)
}

func (s *courseService) FetchCourses(limit int) (*[]entity.Course, error) {
	return s.repo.FetchCourses(limit)
}

func (s *courseService) GetCourseById(id uint) (*entity.Course, error) {
	return s.repo.GetCourseById(id)
}
