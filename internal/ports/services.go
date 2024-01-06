package ports

type Services interface {
	DataService() DataOps
	AccountManagementService() AccountOps
	PluginService() PluginOps
	ScheduleService() ScheduleOps
	GraphQLService() GraphQLOps
}

type DataOps interface {
	Chess() (interface{}, error)
	GitHub() (interface{}, error)
}

type AccountOps interface {
	Register(interface{}) error
	Login(interface{}) (interface{}, error)
	Update(interface{}) error
	Delete(interface{}) error
}

type PluginOps interface {
	Install(interface{}) error
	List() (interface{}, error)
	Info(string, string, ...interface{}) (interface{}, error)
	Upgrade(interface{}) error
	Trigger(interface{}) error
	Uninstall(string) error
}

type ScheduleOps interface {
	List(interface{}) (interface{}, error)
	Create(interface{}) (interface{}, error)
	Update(interface{}) error
	Delete(interface{}) error
}

type GraphQLOps interface {
	GraphQL(interface{}) (interface{}, error)
}
