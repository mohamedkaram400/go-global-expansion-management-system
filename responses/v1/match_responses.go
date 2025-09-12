package responses

import "github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"

type MatchResponse struct {
	ID              uint            `json:"id"`
	MatchingVenders *entities.Match `json:""`
}
