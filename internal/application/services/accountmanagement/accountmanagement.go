package accountmanagement

type Service struct {
	db interface{}
}

// TODO
// Send account confirmation email
func (s Service) RegisterUser(interface{}) error {
	//TODO implement me
	panic("implement me")
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
		db: database,
	}
}
