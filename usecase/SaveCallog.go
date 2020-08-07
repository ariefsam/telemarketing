package usecase

import "github.com/ariefsam/telemarketing/entity"

func (u *Usecase) SaveCallLog(callLog entity.CallLog) (err error) {
	err = u.CallLogRepository.Save(callLog)
	return
}
