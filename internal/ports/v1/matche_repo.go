package ports

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
)

type MatchRepository interface {
	GetVendorsForProject(ctx context.Context, project *entities.Project) ([]entities.Vendor, error)
	UpsertMatch(ctx context.Context, match *entities.Match) error
}
 