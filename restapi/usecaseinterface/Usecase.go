package usecaseinterface

import "github.com/ariefsam/telemarketing/entity"

type Usecase interface {
	AssignCustomer(telemarketerID string, customerID string) (err error)
	ValidateAssignCustomer(telemarketerID string, customerID string) (err error)
	GetCallLog(filter entity.FilterCallLog, limit int) (callLogs []entity.CallLog, err error)
	GetCustomer(filter entity.FilterCustomer, limit int) (customers []entity.Customer, err error)
	SaveCallLog(callLog entity.CallLog) (err error)
	ClosingCustomer(customer entity.Customer) (err error)
	CreateCustomer(customer entity.Customer) (err error)
	Call(telemarketer entity.Telemarketer, customer entity.Customer, status string, currentTimestamp int64) (err error)
	ParseToken(token string) (isValid bool, telemarketer entity.Telemarketer)
	LoginByFirebase(firebaseToken string) (token string, telemarketer entity.Telemarketer, isValid bool, err error)
	CurrentTimestamp() int64
	CreateTelemarketer(telemarketer entity.Telemarketer) (err error)
	SaveTelemarketer(telemarketer entity.Telemarketer) (err error)
	GetTelemarketer(filter entity.FilterTelemarketer, limit int) ([]entity.Telemarketer, error)
	ReportTelemarketer(filter entity.FilterReportTelemarketer, limit int) ([]entity.Telemarketer, error)
	GenerateID() (id string)
}
