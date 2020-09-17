package usecase

import (
	"log"

	"github.com/ariefsam/telemarketing/entity"
)

func (u *Usecase) Call(telemarketer entity.Telemarketer, customer entity.Customer, status string, timestamp int64) (err error) {
	customer.Status = status
	id := u.IDGenerator.Generate()
	callLog := entity.CallLog{
		ID:             id,
		Name:           customer.Name,
		PhoneNumber:    customer.PhoneNumber,
		TelemarketerID: telemarketer.ID,
		Timestamp:      timestamp,
		Status:         status,
	}
	err = u.CallLogRepository.Save(callLog)
	if err != nil {
		log.Println(err)
	}
	u.CustomerRepository.Save(customer)

	telemarketer.Performance.Daily.Call += 1
	telemarketer.Performance.Weekly.Call += 1
	telemarketer.Performance.Monthly.Call += 1
	u.TelemarketerRepository.Save(telemarketer)
	return
}

func (u *Usecase) ValidateCall() {
	return
}
