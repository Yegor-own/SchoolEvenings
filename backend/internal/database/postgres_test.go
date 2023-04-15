package database

import (
	"backend/internal/entity"
	"testing"
)

func TestConnection(t *testing.T) {
	pg, err := ConnectPostgres()
	if err != nil {
		t.Error(err)
	}

	pg.AutoMigrate(&entity.User{}, &entity.Course{}, &entity.Registry{})
}
