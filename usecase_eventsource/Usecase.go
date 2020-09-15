package usecase_eventsource

import (
	"log"
	"time"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/usecase"
	"github.com/pkg/errors"
)

type UsecaseEvent struct {
	usecase.Usecase
}

type Event struct {
	Name      string
	Data      interface{}
	Timestamp int64
}

func save(event string, data interface{}) (err error) {
	ctx, docRef, err := getFirestoreClient()
	if err != nil {
		return
	}
	dataInsert := Event{
		Name:      event,
		Data:      data,
		Timestamp: time.Now().UnixNano(),
	}
	eventID := generateID()
	_, err = docRef.Collection("Event").Doc(eventID).Set(ctx, dataInsert)
	if err != nil {
		err = errors.Wrap(err, "Failed to save data: "+eventID)
		log.Println(err)
		return
	}
	return
}

func (u *UsecaseEvent) CreateTelemarketer(telemarketer entity.Telemarketer) (err error) {
	err = u.Usecase.ValidateCreateTelemarketer(telemarketer)
	log.Println("Create telemarketer")
	if err != nil {
		log.Println(err)
		return
	}
	err = save("CreateTelemarketer", telemarketer)
	return
}

func (u *UsecaseEvent) CreateCustomer(customer entity.Customer) (err error) {
	err = u.Usecase.ValidateCreateCustomer(customer)
	if err != nil {
		return
	}
	err = save("CreateCustomer", customer)
	return
}

func (u *UsecaseEvent) AssignCustomer(telemarketerID string, customerID string) (err error) {
	err = u.Usecase.ValidateAssignCustomer(telemarketerID, customerID)
	if err != nil {
		return
	}
	data := []string{telemarketerID, customerID}
	err = save("AssignCustomer", data)
	return
}

func (u *UsecaseEvent) SaveTelemarketer(telemarketer entity.Telemarketer) (err error) {
	err = save("SaveTelemarketer", telemarketer)
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
