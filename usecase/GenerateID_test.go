package usecase_test

import (
	"testing"

	"github.com/ariefsam/telemarketing/usecase"
	"github.com/ariefsam/telemarketing/usecase/dependency/mockdependency"
	"github.com/stretchr/testify/assert"
)

func TestGenerateID(t *testing.T) {
	var u usecase.Usecase
	var mockIDGenerator mockdependency.IDGenerator
	u.IDGenerator = &mockIDGenerator
	id := "expectedID"
	mockIDGenerator.On("Generate").Return(id)
	assert.Equal(t, id, u.GenerateID())
}
