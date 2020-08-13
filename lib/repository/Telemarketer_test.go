package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/lib/repository"
)

func TestTelemarketer(t *testing.T) {
	var telemarketerRepository repository.Telemarketer
	telemarketer := entity.Telemarketer{
		Email: "arief@gmail.com",
	}
	err := telemarketerRepository.Save(telemarketer)
	assert.NoError(t, err)
	filter := entity.FilterTelemarketer{
		Email: "arief@gmail.com",
	}
	getTelemarketer, err := telemarketerRepository.Get(filter, 1)
	assert.NoError(t, err)
	assert.Equal(t, telemarketer, getTelemarketer[0])
	telemarketerRepository.Delete(telemarketer.Email)
}
