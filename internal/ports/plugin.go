package ports

type Plugin interface {
	Trigger(plugin string, payload interface{}) error              // Trigger the plugin
	GetLastTrigger(plugin string, payload interface{}) interface{} // Get the plugin trigger based on name,  count and failed/successful executions
}

type Builder interface {
	Install(plugin string) error   // Detect language and build the corresponding folder given the plugin name
	Uninstall(plugin string) error // Remove the built binary
	Upgrade(plugin string) error   // Download latest code and rebuild the binary
}
