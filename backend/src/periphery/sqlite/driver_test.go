package sqlite

import (
	"backend/src/intermediary/repository"
	"backend/src/intermediary/storage"
	"backend/src/internal/usecase"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestSqliteDriver(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Error("failed to connect database")
	}

	sl := NewSqliteDriver(db)
	ac := storage.NewDataAccess(sl)
	rp := repository.NewUserRepository(ac)
	usecase.NewUserService(rp)
}
