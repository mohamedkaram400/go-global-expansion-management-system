package services

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/ports/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/responses/v1"
)

type AnalyticsService struct {
	MatchRepo ports.MatchRepository
	ResearchDocumentRepo ports.ResearchDocumentRepository
}

func NewAnalyticsService(matchRepo ports.MatchRepository, researchDocumentRepo ports.ResearchDocumentRepository) *AnalyticsService {
	return &AnalyticsService{MatchRepo: matchRepo, ResearchDocumentRepo: researchDocumentRepo}
}

func (svc *AnalyticsService) GenerateAnalytics(ctx context.Context) ([]responses.VendorAnalyticsResponse, error) {

	// Step 1: Get vendors from MySQL
	vendorData, err := svc.MatchRepo.GetTopVendorsByCountry(ctx, 30) 
	if err != nil {
		return nil, err
	}

	// Step 2: Get research documents count from MongoDB
	researchCounts, err := svc.ResearchDocumentRepo.CountResearchDocsByCountry(ctx)
	if err != nil {
		return nil, err
	}

	// Step 3: Merge results
	result := []responses.VendorAnalyticsResponse{}

	for country, vendors := range vendorData {
		res := responses.VendorAnalyticsResponse{
			Country:    country,
			TopVendors: vendors,
		}

		// Add MongoDB data (if exists)
		if count, ok := researchCounts[country]; ok {
			res.ResearchDocsCount = count
		}

		result = append(result, res)
	}

	return result, nil
}