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

func (r *MatchRepo) GetVendorsForProject(ctx context.Context, project *entities.Project) ([]*entities.Vendor, error) {
	var vendors []*entities.Vendor

	fmt.Println(project.Country)
	if r.DB.Dialector.Name() == "sqlite" {
		// Simplified example: filter by country first
		if err := r.DB.WithContext(ctx).
			Where("EXISTS (SELECT 1 FROM json_each(countries_supported) WHERE value = ?)", project.Country).
			Find(&vendors).Error; err != nil {
			return nil, err
		}
	} else {
		if err := r.DB.WithContext(ctx).
			Where("JSON_CONTAINS(countries_supported, JSON_QUOTE(?))", project.Country).
			Find(&vendors).Error; err != nil {
			return nil, err
		}
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
func (r *MatchRepo) GetTopVendorsByCountry(ctx context.Context, days int) (map[string][]*entities.VendorAnalytics, error) {
	type Result struct {
		Country       string
		VendorID      int
		VendorName    string
		AvgMatchScore float64
	}

	var results []Result

	var query string
	if r.DB.Dialector.Name() == "sqlite" {
		query = `
			SELECT country, vendor_id, vendor_name, avg_match_score FROM (
				SELECT 
					p.country,
					v.id AS vendor_id,
					v.name AS vendor_name,
					AVG(m.score) AS avg_match_score,
					ROW_NUMBER() OVER (PARTITION BY p.country ORDER BY AVG(m.score) DESC) AS rank
				FROM matches m
				INNER JOIN projects p ON m.project_id = p.id
				INNER JOIN vendors v ON m.vendor_id = v.id
				WHERE m.created_at >= datetime('now', '-' || ? || ' day')
				GROUP BY p.country, v.id, v.name
			) ranked
			WHERE rank <= 3;
		`
	} else {
		// MySQL version
		query = `
			SELECT *
			FROM (
				SELECT
					t.*,
					ROW_NUMBER() OVER (
						PARTITION BY country
						ORDER BY avg_match_score DESC
					) AS rn
				FROM (
					SELECT
						p.country,
						v.id AS vendor_id,
						v.name AS vendor_name,
						AVG(m.score) AS avg_match_score
					FROM matches m
					JOIN projects p ON m.project_id = p.id
					JOIN vendors v ON m.vendor_id = v.id
					WHERE m.created_at >= DATE_SUB(NOW(), INTERVAL ? DAY)
					GROUP BY p.country, v.id, v.name
				) t
			) ranked
			WHERE rn <= 3;
		`
	}

	if err := r.DB.WithContext(ctx).Raw(query, days).Scan(&results).Error; err != nil {
		return nil, err
	}

	// Convert results into map[country][]Vendor
	vendorMap := make(map[string][]*entities.VendorAnalytics)
	for _, res := range results {
		vendorMap[res.Country] = append(vendorMap[res.Country], &entities.VendorAnalytics{
			ID:            res.VendorID,
			Name:          res.VendorName,
			AvgMatchScore: res.AvgMatchScore,
		})
	}

	return vendorMap, nil
}

// func (r *MatchRepo) GetTopVendorsByCountry(ctx context.Context, days int) (map[string][]*entities.VendorAnalytics, error) {
// 	type Result struct {
// 		Country       string
// 		VendorID      int
// 		VendorName    string
// 		AvgMatchScore float64
// 	}

// 	var results []Result

// 	var query string
// 	if r.DB.Dialector.Name() == "sqlite" {
// 		query = `
// 			SELECT country, vendor_id, vendor_name, avg_match_score FROM (
// 				SELECT
// 					p.country,
// 					v.id AS vendor_id,
// 					v.name AS vendor_name,
// 					AVG(m.score) AS avg_match_score,
// 					ROW_NUMBER() OVER (PARTITION BY p.country ORDER BY AVG(m.score) DESC) AS rank
// 				FROM matches m
// 				INNER JOIN projects p ON m.project_id = p.id
// 				INNER JOIN vendors v ON m.vendor_id = v.id
// 				WHERE m.created_at >= datetime('now', '-' || ? || ' day')
// 				GROUP BY p.country, v.id, v.name
// 			) ranked
// 			WHERE rank <= 3;
// 		`
// 	} else {
// 		// MySQL version
// 		query = `
// 			SELECT country, vendor_id, vendor_name, avg_match_score FROM (
// 				SELECT
// 					p.country,
// 					v.id AS vendor_id,
// 					v.name AS vendor_name,
// 					AVG(m.score) AS avg_match_score,
// 					ROW_NUMBER() OVER (PARTITION BY p.country ORDER BY AVG(m.score) DESC) AS rank
// 				FROM matches m
// 				INNER JOIN projects p ON m.project_id = p.id
// 				INNER JOIN vendors v ON m.vendor_id = v.id
// 				WHERE m.created_at >= DATE_SUB(NOW(), INTERVAL ? DAY)
// 				GROUP BY p.country, v.id, v.name
// 			) ranked
// 			WHERE rank <= 3;
// 		`
// 	}

// 	if err := r.DB.WithContext(ctx).Raw(query, days).Scan(&results).Error; err != nil {
// 		return nil, err
// 	}

// 	// Convert results into map[country][]Vendor
// 	vendorMap := make(map[string][]*entities.VendorAnalytics)
// 	for _, res := range results {
// 		vendorMap[res.Country] = append(vendorMap[res.Country], &entities.VendorAnalytics{
// 			ID:            res.VendorID,
// 			Name:          res.VendorName,
// 			AvgMatchScore: res.AvgMatchScore,
// 		})
// 	}

// 	return vendorMap, nil
// }

// select AVG(M.score) from match M where M.created_at >= NOW() - INTERVAL 30 DAY group by country, vendor_id limit 3;
