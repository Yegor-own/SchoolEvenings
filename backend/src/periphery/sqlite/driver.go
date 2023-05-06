package sqlite

import (
	"errors"
	"gorm.io/gorm"
)

type SqliteDriver struct {
	DB *gorm.DB
}

func NewSqliteDriver(db *gorm.DB) *SqliteDriver {
	return &SqliteDriver{DB: db}
}

func (d SqliteDriver) Create(dest interface{}) (interface{}, error) {
	res := d.DB.Create(dest)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("No rows affected")
	}

	return dest, nil
}

func (d SqliteDriver) Fetch(dest interface{}) (interface{}, error) {
	res := d.DB.Find(dest)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("No rows affected")
	}

	return dest, nil
}

func (d SqliteDriver) Update(dest interface{}) (interface{}, error) {
	res := d.DB.Save(dest)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("No rows affected")
	}

	return dest, nil
}

func (d SqliteDriver) Delete(dest interface{}) error {
	res := d.DB.Delete(dest)

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("No rows affected")
	}

	return nil
}
