package usecase

import (
	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/lib/timer"
)

func (u *Usecase) GetTelemarketer(filter entity.FilterTelemarketer, limit int) (telemarketers []entity.Telemarketer, err error) {
	telemarketers, err = u.TelemarketerRepository.Get(filter, limit)
	// log.Printf("Log telemarketer, %+v", filter.ReportTimestamp)
	for idx, telemarketer := range telemarketers {
		var t timer.Time
		if filter.ReportTimestamp == nil {
			t = timer.Unix(0, u.CurrentTimestamp())
		} else {
			t = timer.Unix(0, *filter.ReportTimestamp)
			// log.Println("Filter Timestamp ", t.Unix())
		}
		telemarketers[idx].Performance = u.ReportRepository.Get(telemarketer.ID, t)
	}
	return
}
