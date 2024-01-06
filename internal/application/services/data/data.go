package data

import (
	"github.com/shashank-priyadarshi/upgraded-disco/models"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
)

type Service struct {
	db models.Repository
}

// TODO
func (s Service) GetGraphData() (interface{}, error) {
	return s.db.MongoDB.Get(nil) // Provide input to Get as models.MongoDBPayload, with non-nil db query in the data property
}

func (s Service) GetGitHubData() (interface{}, error) {
	return s.db.MongoDB.Get(nil)
}

func NewApplication(log logger.Logger, database interface{}) *Service {
	log.Infof("Initialising data service")

	return &Service{
		db: database.(models.Repository),
	}
}
