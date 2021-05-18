package models

import "time"

// Group is an object representing the database table.
type Group struct {
	// record
	ID          int64  `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"index;unique;not null;" json:"name"`
	Code        string `gorm:"index;unique;not null;" json:"code"`
	Version     string `gorm:"not null;default:'';" json:"version"`
	Lang        string `gorm:"not null;default:'';" json:"lang"`
	NamespaceID int    `gorm:"index;not null;" json:"namespace_id"`
	ServiceID   int    `gorm:"index;not null;" json:"service_id"`
	ClusterID   int    `gorm:"index;not null;" json:"cluster_id"`
	DeployType  int    `gorm:"not null;" json:"deploy_type"`
	Note        string `gorm:"not null;default:'';" json:"note"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	F_Service Service `gorm:"foreignKey:ServiceID" json:"-"`
	F_Cluster Cluster `gorm:"foreignKey:ClusterID" json:"-"`
}

func (r *Group) Cluster() *Cluster {
	clu := Cluster{}
	err := db.First(&clu, r.ClusterID).Error
	if err != nil {
		panic(err)
	}
	return &clu
}

func (r *Group) Service() *Service {
	srv := Service{}
	err := db.First(&srv, r.ServiceID).Error
	if err != nil {
		panic(err)
	}
	return &srv
}

func (r *Group) NodeIPs() []string {
	var nodes []Node
	err := db.Model(&nodes).Where("group_id = ?", r.ID).Find(&nodes)
	if err != nil {
		panic(err)
	}
	ips := []string{}
	for _, n := range nodes {
		ips = append(ips, n.IP)
	}
	return ips
}
