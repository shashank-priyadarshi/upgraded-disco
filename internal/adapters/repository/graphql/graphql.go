package graphql

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/frameworks/driven/databases"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
)

type graphQLRepo struct {
	ports.Database
}

func NewGraphQLRepo(log logger.Logger, config interface{}) ports.GraphQLRepo {
	db, err := databases.NewDatabase("", log, config)
	if err != nil {
		log.Errorf("failed to initialise schedule repo: %v", err)
		return nil
	}
	return &graphQLRepo{
		db,
	}
}
