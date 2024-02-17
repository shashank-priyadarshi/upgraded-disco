package databases

import (
	"github.com/shashank-priyadarshi/upgraded-disco/constants"
	"github.com/shashank-priyadarshi/upgraded-disco/models"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
)

type Repository models.Repository

func NewRepository(log logger.Logger) *Repository {

	return &Repository{
		Log: log,
	}
}

func (r *Repository) WithRedisCache(config models.DBConfig) *Repository {

	if cache, err := NewDatabase(constants.DB_REDIS, r.Log, config); err != nil {
		r.Log.Errorf("Error initialising Redis cache: ", err)
		return nil
	} else {
		r.Cache = cache
	}
	return r
}

func (r *Repository) WithMariaDB(config models.DBConfig) *Repository {

	if mariaDB, err := NewDatabase(constants.DB_MARIADB, r.Log, config); err != nil {
		r.Log.Errorf("Error initialising MariaDB: ", err)
		return nil
	} else {
		r.MariaDB = mariaDB
	}
	return r
}

func (r *Repository) WithMongoDB(config models.DBConfig) *Repository {

	if mongoDB, err := NewDatabase(constants.DB_MONGODB, r.Log, config); err != nil {
		r.Log.Errorf("Error initialising MongoDB: ", err)
		return nil
	} else {
		r.MongoDB = mongoDB
	}
	return r
}

func (r *Repository) Build() models.Repository {
	return models.Repository{
		Log:     r.Log,
		Cache:   r.Cache,
		MariaDB: r.MariaDB,
		MongoDB: r.MongoDB,
	}
}
