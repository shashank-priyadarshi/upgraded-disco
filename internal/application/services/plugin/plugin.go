package plugin

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

func (s Service) Update() {
	//TODO implement me
	panic("implement me")
}

func (s Service) Install() {
	//TODO implement me
	panic("implement me")
}

func (s Service) Trigger() {
	//TODO implement me
	panic("implement me")
}

func NewApplication(database ports.ServiceRepository) *Service {
	return &Service{
		db: database,
	}
}
