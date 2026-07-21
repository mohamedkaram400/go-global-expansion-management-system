package services

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/ports/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses/v1"
)

type AnalyticsService struct {
	MatchRepo            ports.MatchRepository
	ProjectRepo          ports.ProjectRepository
	ResearchDocumentRepo ports.ResearchDocumentRepository
}

func NewAnalyticsService(matchRepo ports.MatchRepository, researchDocumentRepo ports.ResearchDocumentRepository, projectRepo ports.ProjectRepository) *AnalyticsService {
	return &AnalyticsService{MatchRepo: matchRepo, ResearchDocumentRepo: researchDocumentRepo, ProjectRepo: projectRepo}
}

func (svc *AnalyticsService) GenerateAnalytics(ctx context.Context) ([]responses.VendorAnalyticsResponse, error) {
	// Step 1: Get vendors from MySQL
	vendorData, err := svc.MatchRepo.GetTopVendorsByCountry(ctx, 30)
	if err != nil {
		return nil, err
	}

	// Step 2: Get research docs grouped by project_id from Mongo
	researchCounts, err := svc.ResearchDocumentRepo.CountResearchDocsByProject(ctx)
	if err != nil {
		return nil, err
	}

	// Step 3: Map projects → countries
	projectCountries, err := svc.ProjectRepo.GetProjectCountries(ctx)
	if err != nil {
		return nil, err
	}

	// Step 4: Aggregate research docs per country
	researchByCountry := make(map[string]int)
	for projectID, count := range researchCounts {
		if country, ok := projectCountries[projectID]; ok {
			researchByCountry[country] += count
		}
	}

	// Step 5: Merge vendors + research docs
	result := []responses.VendorAnalyticsResponse{}

	for country, vendors := range vendorData {
		res := responses.VendorAnalyticsResponse{
			Country:    country,
			TopVendors: vendors,
		}

		if count, ok := researchByCountry[country]; ok {
			res.ResearchDocsCount = count
		}

		result = append(result, res)
	}

	return result, nil
}
