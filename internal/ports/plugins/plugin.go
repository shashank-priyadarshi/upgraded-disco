package plugin

type PluginWrapper interface{}

type Plugin interface{
	Trigger()
	
}