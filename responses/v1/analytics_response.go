package responses

import "github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"

type VendorAnalyticsResponse struct {
	Country           string                       `json:"country"`
	TopVendors        []*entities.VendorAnalytics `json:"top_vendors"`
	ResearchDocsCount int                          `json:"research_docs_count"`
}
