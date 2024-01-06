package plugin

import "github.com/shashank-priyadarshi/upgraded-disco/utils/logger"

type Service struct {
	db interface{}
}

// TODO
func (s Service) List() (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Update(interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) Install(interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) Trigger(interface{}) error {
	//TODO implement me
	panic("implement me")
}

func NewApplication(log logger.Logger, database interface{}) *Service {

	log.Infof("Initialising plugin service")
	return &Service{
		db: database,
	}
}
