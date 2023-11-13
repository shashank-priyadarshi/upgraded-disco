package data

import "github.com/shashank-priyadarshi/upgraded-disco/models"

type Service struct {
	db models.Repository
}

// TODO
func (s Service) GetGraphData() {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetGitHubData() {
	//TODO implement me
	panic("implement me")
}

func NewApplication(database interface{}) *Service {
	return &Service{
		db: database.(models.Repository),
	}
}
