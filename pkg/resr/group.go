package resr

import (
	"fmt"
	"grape/grape/models"
	"grape/pkg/share"
)

type Group struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Code        string   `json:"code"`
	ServiceCode string   `json:"service_code"`
	ClusterCode string   `json:"cluster_code"`
	DeployType  int      `json:"deploy_type"`
	Nodes       []string `json:"nodes"`
}

func NewGroup(source *models.Group) *Group {
	return &Group{
		ID:          source.ID,
		Name:        source.Name,
		ServiceCode: source.Service().Code,
		DeployType:  source.DeployType,
		Nodes:       source.NodeIPs(),
	}
}

func (r *Group) Marshal() []byte {
	bs, _ := Marshal(r)
	return bs
}

func (r *Group) Key(clusterCode string) string {
	return share.ResourceGroupPrefix(clusterCode) + fmt.Sprintf("%s/%s", r.ServiceCode, r.Code)
}
