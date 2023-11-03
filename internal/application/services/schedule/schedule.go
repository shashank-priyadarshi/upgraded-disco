package schedule

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
)

type Service struct {
	db ports.ServiceRepository
}

func (s Service) List() {
	//TODO implement me
	panic("implement me")
}

func (s Service) Create() {
	//TODO implement me
	panic("implement me")
}

func (s Service) Delete() {
	//TODO implement me
	panic("implement me")
}

func NewApplication(database ports.ServiceRepository) *Service {
	return &Service{
		db: database,
	}
}
