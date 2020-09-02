package usecase_test

import (
	"testing"

	"github.com/ariefsam/telemarketing/usecase"
	"github.com/ariefsam/telemarketing/usecase/dependency/mockdependency"
)

func TestCreateTelemarketer(t *testing.T) {
	var u usecase.Usecase
	var mockTelemarketerRepository mockdependency.TelemarketerRepository
	u.TelemarketerRepository = &mockTelemarketerRepository

}
