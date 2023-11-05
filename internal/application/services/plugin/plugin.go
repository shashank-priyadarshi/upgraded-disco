package plugin

type Service struct {
	db interface{}
}

func (s Service) List() {
	//TODO implement me
	panic("implement me")
}

func (s Service) Update() {
	//TODO implement me
	panic("implement me")
}

func (s Service) Install() {
	//TODO implement me
	panic("implement me")
}

func (s Service) Trigger() {
	//TODO implement me
	panic("implement me")
}

func NewApplication(database interface{}) *Service {
	return &Service{
		db: database,
	}
}
