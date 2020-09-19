package timer

import (
	"log"
	"time"
)

func (t *Time) BeginningWeek() Time {
	var tr Time
	weekday := t.Weekday()
	hour := t.Hour()
	log.Println(hour)
	tr.Time = t.Add(-24 * time.Duration(weekday) * time.Hour).Add(-1 * time.Duration(hour) * time.Hour).Add(-1 * time.Duration(t.Minute()) * time.Minute).Add(-1 * time.Duration(t.Second()) * time.Second)

	return tr
}
