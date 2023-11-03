package services

type Services interface{
	DataService() DataOps
	AccountManagementService() AccountOps
	PluginService() PluginOps
	ScheduleService() ScheduleOps
	GraphQLService() GraphQLOps
}

type DataOps interface{
	GetGraphData()
	GetGitHubData()
}

type AccountOps interface{
	RegisterUser()
	Login()
	ResetPassword()
	DeleteUser()
}

type PluginOps interface{
	List()
	Update()
	Install()
	Trigger()
}

type ScheduleOps interface{
	List()
	Create()
	Delete()
}

type GraphQLOps interface{
	GraphQL()
}