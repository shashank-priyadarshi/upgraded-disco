package ports

type BatchOpsResult struct {
	Output interface{}
	Error  error
	Exists bool
}

type Database interface {
	Create(data interface{}) (interface{}, error)
	Get(data interface{}) (interface{}, error)
	Update(data interface{}) (interface{}, error)
	Delete(data interface{}) (interface{}, error)
	Exists(data interface{}) bool
}

type BatchOps interface {
	BatchCreate(data []interface{}) []BatchOpsResult
	BatchGet(data []interface{}) []BatchOpsResult
	BatchUpdate(data []interface{}) []BatchOpsResult
	BatchDelete(data []interface{}) []BatchOpsResult
	BatchExists(data []interface{}) []BatchOpsResult
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
