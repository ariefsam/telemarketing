package timer

import (
	"time"
)

type Timer struct{}

func (t *Timer) CurrentTimestamp() int64 {
	ts := time.Now().UnixNano()
	return ts
}
