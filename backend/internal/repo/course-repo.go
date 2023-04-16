package repo

import (
	"backend/internal/entity"
	"time"

	"gorm.io/gorm"
)

type CourseRepo interface {
	CreateCourse(title, description string, maxListeners int, expiration, from, to time.Time, timetable []time.Time) (*entity.Course, error)
	UpdateCourse(id uint, title, description string, archive bool) (*entity.Course, error)
	FetchCourses(limit int) (*[]entity.Course, error)
	GetCourseById(id uint) (*entity.Course, error)
}

func NewCourseRepo(db *gorm.DB) CourseRepo {
	return &courseRepo{
		store: db,
	}
}

func (r *courseRepo) CreateCourse(title, description string, maxListeners int, expiration, from, to time.Time, timetable []time.Time) (*entity.Course, error) {
	course := entity.NewCourse(title, description, maxListeners, expiration, from, to, timetable)
	res := r.store.Create(course)

	if res.Error != nil {
		return nil, res.Error
	}

	return course, nil
}

func (r *courseRepo) UpdateCourse(id uint, title, description string, archive bool) (*entity.Course, error) {
	course := &entity.Course{
		ID: id,
	}

	res := r.store.First(course)
	if res.Error != nil {
		return nil, res.Error
	}

	course.Title = title
	course.Description = description
	course.Archive = archive

	res = r.store.Save(course)
	if res.Error != nil {
		return nil, res.Error
	}

	return course, nil
}

func (r *courseRepo) FetchCourses(limit int) (*[]entity.Course, error) {
	courses := &[]entity.Course{}

	var res *gorm.DB
	if limit > 0 {
		res = r.store.Limit(limit).Find(courses)
	} else {
		res = r.store.Find(courses)
	}

	if res.Error != nil {
		return nil, res.Error
	}

	return courses, nil
}

func (r *courseRepo) GetCourseById(id uint) (*entity.Course, error) {
	course := &entity.Course{
		ID: id,
	}

	res := r.store.First(course)
	if res.Error != nil {
		return nil, res.Error
	}

	return course, nil
}
