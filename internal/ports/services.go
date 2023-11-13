package ports

type Services interface {
	DataService() DataOps
	AccountManagementService() AccountOps
	PluginService() PluginOps
	ScheduleService() ScheduleOps
	GraphQLService() GraphQLOps
}

type DataOps interface {
	GetGraphData() (interface{}, error)
	GetGitHubData() (interface{}, error)
}

type AccountOps interface {
	RegisterUser(interface{}) error
	Login(interface{}) (interface{}, error)
	ResetPassword(interface{}) error
	DeleteUser(interface{}) error
}

type PluginOps interface {
	List() (interface{}, error)
	Update(interface{}) error
	Install(interface{}) error
	Trigger(interface{}) error
}

type ScheduleOps interface {
	List(interface{}) (interface{}, error)
	Create(interface{}) (interface{}, error)
	Delete(interface{}) error
}

type GraphQLOps interface {
	GraphQL(interface{}) (interface{}, error)
}
