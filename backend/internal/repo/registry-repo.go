package repo

import (
	"backend/internal/entity"

	"gorm.io/gorm"
)

type RegistryRepo interface {
	EnrollUser(userId, courseId uint) (*entity.Registry, error)
	ConfirmEnrollship(userId, courseId uint) (*entity.Registry, error)
}

func NewRegistryRepo(store *gorm.DB) RegistryRepo {
	return &registryRepo{
		store: store,
	}
}

func (r *registryRepo) EnrollUser(userId, courseId uint) (*entity.Registry, error) {
	registry := entity.NewRegistry(userId, courseId)
	course := &entity.Course{
		ID: courseId,
	}

	res := r.store.First(course)
	if res.Error != nil {
		return nil, res.Error
	}

	if course.MaxListeners

	res = r.store.Create(registry)
	if res.Error != nil {
		return nil, res.Error
	}

	return registry, nil
}

func (r *registryRepo) ConfirmEnrollship(userId, courseId uint) (*entity.Registry, error) {
	var registry *entity.Registry
	res := r.store.Find(registry, "user_id = ? AND course_id = ?", userId, courseId)
	if res.Error != nil {
		return nil, res.Error
	}

	registry.Confirmation = true

	return registry, nil
}
