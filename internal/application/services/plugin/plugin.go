package plugin

import (
	"fmt"
	"github.com/shashank-priyadarshi/upgraded-disco/models"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/pubsub"
)

type Service struct {
	plugin *Plugin
	log    logger.Logger
	db     interface{}
	broker pubsub.PubSub
}

func NewApplication(log logger.Logger, database interface{}, broker pubsub.PubSub, config *models.PluginsConfig) *Service {

	log.Infof("Initialising plugin service")
	db := database.(models.Repository)
	plugin := initPlugin(&db, config, log)

	/*
		1. Initialise PubSub with topics: PLUGINS & JOBS(topic initialisation unnecessary as empty topics will be discarded and if a topic does not exist it will be created)
		2. Initialise worker pool with worker limit read from config
		3. Publish jobs to JOBS topic
		4. Subscribe to PLUGINS topic for update
		5. Worker pool subscribes to JOBS topic and publishes to PLUGINS topic
	*/
	return &Service{
		log:    log,
		db:     database,
		broker: broker,
		plugin: plugin,
	}
}

func (s *Service) Install(i interface{}) error {
	//TODO implement me
	return nil
}

func (s *Service) List() (interface{}, error) {
	//TODO implement me
	return nil, nil
}

func (s *Service) Info(s3 string, s2 string, i ...interface{}) (interface{}, error) {
	//TODO implement me
	return nil, nil
}

func (s *Service) Upgrade(i interface{}) error {
	//TODO implement me
	return nil
}

func (s *Service) Trigger(plugin string) error {
	if err := s.broker.Publish("JOBS", plugin); err != nil {
		s.log.Errorf("Error publishing plugin trigger request to PLUGINS topic for plugin: ", plugin)
		return fmt.Errorf("Error triggering plugin")
	}
	subscriberID := s.broker.Subscribe("PLUGINS", func(message interface{}) {
		s.log.Infof("Result of plugin execution received on channel PLUGINS: %s", message)
	})
	s.log.Infof("Subcriber ID for this job: %s", subscriberID)
	return nil
}

func (s *Service) Uninstall(s2 string) error {
	//TODO implement me
	return nil
}
