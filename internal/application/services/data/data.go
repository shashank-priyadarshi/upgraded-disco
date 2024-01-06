package data

import (
	"github.com/shashank-priyadarshi/upgraded-disco/models"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
)

type Service struct {
	db models.Repository
}

// TODO
func (s Service) Chess() (interface{}, error) {
	return s.db.MongoDB.Get(nil)
}

func (s Service) GitHub() (interface{}, error) {
	return s.db.MongoDB.Get(nil)
}

func NewApplication(log logger.Logger, database interface{}) *Service {
	log.Infof("Initialising data service")

	return &Service{
		db: database.(models.Repository),
	}
}
