package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/usecase"
	"github.com/ariefsam/telemarketing/usecase/dependency/mockdependency"
)

func TestCreateTelemarketer(t *testing.T) {
	var u usecase.Usecase
	var mockTelemarketerRepository mockdependency.TelemarketerRepository
	u.TelemarketerRepository = &mockTelemarketerRepository
	setupTelemarketerByEmailNotExist(&mockTelemarketerRepository, "xx@gmail.com")
	telemarketer := entity.Telemarketer{
		Email: "xx@gmail.com",
	}
	mockTelemarketerRepository.On("Save", telemarketer).Return(nil)
	u.CreateTelemarketer(telemarketer)
	mockTelemarketerRepository.AssertCalled(t, "Save", telemarketer)
}

func TestCreateTelemarketerAlreadyExist(t *testing.T) {
	var u usecase.Usecase
	var mockTelemarketerRepository mockdependency.TelemarketerRepository
	u.TelemarketerRepository = &mockTelemarketerRepository

	setupTelemarketerByEmailExist(&mockTelemarketerRepository, "xx@gmail.com")

	telemarketer := entity.Telemarketer{
		Email: "xx@gmail.com",
		Name:  "anu",
	}
	mockTelemarketerRepository.On("Save", telemarketer).Return(nil)
	err := u.CreateTelemarketer(telemarketer)
	assert.Error(t, err)
	mockTelemarketerRepository.AssertNotCalled(t, "Save", telemarketer)
}

func setupTelemarketerByEmailExist(mockTelemarketerRepository *mockdependency.TelemarketerRepository, email string) {
	existTelemarketer := entity.Telemarketer{
		Email: email,
		Name:  "ani",
	}

	filter := entity.FilterTelemarketer{
		Email: &email,
	}
	mockTelemarketerRepository.On("Get", filter, 1).Return([]entity.Telemarketer{existTelemarketer}, nil)
}

func setupTelemarketerByEmailNotExist(mockTelemarketerRepository *mockdependency.TelemarketerRepository, email string) {
	filterEmail := "xx@gmail.com"
	filter := entity.FilterTelemarketer{
		Email: &filterEmail,
	}
	mockTelemarketerRepository.On("Get", filter, 1).Return([]entity.Telemarketer{}, nil)
}
