package schedule

type Service struct {
	db interface{}
}

func (s Service) List() {
	//TODO implement me
	panic("implement me")
}

func (s Service) Create() {
	//TODO implement me
	panic("implement me")
}

func (s Service) Delete() {
	//TODO implement me
	panic("implement me")
}

func NewApplication(database interface{}) *Service {
	return &Service{
		db: database,
	}
}
