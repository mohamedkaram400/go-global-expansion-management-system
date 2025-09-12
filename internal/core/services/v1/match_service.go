package services

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/ports/v1"
)

type MatchService struct {
	Repo ports.MatchRepository
}

func NewMatchService(repo ports.MatchRepository) *MatchService {
	return &MatchService{Repo: repo}
}

func (svc *MatchService) Rebuild(ctx context.Context, project *entities.Match) (*entities.Match, error) {

	vendors, err := svc.Repo.GetVendorsForProject(ctx, project)
	if err != nil {
		return nil, err
	}

	// Calculate Score formula: services_overlap * 2 + rating + SLA_weight
	Matchs := make([]entities.Match, 0, len(vendors))

	for _, v := range vendors {
		overlap := countOverlap(project.ServicesNeeded, v.ServicesOffered)
		if overlap == 0 {
			continue
		}

		score := float64(overlap*2) + v.Rating + float64(v.ResponseSlaHours)
		match := entities.Match{
			ProjectID: project.ID,
			VendorID:  v.ID,
			Score:     score,
		}

		if err := svc.Repo.UpsertMatch(ctx, &match); err != nil {
			return nil, err
		}

		Matchs = append(Matchs, match)
	}
}

func countOverlap(projectServices, vendorServices []string) int {
	set := make(map[string]struct{})

	for _, s := range vendorServices {
		set[s] = struct{}{}
	}

	count := 0
	for _, ps := range projectServices {
		if _, ok := set[ps]; ok {
			count++
		}
	}
	return count
}
