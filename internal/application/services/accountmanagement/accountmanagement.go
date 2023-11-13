package accountmanagement

type Service struct {
	db interface{}
}

// TODO
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

func NewApplication(database interface{}) *Service {
	return &Service{
		db: database,
	}
}
