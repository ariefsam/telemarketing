package usecase

import "github.com/ariefsam/telemarketing/entity"

func (u *Usecase) GetCallLog(filter entity.FilterCallLog, limit int) (callLogs []entity.CallLog, err error) {
	callLogs, err = u.CallLogRepository.Get(filter, limit)
	return
}
