package graphql

type Service struct {
	db interface{}
}

// TODO
func (s Service) GraphQL() {
	//TODO implement me
	panic("implement me")
}

func NewApplication(database interface{}) *Service {
	return &Service{
		db: database,
	}
}
