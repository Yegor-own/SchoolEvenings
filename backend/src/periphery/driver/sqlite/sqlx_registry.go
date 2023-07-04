package sqlite

import (
	"backend/src/domain/entity"
	"errors"
	"time"
)

func (d SqlxDriver) CreateEntry(registry *entity.Registry) error {
	stmt := `INSERT INTO registry (user_id, course_id, reserve, created_at) VALUES (?, ?, ?, ?)`
	res, err := d.DB.Exec(
		stmt,
		registry.UserID,
		registry.CourseID,
		registry.Reserve,
		registry.CreatedAt.Format(time.RFC822))

	// Handle executing errors
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); n == 0 || err != nil {
		return errors.New("0 rows affected by inserting new entry")
	}

	id, err := res.LastInsertId()
	registry.ID = uint(id)

	return nil
}

func (d SqlxDriver) FetchEntry(id uint) (*entity.Registry, error) {
	row := d.DB.QueryRowx(`SELECT * FROM users WHERE id = ?`, int64(id))

	registry := entity.Registry{}
	err := row.StructScan(&registry)
	if err != nil {
		return nil, err
	}

	return &registry, nil
}

func (d SqlxDriver) FetchEntryParams(params map[string]any) (*entity.Registry, error) {
	return nil, nil
}

func (d SqlxDriver) UpdateEntry(id uint, params map[string]any) error {
	return nil
}
