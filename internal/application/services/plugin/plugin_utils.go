package plugin

import (
	"github.com/shashank-priyadarshi/upgraded-disco/models"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
)

type Plugin struct {
	Databases *models.Repository
	Config    *models.PluginsConfig
	Log       logger.Logger
	//	Initialise worker pool and create new topics: JOBS & PLUGINS
}

// TODO
func initPlugin(databases *models.Repository, config *models.PluginsConfig, log logger.Logger) (plugin *Plugin) {
	/* TODO future:
	   Get OS: set work directory, initialise directory for built binaries
	   Read supported languages, read commands, check if dependencies available
	   asynchronously iterate over plugin source directory and build all plugins
	*/

	//	TODO: Fetch list of plugins using config at initialisation
	//	TODO: Trigger all plugins through the Trigger interface
	//	TODO: Use a global worker pool to run max 5 plugin triggers at one time
	plugin = &Plugin{
		Databases: databases,
		Config:    config,
		Log:       log,
	}

	return
}

// TODO
func (p *Plugin) List() []interface{} { return nil }

func (p *Plugin) Info(plugin string) interface{} { return nil }

func (p *Plugin) Trigger(plugin string, payload ...interface{}) error {
	// Get global worker pool, submit a plugin run job to worker pool
	// Each plugin must implement Trigger interface{}
	// If a plugin has been triggered recently, do not trigger for 8 hours
	return nil
}

func (p *Plugin) GetLastTrigger(plugin string, payload ...interface{}) interface{} {
	// Read last trigger from Redis
	// Return time by default
	// For other details, payload must be received
	return nil
}

func (p *Plugin) Install(plugin string, sourceCode []byte) error {
	// Get OS: set work directory
	// Copy plugin code
	// Detect language
	// Build plugin
	// Remove code
	return nil
}
func (p *Plugin) Upgrade(plugin string, sourceCode []byte) error {
	// Get OS: set work directory
	// Call Uninstall()
	// Copy plugin code
	// Detect language
	// Build plugin
	// Remove code
	return nil
}
func (p *Plugin) Uninstall(plugin string) error {
	// Get OS: remove plugin from work directory
	return nil
}

func getOS() (string, error)          { return "", nil }
func setupWorkspace() error           { return nil }
func cleanupWorkspace() error         { return nil }
func detectLanguage() (string, error) { return "", nil }
func buildPlugin() error              { return nil }
