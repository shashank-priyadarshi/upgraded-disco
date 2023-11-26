package schedule

type Service struct {
	db interface{}
}

// TODO
func (s Service) List(schedule interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Create(schedule interface{}) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Delete(schedule interface{}) error {
	//TODO implement me
	panic("implement me")
}

func NewApplication(database interface{}) *Service {
	return &Service{
		db: database,
	}
}
