package resr

import (
	"grape/grape/models"
	"grape/pkg/share"
)

type Service struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Port     int    `json:"port"`
	Protocol int    `json:"protocol"`
	External int    `json:"external"`
}

func NewService(source *models.Service) *Service {
	return &Service{
		ID:       source.ID,
		Name:     source.Name,
		Port:     source.Port,
		Protocol: source.Protocol,
		External: source.External,
	}
}

func (r *Service) Marshal() []byte {
	bs, _ := Marshal(r)
	return bs
}

func (r *Service) Key(clusterCode string) string {
	return share.ResourceServicePrefix(clusterCode) + r.Code
}
