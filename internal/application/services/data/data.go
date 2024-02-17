package data

import (
	"github.com/shashank-priyadarshi/upgraded-disco/models"
	logger "github.com/shashank-priyadarshi/utilities/logger/ports"
)

type Service struct {
	db models.Repository
}

// TODO
func (s Service) Chess() (interface{}, error) {
	return s.db.MongoDB.Query(nil), nil
}

func (s Service) GitHub() (interface{}, error) {
	return s.db.MongoDB.Query(nil), nil
}

func NewApplication(log logger.Logger, database interface{}) *Service {
	log.Info("Initialising data service")

	return &Service{
		db: database.(models.Repository),
	}
}
