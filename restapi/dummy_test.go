package restapi_test

import "github.com/ariefsam/telemarketing/entity"

func dummyTelemarketerNotAdmin() (expectedTelemarketer entity.Telemarketer) {
	expectedTelemarketer = entity.Telemarketer{
		Name:    "Arief",
		Email:   "arief@gmail.com",
		IsAdmin: false,
	}
	return
}

func dummyTelemarketerAdmin() (expectedTelemarketer entity.Telemarketer) {
	expectedTelemarketer = entity.Telemarketer{
		Name:    "Arief",
		Email:   "arief@gmail.com",
		IsAdmin: true,
	}
	return
}

func dummyCustomers() []entity.Customer {
	return []entity.Customer{
		entity.Customer{
			Name:        "A",
			PhoneNumber: "09123",
		},
		entity.Customer{
			Name:        "B",
			PhoneNumber: "091232",
		},
	}
}
