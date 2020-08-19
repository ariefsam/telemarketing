package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/usecase"
	"github.com/ariefsam/telemarketing/usecase/dependency/mockdependency"
)

func TestSaveTelemarketer(t *testing.T) {
	var u usecase.Usecase
	var mockTelemarketer mockdependency.TelemarketerRepository
	telemarketer := entity.Telemarketer{
		Email: "x@gmail.com",
		Name:  "Andi",
	}
	u.TelemarketerRepository = &mockTelemarketer
	mockTelemarketer.On("Save", telemarketer).Return(nil)

	err := u.SaveTelemarketer(telemarketer)
	assert.NoError(t, err)
	mockTelemarketer.AssertCalled(t, "Save", telemarketer)

}
