package usecase

import (
	"backend/src/domain/entity"
	"time"
)

type CourseService interface {
	GetAll() ([]entity.Course, error)
	Get(id uint) (*entity.Course, error)
	Create(title, description string, ageRange [2]uint, previewUUID string, maxListeners uint, timetable []time.Weekday, from, to, expiresAt time.Time, creatorID uint) (*entity.Course, error)
	ChangeData(course entity.Course) (*entity.Course, error)
	Delete(id uint) error
	AddTag(courseID, tagID uint) error
	RemoveTag(courseID, tagID uint) error
	AddDoc(id uint, doc *entity.Doc) error
	RemoveDoc(docID uint) error
}

type courseService struct {
	repo CourseService
}

func NewCourseService(repo CourseService) CourseService {
	return &courseService{repo: repo}
}

func (s courseService) GetAll() ([]entity.Course, error) {
	return s.repo.GetAll()
}

func (s courseService) Get(id uint) (*entity.Course, error) {
	return s.repo.Get(id)
}

func (s courseService) Create(title, description string, ageRange [2]uint, previewUUID string, maxListeners uint, timetable []time.Weekday, from, to, expiresAt time.Time, creatorID uint) (*entity.Course, error) {
	return s.repo.Create(title, description, ageRange, previewUUID, maxListeners, timetable, from, to, expiresAt, creatorID)
}

func (s courseService) ChangeData(course entity.Course) (*entity.Course, error) {
	return s.repo.ChangeData(course)
}

func (s courseService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s courseService) AddTag(courseID, tagID uint) error {
	return s.repo.AddTag(courseID, tagID)
}

func (s courseService) RemoveTag(courseID, tagID uint) error {
	return s.repo.RemoveTag(courseID, tagID)
}
func (s courseService) AddDoc(id uint, doc *entity.Doc) error {
	return s.repo.AddDoc(id, doc)
}
func (s courseService) RemoveDoc(docID uint) error {
	return s.repo.RemoveDoc(docID)
}
