package storage

import "backend/src/internal/entity"

type Access interface {
	CreateUser(user *entity.User) error
	FetchUser(id uint) (*entity.User, error)
	FetchUserParams(params map[string]any) (*entity.User, error)
	UpdateUser(id uint, params map[string]any) error
	DeleteUser(id uint) error

	CreateCourse(course *entity.Course) error
	FetchCourse(id uint) (*entity.Course, error)
	FetchCourseParams(params map[string]any) (*entity.Course, error)
	UpdateCourse(id uint, params map[string]any) error
	DeleteCourse(id uint) error

	//registry operations
	CreateEntry(registry *entity.Registry) error
	FetchEntry(id uint) error
	FetchEntryParams(params map[string]any) (*entity.Registry, error)
	UpdateEntry(id uint, params map[string]any) error

	//tag operations
	CreateTag(tag *entity.Tag) error
	// DeleteTag deletes relations between course and tag NOT TAG
	DeleteTag(id uint) error
}

type accessDriver struct {
	storage Access
}

func NewDataAccess(driver Access) Access {
	return &accessDriver{storage: driver}
}

func (d accessDriver) CreateUser(user *entity.User) error {
	return d.storage.CreateUser(user)
}

func (d accessDriver) FetchUser(id uint) (*entity.User, error) {
	return d.storage.FetchUser(id)
}

func (d accessDriver) FetchUserParams(params map[string]any) (*entity.User, error) {
	return d.storage.FetchUserParams(params)
}

func (d accessDriver) UpdateUser(id uint, params map[string]any) error {
	return d.storage.UpdateUser(id, params)
}

func (d accessDriver) DeleteUser(id uint) error {
	return d.storage.DeleteUser(id)
}

func (d accessDriver) CreateCourse(course *entity.Course) error {
	return d.storage.CreateCourse(course)
}

func (d accessDriver) FetchCourse(id uint) (*entity.Course, error) {
	return d.storage.FetchCourse(id)
}

func (d accessDriver) FetchCourseParams(params map[string]any) (*entity.Course, error) {
	return d.storage.FetchCourseParams(params)
}

func (d accessDriver) UpdateCourse(id uint, params map[string]any) error {
	return d.storage.UpdateCourse(id, params)
}

func (d accessDriver) DeleteCourse(id uint) error {
	return d.storage.DeleteCourse(id)
}

func (d accessDriver) CreateEntry(registry *entity.Registry) error {
	return d.storage.CreateEntry(registry)
}

func (d accessDriver) FetchEntry(id uint) error {
	return d.storage.FetchEntry(id)
}

func (d accessDriver) FetchEntryParams(params map[string]any) (*entity.Registry, error) {
	return d.storage.FetchEntryParams(params)
}

func (d accessDriver) UpdateEntry(id uint, params map[string]any) error {
	return d.storage.UpdateEntry(id, params)
}

func (d accessDriver) CreateTag(tag *entity.Tag) error {
	return d.storage.CreateTag(tag)
}

// DeleteTag deletes relations between course and tag NOT TAG
func (d accessDriver) DeleteTag(id uint) error {
	return d.storage.DeleteTag(id)
}
