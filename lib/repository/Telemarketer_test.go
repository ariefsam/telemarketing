package repository_test

import (
	"fmt"
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

	telemarketer2 := entity.Telemarketer{
		Email: "arief2@gmail.com",
	}
	err = telemarketerRepository.Save(telemarketer2)
	assert.NoError(t, err)

	getTelemarketer, err = telemarketerRepository.Get(entity.FilterTelemarketer{}, 100)
	assert.NoError(t, err)
	assert.True(t, len(getTelemarketer) > 1, "Failed get telemarketers, length: "+fmt.Sprintf("%d", len(getTelemarketer)))

	telemarketerRepository.Delete(telemarketer.Email)
	telemarketerRepository.Delete(telemarketer2.Email)
}
