package data

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
)

type Service struct {
	db ports.ServiceRepository
}

func (s Service) GetGraphData() {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetGitHubData() {
	//TODO implement me
	panic("implement me")
}

func NewApplication(database ports.ServiceRepository) *Service {
	return &Service{
		db: database,
	}
}
