package plugin

type Service struct {
	db interface{}
}

// TODO
func (s Service) List() (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Update(interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) Install(interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) Trigger(interface{}) error {
	//TODO implement me
	panic("implement me")
}

func NewApplication(database interface{}) *Service {
	return &Service{
		db: database,
	}
}
