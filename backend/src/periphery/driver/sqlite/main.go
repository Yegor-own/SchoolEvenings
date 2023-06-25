package sqlite

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"os"
)

type SqlxDriver struct {
	DB *sqlx.DB
}

// NewSqlxDriver is a function to create new sqlx driver, with scripts var you can execute some sql scripts they should be a path string like path/to/script.sql
func NewSqlxDriver(db *sqlx.DB, scripts ...string) (*SqlxDriver, error) {
	for _, script := range scripts {
		file, err := os.ReadFile(script)
		if err != nil {
			return nil, errors.Join(errors.New("Failed to open script: "+script), err)
		}
		sql := string(file)
		_, err = db.Exec(sql)
		if err != nil {
			return nil, errors.Join(errors.New("Failed to execute script: "+script), err)
		}
	}
	return &SqlxDriver{DB: db}, nil
}
