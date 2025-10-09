package entities

import "time"

type Match struct {
	ID        uint    `json:"id" gorm:"primaryKey;autoIncrement"`

	ProjectID uint    `json:"project_id" gorm:"not null"`
	Project   Project `gorm:"foreignKey:ProjectID"`

	VendorID  uint    `json:"vendor_id" gorm:"not null"`
	Vendor    Vendor  `gorm:"foreignKey:VendorID"`

	Score     float64   `json:"score" gorm:"column:score;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type VendorAnalytics struct {
    ID            uint    `json:"id"`
    Name          string  `json:"name"`
    AvgMatchScore float64 `json:"avg_match_score"`
}
