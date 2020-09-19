package timer_test

import (
	"testing"
	"time"

	"github.com/ariefsam/telemarketing/lib/timer"
	"github.com/stretchr/testify/assert"
)

func TestBeginningWeek(t *testing.T) {
	var todayunix, expectedFirstWeekUnix int64
	todayunix = 1600497808 //Sabtu, 19 September 2020 pukul 13.43.28 GMT+07:00
	loc, _ := time.LoadLocation("Asia/Jakarta")
	timerToTest := timer.Unix(todayunix, 0).In(loc)

	expectedFirstWeekUnix = 1599930000 //Minggu, 13 September 2020 pukul 00.00.00 GMT+07:00

	timerToTest = timerToTest.BeginningWeek()
	assert.Equal(t, expectedFirstWeekUnix, timerToTest.Unix())

}
