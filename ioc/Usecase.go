package ioc

import (
	"os"
	"strconv"

	"github.com/ariefsam/telemarketing/lib/firebaseauth"
	"github.com/ariefsam/telemarketing/lib/idgenerator"
	"github.com/ariefsam/telemarketing/lib/repository"
	"github.com/ariefsam/telemarketing/lib/timer"
	"github.com/ariefsam/telemarketing/lib/token_telemarketer"
	"github.com/ariefsam/telemarketing/usecase"
)

func Usecase() (usecase usecase.Usecase) {
	redis_db := os.Getenv("redis_db")
	repository.RedisDB, _ = strconv.Atoi(redis_db)
	repository.InitRedis()
	repository.FirebaseAccountPath = os.Getenv("firebase_credential_path")
	repository.ProjectionDB = os.Getenv("projection_db")

	usecase.CustomerRepository = &repository.Customer{}
	usecase.CallLogRepository = &repository.CallLog{}
	usecase.TokenService = &token_telemarketer.Token{}
	usecase.FirebaseAuth = &firebaseauth.FirebaseAuth{
		CredentialPath: os.Getenv("firebase_credential_path"),
	}
	usecase.TelemarketerRepository = &repository.Telemarketer{}
	usecase.Timer = &timer.Timer{}
	usecase.IDGenerator = &idgenerator.IDGenerator{}
	usecase.ReportRepository = &repository.Report{}
	return
}
