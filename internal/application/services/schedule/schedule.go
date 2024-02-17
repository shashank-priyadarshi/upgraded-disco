package schedule

import logger "github.com/shashank-priyadarshi/utilities/logger/ports"

type Service struct {
	db interface{}
}

func (s Service) Create(schedule interface{}) (interface{}, error) {
	//TODO implement me
	return nil, nil
}

func (s Service) List(schedule interface{}) (interface{}, error) {
	//TODO implement me
	return nil, nil
}

func (s Service) Update(schedule interface{}) error { return nil }

func (s Service) Delete(schedule interface{}) error {
	//TODO implement me
	return nil
}

func NewApplication(log logger.Logger, database interface{}) *Service {

	log.Info("Initialising schedule service")
	return &Service{
		db: database,
	}
}
