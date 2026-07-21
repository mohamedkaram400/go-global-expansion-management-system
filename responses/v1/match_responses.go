package responses

import (
	"time"
)

type MatchResponse struct {
	ID         int       `json:"id"`
	ProjectID  int       `json:"project_id"`
	VendorID   int       `json:"vendor_id"`
	VendorName string    `json:"vendor_name"`
	Score      float64   `json:"score"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
