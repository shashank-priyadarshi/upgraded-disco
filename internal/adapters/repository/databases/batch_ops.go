package databases

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
)

type BatchOps struct {
	ports.Database
}

func (b BatchOps) BatchCreate(data []interface{}) []ports.BatchOpsResult {
	var batchOpsResult []ports.BatchOpsResult
	for _, value := range data {
		res, err := b.Create(value)
		batchOpsResult = append(batchOpsResult, ports.BatchOpsResult{
			Output: res,
			Error:  err,
		})
	}
	return batchOpsResult
}

func (b BatchOps) BatchGet(data []interface{}) []ports.BatchOpsResult {
	var batchOpsResult []ports.BatchOpsResult
	for _, value := range data {
		res, err := b.Get(value)
		batchOpsResult = append(batchOpsResult, ports.BatchOpsResult{
			Output: res,
			Error:  err,
		})
	}
	return batchOpsResult
}

func (b BatchOps) BatchUpdate(data []interface{}) []ports.BatchOpsResult {
	var batchOpsResult []ports.BatchOpsResult
	for _, value := range data {
		res, err := b.Update(value)
		batchOpsResult = append(batchOpsResult, ports.BatchOpsResult{
			Output: res,
			Error:  err,
		})
	}
	return batchOpsResult
}

func (b BatchOps) BatchDelete(data []interface{}) []ports.BatchOpsResult {
	var batchOpsResult []ports.BatchOpsResult
	for _, value := range data {
		res, err := b.Delete(value)
		batchOpsResult = append(batchOpsResult, ports.BatchOpsResult{
			Output: res,
			Error:  err,
		})
	}
	return batchOpsResult
}

func (b BatchOps) BatchExists(data []interface{}) []ports.BatchOpsResult {
	var batchOpsResult []ports.BatchOpsResult
	for _, value := range data {
		exists := b.Exists(value)
		batchOpsResult = append(batchOpsResult, ports.BatchOpsResult{
			Exists: exists,
		})
	}
	return batchOpsResult
}
