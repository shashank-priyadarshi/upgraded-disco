package plugin

import "github.com/shashank-priyadarshi/upgraded-disco/utils/logger"

type Service struct {
	db interface{}
}

func (s Service) Install(i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) List() (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Info(s3 string, s2 string, i ...interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Upgrade(i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) Trigger(i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) Uninstall(s2 string) error {
	//TODO implement me
	panic("implement me")
}

func NewApplication(log logger.Logger, database interface{}) *Service {

	log.Infof("Initialising plugin service")
	return &Service{
		db: database,
	}
}
