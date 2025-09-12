package requests

import "github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"

type ProjectRequest struct {
	ClientId        uint      `json:"client_id"		binding:"required,gt=0"`
	ServicesNeeded  []string  `json:"service_needed"	binding:"required"`
	Country    		string    `json:"country"	binding:"required"`
	Budget          float64   `json:"budget"	binding:"required,gte=1"`
    Status         entities.ProjectStatus `json:"status"	binding:"omitempty,oneof=active completed cancelled"`
}
