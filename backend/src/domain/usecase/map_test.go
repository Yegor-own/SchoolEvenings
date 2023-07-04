package usecase

import (
	"backend/src/domain/entity"
	"encoding/json"
	"log"
	"reflect"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	course := entity.Course{
		ID:           0,
		Title:        "aboba",
		Description:  "",
		AgeRange:     [2]uint{},
		PreviewUUID:  "",
		MaxListeners: 12,
		Timetable:    nil,
		From:         time.Time{},
		To:           time.Time{},
		ExpiresAt:    time.Now(),
	}

	var params map[string]any

	b, err := json.Marshal(course)
	if err != nil {
		t.Error(err)
	}

	err = json.Unmarshal(b, &params)
	if err != nil {
		t.Error(err)
	}

	for s, a := range params {
		if reflect.DeepEqual(a, reflect.Zero(reflect.TypeOf(a)).Interface()) {
			log.Println("Empty:", s)
		} else {
			log.Printf("%v: %v", s, a)
		}
	}
}
