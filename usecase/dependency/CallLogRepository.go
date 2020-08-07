package dependency

import "github.com/ariefsam/telemarketing/entity"

type CallLogRepository interface {
	Save(callLog entity.CallLog) (err error)
	Get(filter entity.FilterCallLog, limit int) ([]entity.CallLog, error)
}
