package usecase

import (
	"time"

	"github.com/ariefsam/telemarketing/lib/repository"
	"github.com/ariefsam/telemarketing/usecase/dependency"
)

type Usecase struct {
	CustomerRepository     dependency.CustomerRepository
	CallLogRepository      dependency.CallLogRepository
	TokenService           dependency.TokenService
	FirebaseAuth           dependency.FirebaseAuth
	TelemarketerRepository dependency.TelemarketerRepository
	Timer                  dependency.Timer
	IDGenerator            dependency.IDGenerator
	ReportRepository       *repository.Report
}

var loc *time.Location

func init() {
	loc, _ = time.LoadLocation("Asia/Jakarta")
}
