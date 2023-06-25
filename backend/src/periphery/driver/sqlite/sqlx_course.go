package sqlite

import (
	"backend/src/internal/entity"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (d SqlxDriver) CreateCourse(course *entity.Course) error {
	stmt := `INSERT INTO courses (title, description, age_range, preview_uuid, max_listeners, timetable, "from", "to", expires_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	var timetable string
	for i, weekday := range course.Timetable {
		timetable += strconv.Itoa(int(weekday))
		if len(course.Timetable)-1 != i {
			timetable += " "
		}
	}

	res, err := d.DB.Exec(stmt, course.Title, course.Description, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(course.AgeRange)), " "), "[]"), course.PreviewUUID, course.MaxListeners, timetable, course.From.Format(time.RFC822), course.To.Format(time.RFC822), course.ExpiresAt.Format(time.RFC822))
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
	var course entity.Course

	res := d.DB.QueryRowx(`SELECT * FROM courses WHERE id = ?`, id)

	var ageRange, timetable string
	err := res.Scan(&course.ID, &course.Title, &course.Description, &ageRange, &course.PreviewUUID, &course.MaxListeners, &timetable, &course.To, &course.From, &course.ExpiresAt)
	if err != nil {
		return nil, err
	}

	// Parsing AgeRange
	for i, s := range strings.Split(ageRange, " ") {
		age, _ := strconv.Atoi(s)
		course.AgeRange[i] = uint(age)
	}

	// Parsing Timetable
	for _, s := range strings.Split(timetable, " ") {
		n, _ := strconv.Atoi(s)
		course.Timetable = append(course.Timetable, time.Weekday(n))
	}

	return &course, nil
}

func (d SqlxDriver) FetchCourseParams(params map[string]any) (*entity.Course, error) {
	return nil, nil
}
func (d SqlxDriver) UpdateCourse(id uint, params map[string]any) error {
	return nil
}
func (d SqlxDriver) DeleteCourse(id uint) error {
	return nil
}
