package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/repository"
)

func TestCustomer(t *testing.T) {
	var c repository.Customer
	customer := entity.Customer{
		Name:           "Member 01",
		PhoneNumber:    "01234",
		TelemarketerID: "tele001",
		Status:         "Tertarik",
	}
	err := c.Save(customer)

	customer2 := entity.Customer{
		PhoneNumber:    "012345",
		TelemarketerID: "tele001",
		Status:         "Tidak diangkat",
	}
	err = c.Save(customer2)
	assert.NoError(t, err)
	filter := entity.FilterCustomer{
		PhoneNumber: customer.PhoneNumber,
	}

	customers, err := c.Get(filter, 1)
	assert.Equal(t, 1, len(customers))
	if len(customers) == 1 {
		assert.Equal(t, []entity.Customer{customer}, customers)
	}

	filter = entity.FilterCustomer{
		TelemarketerID: customer2.TelemarketerID,
		Status:         customer2.Status,
	}
	customers, err = c.Get(filter, 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(customers))
	if len(customers) == 1 {
		assert.Equal(t, []entity.Customer{customer2}, customers)
	}

}
