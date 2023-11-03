package accountmanagement

import "github.com/shashank-priyadarshi/upgraded-disco/internal/ports/frameworks/driven/database"

type Service struct {
	db database.ServiceRepository
}

func NewApplication(database database.ServiceRepository) *Service {
	return &Service{
		db: database,
	}
}
