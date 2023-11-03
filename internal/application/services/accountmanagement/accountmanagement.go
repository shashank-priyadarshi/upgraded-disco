package accountmanagement

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
)

type Service struct {
	db ports.ServiceRepository
}

func (s Service) RegisterUser() {
	//TODO implement me
	panic("implement me")
}

func (s Service) Login() {
	//TODO implement me
	panic("implement me")
}

func (s Service) ResetPassword() {
	//TODO implement me
	panic("implement me")
}

func (s Service) DeleteUser() {
	//TODO implement me
	panic("implement me")
}

func NewApplication(database ports.ServiceRepository) *Service {
	return &Service{
		db: database,
	}
}
