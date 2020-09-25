package ioc_eventsource

import (
	"os"
	"strconv"

	"github.com/ariefsam/telemarketing/ioc"
	"github.com/ariefsam/telemarketing/usecase_eventsource"
)

func Usecase() (usecase usecase_eventsource.UsecaseEvent) {
	redis_db := os.Getenv("redis_db")
	usecase_eventsource.RedisDB, _ = strconv.Atoi(redis_db)
	usecase_eventsource.FirebaseAccountPath = os.Getenv("firebase_credential_path")
	usecase_eventsource.EventSourceDB = os.Getenv("event_source_db")
	usecase.Usecase = ioc.Usecase()
	return
}
