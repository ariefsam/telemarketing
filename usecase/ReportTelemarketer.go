package usecase

import (
	"time"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/lib/timer"
)

func (u *Usecase) ReportTelemarketer(filter entity.FilterReportTelemarketer, limit int) (telemarketers []entity.Telemarketer, err error) {
	var filterTelemarketer entity.FilterTelemarketer
	f := false
	filterTelemarketer.IsAdmin = &f
	telemarketers, err = u.TelemarketerRepository.Get(filterTelemarketer, limit)
	// log.Printf("Log telemarketer, %+v", filter.ReportTimestamp)
	for idx, telemarketer := range telemarketers {
		var t timer.Time
		t.Time = time.Now()
		telemarketers[idx].Performance = u.ReportRepository.Get(telemarketer.ID, t)
	}
	return
}
