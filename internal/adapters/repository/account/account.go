package account

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/frameworks/driven/databases"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
)

type accountRepo struct {
	ports.Database
}

func NewAccountRepo(log logger.Logger, config interface{}) ports.AccountRepo {
	db, err := databases.NewDatabase("", log, config)
	if err != nil {
		log.Errorf("failed to initialise schedule repo: %v", err)
		return nil
	}
	return &accountRepo{
		db,
	}
}
