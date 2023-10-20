package plugin


type Plugin interface{
	Install(plugin string) error // Build the corresponding folder given the plugin name
	Uninstall(plugin string) error // Remove the built binary
	Upgrade(plugin string) error // Download latest code and rebuild the binary
	Trigger(plugin string, payload interface{}) error // Trigger the plugin
	GetLastTrigger(plugin string, payload interface{}) interface{} // Get the plugin trigger based on name,  count and failed/successful executions
}

type PluginWrapper interface{}