package sqlite

import (
	"backend/src/domain/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
	"testing"
	"time"
)

func TestSqlxDriver_CreateCourse(t *testing.T) {
	//Connecting db
	var db *sqlx.DB
	db, err := sqlx.Open("sqlite3", "../../database/backend.db")
	if err != nil {
		t.Error(err)
	}

	//Declare new driver
	driver, err := NewSqlxDriver(db)
	if err != nil {
		t.Error(err)
	}

	//Testing CreateCourse function
	course := entity.NewCourse("ttl", "dscr", [2]uint{11, 14}, uuid.New().String(), 50, []time.Weekday{1, 2, 3}, time.Now(), time.Now().Add(time.Hour*24*20), time.Now())
	err = driver.CreateCourse(course)
	if err != nil {
		t.Error(err)
	}

	log.Println(course.ID)
}

func TestSqlxDriver_FetchCourse(t *testing.T) {
	var db *sqlx.DB
	db, err := sqlx.Open("sqlite3", "../../database/backend.db")
	if err != nil {
		t.Error(err)
	}

	//Declare new driver
	driver, err := NewSqlxDriver(db)
	if err != nil {
		t.Error(err)
	}

	// Testing FetchCourse function
	course, err := driver.FetchCourse(2)
	if err != nil {
		t.Error(err)
	}
	log.Println(course)
}
