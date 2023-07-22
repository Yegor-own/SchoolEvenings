package repository

import (
	"backend/src/domain/entity"
	"backend/src/intermediary/storage"
	"encoding/json"
	"time"
)

type CourseRepository struct {
	Storage storage.Access
}

func NewCourseRepository(access storage.Access) *CourseRepository {
	return &CourseRepository{
		Storage: access,
	}
}

func (r CourseRepository) GetAll() ([]entity.Course, error) {
	return r.Storage.FetchAllCourses()
}

func (r CourseRepository) Get(id uint) (*entity.Course, error) {
	return r.Storage.FetchCourse(id)
}

func (r CourseRepository) Create(title, description string, ageRange [2]uint, previewUUID string, maxListeners uint, timetable []time.Weekday, from, to, expiresAt time.Time, creatorID uint) (*entity.Course, error) {
	course := entity.NewCourse(title, description, ageRange, previewUUID, maxListeners, timetable, from, to, expiresAt, creatorID)
	err := r.Storage.CreateCourse(course)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (r CourseRepository) ChangeData(course entity.Course) (*entity.Course, error) {
	var params map[string]any

	b, err := json.Marshal(course)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &params)
	if err != nil {
		return nil, err
	}
	err = r.Storage.UpdateCourse(course.ID, params)
	if err != nil {
		return nil, err
	}

	return r.Storage.FetchCourse(course.ID)
}
func (r CourseRepository) Delete(id uint) error {
	return r.Storage.DeleteCourse(id)
}
func (r CourseRepository) AddTag(courseID, tagID uint) error {
	return nil
}

func (r CourseRepository) RemoveTag(courseID, tagID uint) error {
	//TODO develop functions for tags
	return nil
}
func (r CourseRepository) AddDoc(id uint, doc *entity.Doc) error {
	//TODO develop functions for tags
	return nil
}
func (r CourseRepository) RemoveDoc(docID uint) error {
	//TODO develop functions for tags
	return nil
}
