package config

import (
	"fmt"
	"github.com/shashank-priyadarshi/upgraded-disco/constants"
	models "github.com/shashank-priyadarshi/upgraded-disco/models"
	"github.com/spf13/viper"
)

type Source struct {
	ConfigSource, ConfigPath string
}

// Load based on config Source and ConfigPath, use switch case to Load config from env var or from yaml, toml, json using viper
func (s *Source) Load() (config models.Config, err error) {
	if err = s.getConfigFile(); err != nil {
		return config, fmt.Errorf("error looking up for file %s at location %s: %v", s.ConfigSource, s.ConfigPath, err)
	}
	if config, err = setConfig(); err != nil {
		return config, fmt.Errorf("error reading config from file %s at location %s: %v", s.ConfigSource, s.ConfigPath, err)
	}
	return config, nil
}

// Refresh based on config Source and ConfigPath
func (s *Source) Refresh() (config models.Config, err error) {
	if config, err = setConfig(); err != nil {
		return config, fmt.Errorf("error refreshing config from file %s at location %s: %v", s.ConfigSource, s.ConfigPath, err)
	}
	return config, nil
}

func (s *Source) getConfigFile() error {
	viper.SetConfigFile(fmt.Sprintf("%s/%s", s.ConfigPath, s.ConfigSource))
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func setConfig() (config models.Config, err error) {
	dbConfig := make(map[string]models.DBConfig)
	dbConfig[constants.DB_REDIS] = models.DBConfig{
		Username:           viper.Get("database.redis.username").(string),
		Password:           viper.Get("database.redis.password").(string),
		Host:               viper.Get("database.redis.host").(string),
		Database:           viper.Get("database.redis.database").(string),
		MaxIdleConnections: viper.Get("database.redis.idle").(int),
		MaxOpenConnections: viper.Get("database.redis.idle").(int),
	}
	dbConfig[constants.DB_MARIADB] = models.DBConfig{
		Username:           viper.Get("database.mariadb.username").(string),
		Password:           viper.Get("database.mariadb.password").(string),
		Host:               viper.Get("database.mariadb.host").(string),
		Database:           viper.Get("database.mariadb.database").(string),
		MaxIdleConnections: viper.Get("database.mariadb.idle").(int),
		MaxOpenConnections: viper.Get("database.mariadb.open").(int),
	}
	dbConfig[constants.DB_MONGODB] = models.DBConfig{
		Username: viper.Get("database.mongo.username").(string),
		Password: viper.Get("database.mongo.password").(string),
		Host:     viper.Get("database.mongo.host").(string),
	}
	config.DBConfig = dbConfig
	config.ServerConfig.Port = viper.Get("server.port").(string)
	return
}
