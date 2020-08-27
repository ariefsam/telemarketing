package usecase_eventsource

import (
	"time"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/usecase"
	"github.com/pkg/errors"
)

type Usecase struct {
	usecase.Usecase
}

type Event struct {
	Name      string
	Data      []interface{}
	Timestamp int64
}

func save(event string, data []interface{}) (err error) {
	ctx, docRef, err := getFirestoreClient()
	if err != nil {
		return
	}
	dataInsert := Event{
		Name:      event,
		Data:      data,
		Timestamp: time.Now().UnixNano(),
	}

	_, err = docRef.Collection("Event").Doc(generateID()).Set(ctx, dataInsert)
	if err != nil {
		err = errors.Wrap(err, "Failed to save data")
		return
	}
	return
}

func (u *Usecase) SaveCustomer(customer entity.Customer) (err error) {
	err = save("SaveCustomer", []interface{}{customer})
	return
}

/*
AssignCustomer(telemarketerEmail string) (customer entity.Customer, err error)
	GetCallLog(filter entity.FilterCallLog, limit int) (callLogs []entity.CallLog, err error)
	GetCustomer(filter entity.FilterCustomer, limit int) (customers []entity.Customer, err error)
	SaveCallLog(callLog entity.CallLog) (err error)
	SaveCustomer(customer entity.Customer) (err error)
	ParseToken(token string) (isValid bool, telemarketer entity.Telemarketer)
	LoginByFirebase(firebaseToken string) (token string, telemarketer entity.Telemarketer, isValid bool, err error)
	CurrentTimestamp() int64
	SaveTelemarketer(telemarketer entity.Telemarketer) (err error)
	GetTelemarketer(filter entity.FilterTelemarketer, limit int) ([]entity.Telemarketer, error)*/
