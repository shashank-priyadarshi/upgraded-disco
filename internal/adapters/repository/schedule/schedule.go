package schedule

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/frameworks/driven/databases"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
)

type scheduleRepo struct {
	ports.Database
}

func NewScheduleRepo(log logger.Logger, config interface{}) ports.ScheduleRepo {
	db, err := databases.NewDatabase("", log, config)
	if err != nil {
		log.Errorf("failed to initialise schedule repo: %v", err)
		return nil
	}
	return &scheduleRepo{
		db,
	}
}
