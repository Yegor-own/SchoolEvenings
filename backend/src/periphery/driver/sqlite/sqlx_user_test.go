package sqlite

import (
	"backend/src/intermediary/repository"
	"backend/src/intermediary/storage"
	"backend/src/internal/usecase"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestSqliteConnection(t *testing.T) {
	var db *sqlx.DB

	// exactly the same as the built-in
	db, err := sqlx.Open("sqlite3", "backend.db")

	if err != nil {
		t.Error(err)
	}
	// from a pre-existing sql.DB; note the required driverName
	//db = sqlx.NewDb(sql.Open("sqlite3", ":memory:"), "sqlite3")

	// force a connection and test that it worked
	err = db.Ping()
	if err != nil {
		t.Error(err)
	}
}

func TestSqliteSqlxDriver(t *testing.T) {
	var db *sqlx.DB

	// exactly the same as the built-in
	db, _ = sqlx.Open("sqlite3", "backend.db")
	driver, _ := NewSqlxDriver(db)
	access := storage.NewDataAccess(driver)
	repo := repository.NewUserRepository(access)
	usecase.NewUserService(repo)

}
