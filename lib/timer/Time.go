package timer

import "time"

type Time struct {
	time.Time
}

func NewFromTime(t time.Time) Time {
	return Time{Time: t}
}

func Now() (t Time) {
	return NewFromTime(time.Now())
}

func Unix(second, nanosecond int64) (t Time) {
	return NewFromTime(time.Unix(second, nanosecond))
}

func (t Time) In(loc *time.Location) Time {
	return NewFromTime(t.Time.In(loc))
}
