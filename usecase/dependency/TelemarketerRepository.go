package dependency

import "github.com/ariefsam/telemarketing/entity"

type TelemarketerRepository interface {
	Save(telemarketer entity.Telemarketer) (err error)
	Get(filter entity.FilterTelemarketer, limit int) ([]entity.Telemarketer, error)
}
