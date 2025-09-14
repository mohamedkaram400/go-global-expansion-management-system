package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/ports/v1"
)

type MatchService struct {
	MatchRepo ports.MatchRepository
	ProjectService *ProjectService
}

func NewMatchService(repo ports.MatchRepository, projectService *ProjectService) *MatchService {
	return &MatchService{MatchRepo: repo, ProjectService: projectService}
}

func (svc *MatchService) Rebuild(ctx context.Context, projectID int) ([]entities.Match, error) {

	// Load project from DB
	project, err := svc.ProjectService.FindProjectByID(ctx, projectID)
	if err != nil {
		return nil, err
	}
 
	fmt.Println(project)

	// Parse JSON []byte → []string
	var projectServices []string
	if err := json.Unmarshal(project.ServicesNeeded, &projectServices); err != nil {
		return nil, fmt.Errorf("failed to parse services_needed: %w", err)
	}


	// Get all vendors that apply rules
	vendors, err := svc.MatchRepo.GetVendorsForProject(ctx, project)
	if err != nil {
		return nil, err
	}
 
	// Loop over vendors and Calculate Score
	Matchs := make([]entities.Match, 0, len(vendors))

	// Count number of service overlap in project and offered by vendor
	for _, v := range vendors {
		var vendorServices []string
		if err := json.Unmarshal(v.ServicesOffered, &vendorServices); err != nil {
			return nil, fmt.Errorf("failed to parse vendor services_offered: %w", err)
		}

		// Count how many service overlap with project and vendor
		overlap := countOverlap(projectServices, vendorServices)

		if overlap == 0 {
			continue
		}

		// Calculate Score formula: services_overlap * 2 + rating + SLA_weight
		score := float64(overlap*2) + v.Rating + float64(v.ResponseSlaHours)
		match := entities.Match{
			ProjectId: project.ID,
			VendorId:  v.ID,
			Score:     score,
		}

		// Save match result with socre in matches table 
		if err := svc.MatchRepo.UpsertMatch(ctx, &match); err != nil {
			return nil, err
		}

		// Append all match result to matchs
		Matchs = append(Matchs, match)
	}

	return Matchs, nil
}

func countOverlap(projectServices, vendorServices []string) int {
	set := make(map[string]struct{})

	for _, vs := range vendorServices {
		set[vs] = struct{}{}
	}

	count := 0
	for _, ps := range projectServices {
		if _, ok := set[ps]; ok {
			count++
		}
	}
	return count
}
