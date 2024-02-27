package accountmanagement

import (
	"context"
	"github.com/shashank-priyadarshi/upgraded-disco/models"
	logger "github.com/shashank-priyadarshi/utilities/logger/ports"
)

type Service struct {
	db models.Repository
}

// TODO
// Send account confirmation email
func (s Service) Register(userData interface{}) error {
	user := userData.(models.RegisterUser)
	//if s.db.MariaDB.Exists(user.Username) || s.db.MariaDB.Exists(user.Email) {
	//}
	_, err := s.db.MariaDB.Create(context.TODO(), user)
	if err != nil {
	}
	return nil
}

// Can be email or username, if empty, needs to be API or JWT token
func (s Service) Login(interface{}) (interface{}, error) {
	//TODO implement me
	return nil, nil
}

// Send password reset email
func (s Service) Update(interface{}) error {
	//TODO implement me
	return nil
}

// Send account deletion confirmation email
func (s Service) Delete(interface{}) error {
	//TODO implement me
	return nil
}

func NewApplication(log logger.Logger, database interface{}) *Service {

	log.Info("Initialising account management service")
	return &Service{
		db: database.(models.Repository),
	}
}
