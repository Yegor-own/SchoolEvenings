package repo

import "gorm.io/gorm"

type userRepo struct {
	store *gorm.DB
}

type courseRepo struct {
	store *gorm.DB
}

type registryRepo struct {
	store *gorm.DB
}
