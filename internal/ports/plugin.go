package ports

import "github.com/shashank-priyadarshi/upgraded-disco/constants"

type Trigger interface {
	Run(...interface{}) (constants.HTTPStatusCode, error)
}

type Plugin interface {
	List() []interface{}
	Info(string) interface{}
	Trigger(string, ...interface{}) error              // Trigger the plugin
	GetLastTrigger(string, ...interface{}) interface{} // Get the plugin trigger based on name,  count and failed/successful executions
}

type Builder interface {
	Install(string, []byte) error // Detect language and build the corresponding folder given the plugin name
	Upgrade(string, []byte) error // Download latest code and rebuild the binary
	Uninstall(string) error       // Remove the built binary
}
