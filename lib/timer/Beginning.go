package timer

import (
	"time"
)

func (t Time) BeginningMonth() Time {
	var tr Time
	tr.Time = t.Add(-24 * time.Duration(t.Day()-1) * time.Hour)
	tr = tr.BeginningDay()
	tr.Time = time.Unix(tr.Time.Unix(), 0)
	return tr
}

func (t Time) BeginningWeek() Time {
	var tr Time
	weekday := t.Weekday()
	tr.Time = t.Add(-24 * time.Duration(weekday) * time.Hour)
	tr = tr.BeginningDay()
	tr.Time = time.Unix(tr.Time.Unix(), 0)
	return tr
}

func (t Time) BeginningDay() Time {
	var tr Time
	tr.Time = t.Add(-1 * time.Duration(t.Hour()) * time.Hour).Add(-1 * time.Duration(t.Minute()) * time.Minute).Add(-1 * time.Duration(t.Second()) * time.Second)
	tr.Time = time.Unix(tr.Time.Unix(), 0)
	return tr
}
