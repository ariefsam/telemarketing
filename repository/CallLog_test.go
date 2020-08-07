package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/repository"
)

func TestCallLog(t *testing.T) {
	var c repository.CallLog
	callLog := entity.CallLog{
		Name:           "Member 01",
		PhoneNumber:    "0123",
		Status:         "Tidak diangkat",
		TelemarketerID: "tele001",
		Timestamp:      1230000,
	}

	err := c.Save(callLog)
	assert.NoError(t, err)

	callLog2 := entity.CallLog{
		PhoneNumber:    "0123",
		Status:         "Tertarik",
		TelemarketerID: "tele001",
		Timestamp:      1240000,
	}

	err = c.Save(callLog2)
	assert.NoError(t, err)

	callLog3 := entity.CallLog{
		PhoneNumber:    "0123",
		Status:         "Tertarik",
		TelemarketerID: "tele001",
		Timestamp:      1250000,
	}

	err = c.Save(callLog3)
	assert.NoError(t, err)

	filter := entity.FilterCallLog{
		TelemarketerID: callLog2.TelemarketerID,
		Status:         callLog2.Status,
		TimestampStart: 1231000,
		TimestampEnd:   1250000,
	}
	callLogs, err := c.Get(filter, 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(callLogs))
	if len(callLogs) == 1 {
		assert.Equal(t, []entity.CallLog{callLog2}, callLogs)
	}
}
