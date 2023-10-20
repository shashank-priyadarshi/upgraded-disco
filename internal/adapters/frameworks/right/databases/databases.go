package databases

import (
	"errors"
	"strings"

	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/constants"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/frameworks/right/databases/mariadb"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/frameworks/right/databases/mongodb"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/frameworks/right/databases/redis"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports/frameworks/right/database"
)

func NewDatabase(database string, log, config interface{})(database.Database,error){
	switch{
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