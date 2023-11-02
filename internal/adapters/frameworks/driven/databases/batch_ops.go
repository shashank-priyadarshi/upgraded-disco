package databases

import "github.com/shashank-priyadarshi/upgraded-disco/internal/ports/frameworks/driven/database"

type BatchOps struct {
	database.Database
}

func (b BatchOps) BatchCreate(data []interface{}) []database.BatchOpsResult {
	var batchOpsResult []database.BatchOpsResult
	for _, value := range data {
		res, err := b.Create(value)
		batchOpsResult = append(batchOpsResult, database.BatchOpsResult{
			Output: res,
			Error:  err,
		})
	}
	return batchOpsResult
}

func (b BatchOps) BatchGet(data []interface{}) []database.BatchOpsResult {
	var batchOpsResult []database.BatchOpsResult
	for _, value := range data {
		res, err := b.Get(value)
		batchOpsResult = append(batchOpsResult, database.BatchOpsResult{
			Output: res,
			Error:  err,
		})
	}
	return batchOpsResult
}

func (b BatchOps) BatchUpdate(data []interface{}) []database.BatchOpsResult {
	var batchOpsResult []database.BatchOpsResult
	for _, value := range data {
		res, err := b.Update(value)
		batchOpsResult = append(batchOpsResult, database.BatchOpsResult{
			Output: res,
			Error:  err,
		})
	}
	return batchOpsResult
}

func (b BatchOps) BatchDelete(data []interface{}) []database.BatchOpsResult {
	var batchOpsResult []database.BatchOpsResult
	for _, value := range data {
		res, err := b.Delete(value)
		batchOpsResult = append(batchOpsResult, database.BatchOpsResult{
			Output: res,
			Error:  err,
		})
	}
	return batchOpsResult
}
