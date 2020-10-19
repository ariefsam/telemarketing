package usecase

import (
	"log"
	"time"

	"github.com/ariefsam/telemarketing/entity"
	"github.com/ariefsam/telemarketing/lib/timer"
)

func (u *Usecase) ReportTelemarketer(filter entity.FilterReportTelemarketer, limit int) (telemarketers []entity.Telemarketer, err error) {
	var filterTelemarketer entity.FilterTelemarketer
	f := false
	filterTelemarketer.IsAdmin = &f
	telemarketers, err = u.TelemarketerRepository.Get(filterTelemarketer, limit)
	now := time.Now().UnixNano()
	if filter.TimestampStart == nil {
		filter.TimestampStart = &now
	}
	if filter.TimestampEnd == nil {
		filter.TimestampEnd = &now
	}

	for idx, telemarketer := range telemarketers {
		t := timer.Unix(0, *filter.TimestampStart)
		end := timer.Unix(0, *filter.TimestampEnd)
		var i int
		for {
			i++
			log.Println("Counting ", t.Time.Format("20060102"))
			telemarketers[idx].Performance.Daily.BuyAmount += u.ReportRepository.Get(telemarketer.ID, t).Daily.BuyAmount
			telemarketers[idx].Performance.Daily.Call += u.ReportRepository.Get(telemarketer.ID, t).Daily.Call
			telemarketers[idx].Performance.Daily.Closing += u.ReportRepository.Get(telemarketer.ID, t).Daily.Closing
			t.Time = t.Time.Add(24 * time.Hour)
			if t.UnixNano() >= end.UnixNano() {

				log.Println("Total i: ", i)
				break
			}
		}
		telemarketers[idx].Target.Daily.BuyAmount *= i
		telemarketers[idx].Target.Daily.Closing *= i
		telemarketers[idx].Target.Daily.Call *= i
	}
	return
}
