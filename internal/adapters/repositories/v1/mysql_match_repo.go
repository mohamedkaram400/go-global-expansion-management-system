package repositories

import (
	"context"
	"fmt"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MatchRepo struct {
	DB *gorm.DB
}

func NewMatchRepo(db *gorm.DB) *MatchRepo {
	return &MatchRepo{DB: db}
}

func (r *MatchRepo) GetVendorsForProject(ctx context.Context, project *entities.Project) ([]entities.Vendor, error) {
	var vendors []entities.Vendor

	// Simplified example: filter by country first
	if err := r.DB.WithContext(ctx).
		Where("JSON_CONTAINS(countries_supported, ?)", fmt.Sprintf(`"%s"`, project.Country)).
		Find(&vendors).Error; err != nil {
		return nil, err
	}

	// fmt.Printf("✅ Project retruned: %+v\n", vendors) // debug

	return vendors, nil
}

func (r *MatchRepo) UpsertMatch(ctx context.Context, match *entities.Match) error {
	return r.DB.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "project_id"}, {Name: "vendor_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"score"}),
	}).Create(match).Error
}

// SELECT * FROM vendors V INNER JOIN projects P WHERE P.country in V.countries_supported AND P.services_needed in V.services_offered;
