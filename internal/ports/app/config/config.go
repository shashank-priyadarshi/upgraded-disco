package config

import models "github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/core/domain"

type Config interface {
	Load() (models.Config, error)
	Refresh() (models.Config, error)
}
