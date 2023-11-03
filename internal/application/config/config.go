package config

import models "github.com/shashank-priyadarshi/upgraded-disco/models"

type Source struct {
	Source, SourcePath string
}

// Load based on config Source and SourcePath, use switch case to Load config from env var or from yaml, toml, json using viper
func (a *Source) Load() (models.Config, error) {
	return models.Config{}, nil
}

// Refresh based on config Source and SourcePath
func (a *Source) Refresh() (models.Config, error) {
	return models.Config{}, nil
}
