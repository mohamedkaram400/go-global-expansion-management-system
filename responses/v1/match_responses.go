package responses

import (
	"time"
)

type MatchResponse struct {
	ID         uint      `json:"id"`
	ProjectID  uint      `json:"project_id"`
	VendorID   uint      `json:"vendor_id"`
	VendorName string    `json:"vendor_name"`
	Score      float64   `json:"score"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}