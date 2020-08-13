package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/telemarketing/usecase"
	"github.com/ariefsam/telemarketing/usecase/dependency/mockdependency"
)

func TestCurrentTimestamp(t *testing.T) {
	var u usecase.Usecase
	var mockTimer mockdependency.Timer
	var expectedTime int64
	expectedTime = 1231231
	mockTimer.On("CurrentTimestamp").Return(expectedTime)
	u.Timer = &mockTimer
	getTime := u.CurrentTimestamp()
	assert.Equal(t, expectedTime, getTime)
}
