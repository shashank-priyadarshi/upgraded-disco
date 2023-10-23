package database

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
	BatchUpdate(fields, data []interface{}) []BatchOpsResult
	BatchDelete(data []interface{}) []BatchOpsResult
}
