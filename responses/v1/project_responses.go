package responses

import (
	"encoding/json"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
)

type ProjectResponse struct {
	ID             int      `json:"id"`
	ClientId       int      `json:"client_id"`
	ServicesNeeded []string `json:"service_needed"`
	Country        string   `json:"country"`
	Budget         float64  `json:"budget"`
	Status         string   `json:"status"`
}

func FormatProject(project *entities.Project) ProjectResponse {
	var services []string
	_ = json.Unmarshal(project.ServicesNeeded, &services)

	return ProjectResponse{
		ID:             project.ID,
		ClientId:       project.ClientID,
		ServicesNeeded: services,
		Country:        project.Country,
		Budget:         project.Budget,
		Status:         project.Status,
	}
}
func FormatProjects(projects []*entities.Project) []ProjectResponse {
	responses := make([]ProjectResponse, 0, len(projects))
	for _, v := range projects {
		responses = append(responses, FormatProject(v))
	}
	return responses
}
