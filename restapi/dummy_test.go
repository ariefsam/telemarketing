package restapi_test

import "github.com/ariefsam/telemarketing/entity"

func dummyTelemarketerNotAdmin() (expectedTelemarketer entity.Telemarketer) {
	expectedTelemarketer = entity.Telemarketer{
		ID:      "idtelemarketer01",
		Name:    "Arief",
		Email:   "arief@gmail.com",
		IsAdmin: false,
	}
	return
}

func dummyTelemarketerAdmin() (expectedTelemarketer entity.Telemarketer) {
	expectedTelemarketer = entity.Telemarketer{
		ID:      "idtelemarketeradmin01",
		Name:    "Arief",
		Email:   "arief@gmail.com",
		IsAdmin: true,
	}
	return
}

func dummyCustomers() []entity.Customer {
	return []entity.Customer{
		entity.Customer{
			ID:          "idcutomer01",
			Name:        "A",
			PhoneNumber: "09123",
		},
		entity.Customer{
			ID:          "idcutomer02",
			Name:        "B",
			PhoneNumber: "091232",
		},
	}
}
