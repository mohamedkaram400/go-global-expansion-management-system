package entities

import "time"
type Match struct {
	ID 		   uint `json:"id" 			 gorm:"primaryKey;autoIncrement"`
	ProjectId uint `json:"project_id"   gorm:"column:"project_id;not null"`
	VendorId  uint `json:"vendor_id"	 gorm:"column:"vendor_id;not null"`
	Score 	   float64 `json:"score"	 gorm:"column:"score;not null"`
	CreatedAt    time.Time `json:"created_at" 		gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" 		gorm:"column:updated_at;autoUpdateTime"`
}