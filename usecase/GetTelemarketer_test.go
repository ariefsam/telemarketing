package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/usecase"
	"github.com/ariefsam/telemarketing/usecase/dependency/mockdependency"
)

func TestGetTelemarketer(t *testing.T) {
	var u usecase.Usecase
	var mockTelemarketerRepository mockdependency.TelemarketerRepository
	u.TelemarketerRepository = &mockTelemarketerRepository

	telemarketers := []entity.Telemarketer{
		entity.Telemarketer{
			Name:  "A",
			Email: "x@gmail.com",
		},
	}
	filter := entity.FilterTelemarketer{}
	limit := 10
	mockTelemarketerRepository.On("Get", filter, limit).Return(telemarketers, nil)
	getTelemarketers, err := u.GetTelemarketer(filter, limit)
	assert.NoError(t, err)
	assert.Equal(t, telemarketers, getTelemarketers)

}
