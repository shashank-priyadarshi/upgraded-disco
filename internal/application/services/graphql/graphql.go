package graphql

import logger "github.com/shashank-priyadarshi/utilities/logger/ports"

type Service struct {
	db interface{}
}

// TODO
func (s Service) GraphQL(query interface{}) (interface{}, error) {
	//TODO implement me
	return nil, nil
}

func NewApplication(log logger.Logger, database interface{}) *Service {

	log.Info("Initialising GraphQL service")
	return &Service{
		db: database,
	}
}
