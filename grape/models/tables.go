package models

import "gorm.io/gorm"

type TableName = string

const (
	TableNameCluster   TableName = "clusters"
	TableNameEtcdLink  TableName = "etcd_links"
	TableNameGroup     TableName = "groups"
	TableNameNamespace TableName = "namespaces"
	TableNameNode      TableName = "nodes"
	TableNamePolicy    TableName = "policies"
	TableNameService   TableName = "services"
	TableNameUser      TableName = "users"
)

func ModelCluster() *gorm.DB {
	return db.Table(TableNameCluster)
}

func ModelEtcdLink() *gorm.DB {
	return db.Table(TableNameEtcdLink)
}

func ModelGroup() *gorm.DB {
	return db.Table(TableNameGroup)
}

func ModelNs() *gorm.DB {
	return db.Table(TableNameNamespace)
}

func ModelNode() *gorm.DB {
	return db.Table(TableNameNode)
}

func ModelPolicy() *gorm.DB {
	return db.Table(TableNamePolicy)
}

func ModelService() *gorm.DB {
	return db.Table(TableNameService)
}

func ModelUser() *gorm.DB {
	return db.Table(TableNameUser)
}
