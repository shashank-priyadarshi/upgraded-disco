package plugins

import (
	models "github.com/shashank-priyadarshi/upgraded-disco/models"
	"go.uber.org/zap"
)

type Plugin struct {
	Databases models.Repository
	Log       zap.Logger
}

// TODO
func (p *Plugin) Trigger(plugin string, payload interface{}) error {
	// Get OS: search for plugin
	// Trigger plugin
	// Save trigger details to Redis
	return nil
}
func (p *Plugin) GetLastTrigger(plugin string, payload interface{}) interface{} {
	// Read last trigger from Redis
	// Return time by default
	// For other details, payload must be received
	return nil
}

func (p *Plugin) Install(plugin string) error {
	// Get OS: set work directory
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
func (p *Plugin) Upgrade(plugin string) error {
	// Get OS: set work directory
	// Call Uninstall()
	// Copy plugin code
	// Detect language
	// Build plugin
	// Remove code
	return nil
}

func getOS() (string, error)          { return "", nil }
func setupWorkspace() error           { return nil }
func cleanupWorkspace() error         { return nil }
func detectLanguage() (string, error) { return "", nil }
func buildPlugin() error              { return nil }
