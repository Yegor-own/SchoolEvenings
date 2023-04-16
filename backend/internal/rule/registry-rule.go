package rule

import "backend/internal/entity"

type RegistryService interface {
	EnrollUser(userId, courseId uint) (*entity.Registry, error)
	ConfirmEnrollship(userId, courseId uint) (*entity.Registry, error)
}

type registryService struct {
	repo repo.RegistryRepo
}

func (s *registryService) EnrollUser(userId, courseId uint) (*entity.Registry, error) {
	return s.repo.EnrollUser(userId, courseId)
}

func (s *registryService) ConfirmEnrollship(userId, courseId uint) (*entity.Registry, error) {
	return s.repo.ConfirmEnrollship(userId, courseId)
}
