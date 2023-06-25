package sqlite

import (
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestNewSqlxDriver(t *testing.T) {
	//Connecting to sqlite
	var db *sqlx.DB
	db, err := sqlx.Open("sqlite3", "../../database/backend.db")

	if err != nil {
		t.Error(err)
	}

	//Declare new driver with scripts
	_, err = NewSqlxDriver(db, "scripts/courses.sql", "scripts/users.sql")
	if err != nil {
		t.Error(err)
	}
}
