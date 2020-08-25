package usecaseinterface

import "github.com/ariefsam/telemarketing/entity"

type Usecase interface {
	AssignCustomer(telemarketerID string) (customer entity.Customer, err error)
	GetCallLog(filter entity.FilterCallLog, limit int) (callLogs []entity.CallLog, err error)
	GetCustomer(filter entity.FilterCustomer, limit int) (customers []entity.Customer, err error)
	SaveCallLog(callLog entity.CallLog) (err error)
	SaveCustomer(customer entity.Customer) (err error)
	ParseToken(token string) (isValid bool, telemarketer entity.Telemarketer)
	LoginByFirebase(firebaseToken string) (token string, telemarketer entity.Telemarketer, isValid bool, err error)
	CurrentTimestamp() int64
	SaveTelemarketer(telemarketer entity.Telemarketer) (err error)
	GetTelemarketer(filter entity.FilterTelemarketer, limit int) ([]entity.Telemarketer, error)
}
