package sqlite

import (
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestNewSqlxDriver(t *testing.T) {
	var db *sqlx.DB
	// exactly the same as the built-in
	db, err := sqlx.Open("sqlite3", "../../database/backend.db")

	if err != nil {
		t.Error(err)
	}

	_, err = NewSqlxDriver(db, "scripts/course.sql")
	if err != nil {
		t.Error(err)
	}
}
