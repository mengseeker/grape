package bl

import "grape/grape/models"

func CreateNs(name, code, note string, clusterID int64) models.Namespace {
	record := models.Namespace{
		Name:      name,
		Code:      code,
		ClusterID: clusterID,
		Note:      note,
	}
	PanicErr(db().Create(&record).Error)
	return record
}
