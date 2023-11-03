package ports

type BatchOpsResult struct {
	Output interface{}
	Error  error
}

type Database interface {
	Create(data interface{}) (interface{}, error)
	Get(data interface{}) (interface{}, error)
	Update(data interface{}) (interface{}, error)
	Delete(data interface{}) (interface{}, error)
}

type BatchOps interface {
	BatchCreate(data []interface{}) []BatchOpsResult
	BatchGet(data []interface{}) []BatchOpsResult
	BatchUpdate(data []interface{}) []BatchOpsResult
	BatchDelete(data []interface{}) []BatchOpsResult
}

type ServiceRepository interface {
	GetDataServiceRepository() DataRepo
	GetAccountManagementRepository() AccountRepo
	GetPluginServiceRepository() PluginRepo
	GetScheduleServiceRepository() ScheduleRepo
	GetGraphQLServiceRepository() GraphQLRepo
}

type DataRepo interface {
	Database
}

type AccountRepo interface {
	Database
}

type PluginRepo interface {
	Database
}

type ScheduleRepo interface {
	Database
}

type GraphQLRepo interface {
	Database
}
