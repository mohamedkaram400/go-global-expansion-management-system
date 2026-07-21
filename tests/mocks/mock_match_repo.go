package mocks

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
)

type MockMatchRepo struct {
	GetVendorsForProjectFunc   func(ctx context.Context, project *entities.Project) ([]*entities.Vendor, error)
	UpsertMatchFunc            func(ctx context.Context, match *entities.Match) error
	GetTopVendorsByCountryFunc func(ctx context.Context, days int) (map[string][]*entities.VendorAnalytics, error)
}

func (m *MockMatchRepo) GetVendorsForProject(ctx context.Context, project *entities.Project) ([]*entities.Vendor, error) {
	return m.GetVendorsForProjectFunc(ctx, project)
}

func (m *MockMatchRepo) UpsertMatch(ctx context.Context, match *entities.Match) error {
	return m.UpsertMatchFunc(ctx, match)
}

func (m *MockMatchRepo) GetTopVendorsByCountry(ctx context.Context, days int) (map[string][]*entities.VendorAnalytics, error) {
	return m.GetTopVendorsByCountryFunc(ctx, days)
}
