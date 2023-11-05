package repository

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
	"github.com/shashank-priyadarshi/upgraded-disco/models"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
)

type Repository struct {
	log                     logger.Logger
	Cache, MariaDB, MongoDB ports.Database
}

func NewRepository(log logger.Logger) *Repository {
	return &Repository{
		log: log,
	}
}

func (r *Repository) WithRedisCache(config models.DBConfig) *Repository {
	return &Repository{}
}

func (r *Repository) WithMariaDB(config models.DBConfig) *Repository {
	return &Repository{}
}

func (r *Repository) WithMongoDB(config models.DBConfig) *Repository {
	return &Repository{}
}
