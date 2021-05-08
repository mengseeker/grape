package bl

import (
	"grape/grape/models"
)

func CreateEtcd(name, address, note string) models.EtcdLink {
	record := models.EtcdLink{
		Name:    name,
		Address: address,
		Note:    note,
	}
	db().Create(&record)
	return record
}
