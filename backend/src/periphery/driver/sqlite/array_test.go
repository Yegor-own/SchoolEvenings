package sqlite

import (
	"backend/src/domain/entity"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
	"testing"
	"time"
)

func TestArrayInsert(t *testing.T) {
	var db *sqlx.DB

	// exactly the same as the built-in
	//TODO may I can use there regex
	db, _ = sqlx.Open("sqlite3", "test.db")
	stmt := `INSERT INTO courses 
    (title, description, age_range, preview_uuid, max_listeners, timetable, "from", "to", expires_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	c := entity.NewCourse("ttl", "dscr", [2]uint{11, 14}, uuid.New().String(), 50, []time.Weekday{1, 2, 3}, time.Now(), time.Now().Add(time.Hour*24*20), time.Now())
	res, err := db.Exec(stmt, c.Title, c.Description, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(c.AgeRange)), " "), "[]"), c.PreviewUUID, c.MaxListeners, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(c.Timetable)), " "), "[]"), c.From, c.To, c.ExpiresAt)
	if err != nil {
		t.Error(err)
	}

	log.Println(res.LastInsertId())
}
