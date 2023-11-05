package databases

import (
	"errors"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
	"strings"

	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/constants"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/frameworks/driven/databases/mariadb"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/frameworks/driven/databases/mongodb"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/frameworks/driven/databases/redis"
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
