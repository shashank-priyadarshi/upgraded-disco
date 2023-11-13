package data

import "github.com/shashank-priyadarshi/upgraded-disco/models"

type Service struct {
	db models.Repository
}

// TODO
func (s Service) GetGraphData() (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetGitHubData() (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func NewApplication(database interface{}) *Service {
	return &Service{
		db: database.(models.Repository),
	}
}
