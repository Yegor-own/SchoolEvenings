package sqlite

import (
	"backend/src/domain/entity"
	"errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// CreateUser uses a pointer for user var and writes in it, if fails return err
func (d SqlxDriver) CreateUser(user *entity.User) error {
	//TODO make it more efficient ->
	stmt := `INSERT INTO users (name, surname, patronymic, email, phone, password) VALUES (?, ?, ?, ?, ?, ?)`
	res, err := d.DB.Exec(stmt, user.Name, user.Surname, user.Patronymic, user.Email, user.Phone, user.Password)

	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); n == 0 || err != nil {
		return errors.New("0 rows affected by inserting new user")
	}

	id, err := res.LastInsertId()
	user.ID = uint(id)

	return nil
}

func (d SqlxDriver) FetchUser(id uint) (*entity.User, error) {
	row := d.DB.QueryRowx(`SELECT * FROM users WHERE id = ?`, int64(id))

	user := entity.User{}
	err := row.StructScan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (d SqlxDriver) FetchUserParams(params map[string]any) (*entity.User, error) {
	var query = `SELECT * FROM users WHERE `

	first := false
	for param, _ := range params {
		if first {
			query += ` AND ` + param + ` = :` + param
			continue
		}
		query += param + ` = :` + param
		first = true
	}

	query, args, err := sqlx.Named(query, params)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	query = d.DB.Rebind(query)
	row := d.DB.QueryRowx(query, args...)

	user := entity.User{}
	err = row.StructScan(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (d SqlxDriver) UpdateUser(id uint, params map[string]any) error {
	var stmt = `UPDATE users SET `

	first := false
	for param, _ := range params {
		if first {
			stmt += `, ` + param + ` = :` + param
			continue
		}
		stmt += param + ` = :` + param
		first = true
	}
	stmt += ` WHERE id = :id`
	params["id"] = id

	stmt, args, err := sqlx.Named(stmt, params)
	if err != nil {
		return err
	}

	stmt, args, err = sqlx.In(stmt, args...)
	if err != nil {
		return err
	}

	stmt = d.DB.Rebind(stmt)
	res, err := d.DB.Exec(stmt, args...)
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); n == 0 || err != nil {
		return errors.New("0 rows affected by updating user")
	}
	return nil
}

func (d SqlxDriver) DeleteUser(id uint) error {
	stmt := `DELETE FROM users WHERE id = ?`
	res, err := d.DB.Exec(stmt, int64(id))
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); n == 0 || err != nil {
		return errors.New("0 rows affected by deleting user")
	}

	return nil
}
