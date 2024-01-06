package models

type PluginsConfig struct {
	Endpoint, Token string
}

type Metadata struct {
	Name, ID, Installed, Updated, Language string
	Tags                                   []string
}

type Trigger struct {
	Latest        string
	Total, Failed int
	Failures      []Failed
}

type Failed struct {
	TimeStamp  string
	Duration   int
	Parameters []interface{}
	Error      string
}
