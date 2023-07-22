package sqlite

import (
	"backend/src/domain/entity"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
	"time"
)

func convertTimetable(tt []time.Weekday) string {
	var timetable string
	for i, weekday := range tt {
		timetable += strconv.Itoa(int(weekday))
		if len(tt)-1 != i {
			timetable += " "
		}
	}

	return timetable
}

func parseTimetable(tt string) []time.Weekday {
	var timetable []time.Weekday
	for _, s := range strings.Split(tt, " ") {
		n, _ := strconv.Atoi(s)
		timetable = append(timetable, time.Weekday(n))
	}

	return timetable
}

func convertAgeRange(ar [2]uint) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ar)), " "), "[]")
}

func parseAgeRange(ar string) [2]uint {
	var ageRange [2]uint
	for i, s := range strings.Split(ar, " ") {
		age, _ := strconv.Atoi(s)
		ageRange[i] = uint(age)
	}

	return ageRange
}

func convertResponseToCourse(res *sqlx.Row) (*entity.Course, error) {
	var (
		course              entity.Course
		ageRange, timetable string
	)

	err := res.Scan(
		&course.ID,
		&course.Title,
		&course.Description,
		&ageRange,
		&course.PreviewUUID,
		&course.MaxListeners,
		&timetable,
		&course.To,
		&course.From,
		&course.ExpiresAt)
	if err != nil {
		return nil, err
	}

	// Parsing AgeRange
	course.AgeRange = parseAgeRange(ageRange)

	// Parsing Timetable
	course.Timetable = parseTimetable(timetable)

	return &course, nil
}

func (d SqlxDriver) CreateCourse(course *entity.Course) error {
	stmt := `INSERT INTO courses (title, description, age_range, preview_uuid, max_listeners, timetable, "from", "to", expires_at, creator_id) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// Convert timetable to string
	timetable := convertTimetable(course.Timetable)

	res, err := d.DB.Exec(
		stmt,
		course.Title,
		course.Description,
		convertAgeRange(course.AgeRange),
		course.PreviewUUID,
		course.MaxListeners,
		timetable,
		course.From.Format(time.RFC822),
		course.To.Format(time.RFC822),
		course.ExpiresAt.Format(time.RFC822))

	// Handle executing errors
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); n == 0 || err != nil {
		return errors.New("0 rows affected by inserting new course")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	course.ID = uint(id)

	return nil
}

func (d SqlxDriver) FetchCourse(id uint) (*entity.Course, error) {
	// Query to db
	res := d.DB.QueryRowx(`SELECT * FROM courses WHERE id = ?`, id)

	// Handle query result to entity
	return convertResponseToCourse(res)
}

func (d SqlxDriver) FetchCourseParams(params map[string]any) (*entity.Course, error) {
	var query = `SELECT * FROM courses WHERE `

	first := false
	for param, arg := range params {
		// Convert argument to  string, if it is time
		if t, ok := arg.(time.Time); ok {
			params[param] = t.Format(time.RFC822)
		}

		// Convert argument to  string, if it is list of weekdays
		if w, ok := arg.([]time.Weekday); ok {
			// Convert timetable to string
			params[param] = convertTimetable(w)
		}

		// Convert argument to  string, if it is list of ages
		if a, ok := arg.([2]uint); ok {
			// Convert agerange to string
			params[param] = convertAgeRange(a)
		}

		if first {
			query += ` AND ` + param + ` = :` + param
			continue
		}
		query += param + ` = :` + param
		first = true
	}

	// Setting named query
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

	// Handle query result to entity
	return convertResponseToCourse(row)
}

func (d SqlxDriver) FetchAllCourses() ([]entity.Course, error) {
	rows, err := d.DB.Queryx(`SELECT * FROM courses`)
	if err != nil {
		return nil, err
	}

	var courses []entity.Course
	for rows.Next() {
		var c entity.Course
		err = rows.StructScan(&c)
		if err != nil {
			return nil, err
		}

		courses = append(courses, c)
	}
	// check the error from rows
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (d SqlxDriver) UpdateCourse(id uint, params map[string]any) error {
	var stmt = `UPDATE courses SET `

	first := false
	for param, arg := range params {
		// Convert argument to  string, if it is time
		if t, ok := arg.(time.Time); ok {
			params[param] = t.Format(time.RFC822)
		}

		// Convert argument to  string, if it is list of weekdays
		if w, ok := arg.([]time.Weekday); ok {
			// Convert timetable to string
			params[param] = convertTimetable(w)
		}

		// Convert argument to  string, if it is list of ages
		if a, ok := arg.([2]uint); ok {
			// Convert agerange to string
			params[param] = convertAgeRange(a)
		}

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

	// Handle execution errors
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); n == 0 || err != nil {
		return errors.New("0 rows affected by updating user")
	}

	return nil
}

func (d SqlxDriver) DeleteCourse(id uint) error {
	stmt := `DELETE FROM courses WHERE id = ?`
	res, err := d.DB.Exec(stmt, int64(id))

	// Handle execution errors
	if err != nil {
		return err
	}
	if n, err := res.RowsAffected(); n == 0 || err != nil {
		return errors.New("0 rows affected by deleting user")
	}

	return nil
}
