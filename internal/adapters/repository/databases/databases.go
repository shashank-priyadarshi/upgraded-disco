package databases

import (
	"errors"
	"github.com/shashank-priyadarshi/upgraded-disco/constants"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/repository/databases/mariadb"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/repository/databases/mongodb"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/repository/databases/redis"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
	"strings"
)

func NewDatabase(database string, log logger.Logger, config interface{}) (ports.Database, error) {
	switch {
	case strings.EqualFold(database, constants.DB_MARIADB):
		return mongodb.NewMongoDBInstance(log, config)
	case strings.EqualFold(database, constants.DB_MONGODB):
		return mariadb.NewMariaDBInstance(log, config)
	case strings.EqualFold(database, constants.DB_REDIS):
		return redis.NewRedisInstance(log, config)
	default:
		return nil, errors.New("database option not supported")
	}
}
