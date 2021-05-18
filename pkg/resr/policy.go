package resr

import (
	"fmt"
	"grape/grape/models"
	"grape/pkg/share"
)

type Policy struct {
	ID          int64                  `json:"id"`
	Category    int                    `json:"category"`
	Code        string                 `json:"code"`
	ServiceCode string                 `json:"service_code"`
	Active      int                    `json:"active"`
	Options     map[string]interface{} `json:"options"`
}

func NewPolicy(svc *Service, source *models.Policy) *Policy {
	return &Policy{
		ID:          source.ID,
		Code:        source.Code,
		Category:    source.Category,
		ServiceCode: svc.Code,
		Options:     source.OptionsMap(),
	}
}

func (r *Policy) Marshal() []byte {
	bs, _ := Marshal(r)
	return bs
}

func (r *Policy) Key(clusterCode string) string {
	return share.ResourcePolicyPrefix(clusterCode) + fmt.Sprintf("%s/%s", r.ServiceCode, r.Code)
}
