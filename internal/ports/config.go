package ports

import models "github.com/shashank-priyadarshi/upgraded-disco/models"

type Config interface {
	Load() (models.Config, error)
	Refresh() (models.Config, error)
}
