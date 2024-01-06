package graphql

import "github.com/shashank-priyadarshi/upgraded-disco/utils/logger"

type Service struct {
	db interface{}
}

// TODO
func (s Service) GraphQL(query interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func NewApplication(log logger.Logger, database interface{}) *Service {

	log.Infof("Initialising GraphQL service")
	return &Service{
		db: database,
	}
}
