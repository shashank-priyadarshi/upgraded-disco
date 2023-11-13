package repository

import (
	"github.com/shashank-priyadarshi/upgraded-disco/constants"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/repository/databases"
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
	r.Cache, _ = databases.NewDatabase(constants.DB_REDIS, r.Log, config)
	return r
}

func (r *Repository) WithMariaDB(config models.DBConfig) *Repository {
	r.MariaDB, _ = databases.NewDatabase(constants.DB_MARIADB, r.Log, config)
	return r
}

func (r *Repository) WithMongoDB(config models.DBConfig) *Repository {
	r.MongoDB, _ = databases.NewDatabase(constants.DB_MONGODB, r.Log, config)
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
