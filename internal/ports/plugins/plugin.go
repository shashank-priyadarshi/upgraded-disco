package plugin


type Plugin interface{
	Install(plugin string) error
	Uninstall(plugin string) error
	Upgrade(plugin string) error
	Trigger(plugin string, payload interface{}) error
	GetLastTrigger() interface{}
}

type PluginWrapper interface{}