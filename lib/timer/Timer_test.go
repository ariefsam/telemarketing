package timer_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/ariefsam/telemarketing/lib/timer"
	"github.com/stretchr/testify/assert"
)

func TestTimer(t *testing.T) {
	var tm timer.Timer
	now := time.Now().UnixNano()
	expectedNow := tm.CurrentTimestamp()
	expectedThen := tm.CurrentTimestamp()
	assert.True(t, expectedNow >= now)
	assert.True(t, expectedNow < expectedThen, fmt.Sprintf("%d - %d", expectedNow, expectedThen))
}

func TestGenerateNumberCode(t *testing.T) {
	i := 0
	num := [1000]int{}
	for {
		x := timer.GenerateNumberCode(3)
		num[x]++
		if num[x] > 1 {
			log.Println(x, num[x])
		}
		i++
		if i == 10 {
			break
		}
	}
}
