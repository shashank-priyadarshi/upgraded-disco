package accountmanagement

import "github.com/shashank-priyadarshi/upgraded-disco/models"

type Service struct {
	db models.Repository
}

// TODO
// Send account confirmation email
func (s Service) RegisterUser(userData interface{}) error {
	user := userData.(models.RegisterUser)
	if s.db.MariaDB.Exists(user.Username) || s.db.MariaDB.Exists(user.Email) {
	}
	_, err := s.db.MariaDB.Create(user)
	if err != nil {
	}
	return nil
}

// Can be email or username, if empty, needs to be API or JWT token
func (s Service) Login(interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

// Send password reset email
func (s Service) ResetPassword(interface{}) error {
	//TODO implement me
	panic("implement me")
}

// Send account deletion confirmation email
func (s Service) DeleteUser(interface{}) error {
	//TODO implement me
	panic("implement me")
}

func NewApplication(database interface{}) *Service {
	return &Service{
		db: database.(models.Repository),
	}
}
