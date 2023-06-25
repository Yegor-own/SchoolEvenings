package sqlite

import "backend/src/internal/entity"

func (d SqlxDriver) CreateCourse(course *entity.Course) error {
	//stmt := `INSERT INTO users ()`
	return nil
}
func (d SqlxDriver) FetchCourse(id uint) (*entity.Course, error) {
	return nil, nil
}
func (d SqlxDriver) FetchCourseParams(params map[string]any) (*entity.Course, error) {
	return nil, nil
}
func (d SqlxDriver) UpdateCourse(id uint, params map[string]any) error {
	return nil
}
func (d SqlxDriver) DeleteCourse(id uint) error {
	return nil
}
