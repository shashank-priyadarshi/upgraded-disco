package main

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/plugins"
)

func main() {
	//	TODO: Add Strategy Pattern to add services as required and start server
	//	TODO: Structure init for all services
	// init server
	// accepts config, maria, redis, mongo and logger
}

func initPlugin(pluginConf *plugins.Plugin) {
	// Get OS: set work directory, initialise directory for built binaries
	// Read supported languages, read commands, check if dependencies available
	// asynchronously iterate over plugin source directory and build all plugins
}
