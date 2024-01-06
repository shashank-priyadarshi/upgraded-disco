package schedule

import "github.com/shashank-priyadarshi/upgraded-disco/utils/logger"

type Service struct {
	db interface{}
}

func (s Service) Create(schedule interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) List(schedule interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Update(schedule interface{}) error { return nil }

func (s Service) Delete(schedule interface{}) error {
	//TODO implement me
	panic("implement me")
}

func NewApplication(log logger.Logger, database interface{}) *Service {

	log.Infof("Initialising schedule service")
	return &Service{
		db: database,
	}
}
